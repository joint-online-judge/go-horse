package main

import (
	"os"

	"github.com/joint-online-judge/go-horse/db"
	"github.com/joint-online-judge/go-horse/handlers/utils"
	"github.com/joint-online-judge/go-horse/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func getenv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.Panic,
	})
	db.ConnectDB()
	routers.Initalize(app)
	log.Fatal(app.Listen(":" + getenv("PORT", "3000")))
}
