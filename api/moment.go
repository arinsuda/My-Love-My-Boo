package handler

import (
	"encoding/json"
	"net/http"
)

type Moment struct {
	ID          int      `json:"id"`
	Date        string   `json:"date"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
	Tag         string   `json:"tag"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	moments := []Moment{
		{
			ID:          1,
			Date:        "2019-08-10",
			Title:       "วันแรกที่เราเจอกัน",
			Description: "ตอนนั้นยังไม่รู้เลยว่าจะกลายเป็นคนสำคัญขนาดนี้ 💖",
			Images:      []string{"/images/first-day-1.jpg"},
			Tag:         "เริ่มต้น",
		},
		{
			ID:          2,
			Date:        "2020-02-14",
			Title:       "วาเลนไทน์ครั้งแรก",
			Description: "ช็อกโกแลตวันนั้น กับรอยยิ้มของคุณ ยังจำได้อยู่เลย 🍫",
			Images:      []string{"/images/valentine-2020.jpg"},
			Tag:         "เดต",
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(moments)
}
