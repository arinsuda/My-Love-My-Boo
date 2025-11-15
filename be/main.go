package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	SupabaseURL    string
	SupabaseKey    string
	SupabaseBucket string
	Port           string
	AppOrigin      string
}

var cfg Config

func main() {
	// โหลด .env
	_ = godotenv.Load(".env", ".env.production")

	cfg = Config{
		SupabaseURL:    os.Getenv("SUPABASE_URL"),
		SupabaseKey:    os.Getenv("SUPABASE_SERVICE_ROLE_KEY"),
		SupabaseBucket: os.Getenv("SUPABASE_BUCKET"),
		Port:           os.Getenv("APP_PORT"),
		AppOrigin:      os.Getenv("APP_ORIGIN"),
	}

	if cfg.Port == "" {
		cfg.Port = os.Getenv("PORT")
	}
	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	if cfg.AppOrigin == "" {
		// default ตอน dev
		cfg.AppOrigin = "http://localhost:5173"
	}
	if cfg.SupabaseURL == "" || cfg.SupabaseKey == "" || cfg.SupabaseBucket == "" {
		log.Fatal("missing SUPABASE_URL / SUPABASE_SERVICE_ROLE_KEY / SUPABASE_BUCKET in .env")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/api/upload", uploadHandler)
	mux.HandleFunc("/api/moments", momentsHandler)

	handler := withCORS(mux)

	log.Printf("Server running on :%s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handler))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// รองรับไฟล์ไม่เกิน ~20MB
	if err := r.ParseMultipartForm(20 << 20); err != nil {
		http.Error(w, "cannot parse form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "cannot read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// ตั้งชื่อไฟล์บน Supabase
	filename := fmt.Sprintf("memories/%d-%s", time.Now().Unix(), header.Filename)

	// อ่านไฟล์มาเก็บใน buffer
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, file); err != nil {
		http.Error(w, "cannot read file", http.StatusInternalServerError)
		return
	}

	// อัปโหลดไป Supabase Storage
	uploadURL := fmt.Sprintf("%s/storage/v1/object/%s/%s", cfg.SupabaseURL, cfg.SupabaseBucket, filename)

	req, err := http.NewRequest("POST", uploadURL, bytes.NewReader(buf.Bytes()))
	if err != nil {
		http.Error(w, "cannot create request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Authorization", "Bearer "+cfg.SupabaseKey)
	req.Header.Set("Content-Type", header.Header.Get("Content-Type"))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, "upload failed", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		body, _ := io.ReadAll(res.Body)
		log.Println("Supabase upload error:", string(body))
		http.Error(w, "upload error", http.StatusInternalServerError)
		return
	}

	// bucket เป็น private → ขอ signed URL ให้เลย
	signedURL, err := createSignedURL(filename)
	if err != nil {
		log.Println("Supabase signed url error:", err)
		http.Error(w, "cannot create signed url", http.StatusInternalServerError)
		return
	}

	resp := map[string]string{
		"path":      filename,
		"signedUrl": signedURL,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func momentsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data, err := os.ReadFile("moments.json")
	if err != nil {
		http.Error(w, "cannot load moments", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func createSignedURL(path string) (string, error) {
	url := fmt.Sprintf("%s/storage/v1/object/sign/%s/%s", cfg.SupabaseURL, cfg.SupabaseBucket, path)

	// ใช้ตัวเลขจริง ๆ ไม่ใช่ "60*60*24"
	payload := map[string]int{
		"expiresIn": 60 * 60 * 24, // 1 วัน
	}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+cfg.SupabaseKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		b, _ := io.ReadAll(res.Body)
		return "", fmt.Errorf("supabase sign error: %s", string(b))
	}

	var data struct {
		SignedURL string `json:"signedURL"`
	}
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return "", err
	}

	// Supabase คืน path ต่อจาก base url → ต่อเอง
	return cfg.SupabaseURL + data.SignedURL, nil
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"status": "ok"}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", cfg.AppOrigin)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
