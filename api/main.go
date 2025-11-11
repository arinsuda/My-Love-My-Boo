package handler

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	storage "github.com/supabase-community/storage-go"
)

func main() {
	_ = godotenv.Load(".env")

	app := fiber.New()

	supabaseURL := os.Getenv("SUPABASE_URL")
	serviceKey := os.Getenv("SUPABASE_SERVICE_ROLE_KEY")
	_ = storage.NewClient(supabaseURL, serviceKey, nil)

	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("OK 💖")
	})

	log.Println("listening on :8080")
	app.Listen(":8080")
}
