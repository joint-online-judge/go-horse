package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Initalize(router *fiber.App) {
	router.Use(cors.New(cors.Config{
		AllowOrigins: "*", // comma string format
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	router.Use(Recover)
	router.Use(Security)
	// FIXME: there may have file uploads
	// router.Use(Json)
}
