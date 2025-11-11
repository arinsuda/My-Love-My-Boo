package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	// โหลดค่า .env
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found, using default PORT=8080")
	}

	// อ่านพอร์ตจาก .env หรือใช้ 8080 ถ้าไม่ตั้งค่า
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := fiber.New()

	// Health check route
	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("OK 💖")
	})

	// ตัวอย่าง route /api/moments
	app.Get("/api/moments", func(c fiber.Ctx) error {
		moments := []fiber.Map{
			{
				"id":          1,
				"date":        "2019-08-10",
				"title":       "วันแรกที่เราเจอกัน",
				"description": "ตอนนั้นยังไม่รู้เลยว่าจะกลายเป็นคนสำคัญขนาดนี้",
				"images":      []string{"/public/first-meeting.jpg"},
				"tag":         "เริ่มต้น",
			},
			{
				"id":          2,
				"date":        "2020-02-14",
				"title":       "วาเลนไทน์ครั้งแรก",
				"description": "ช็อกโกแลตวันนั้น กับรอยยิ้มของคุณ ยังจำได้อยู่เลย",
				"images":      []string{"/public/valentine-2020.jpg"},
				"tag":         "เดต",
			},
		}
		return c.JSON(moments)
	})

	// เริ่ม server
	log.Printf("🚀 Server running on http://localhost:%s", port)
	log.Fatal(app.Listen(":" + port))
}
