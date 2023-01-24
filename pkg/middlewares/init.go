package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/joint-online-judge/go-horse/pkg/configs"
)

func Initalize(router *fiber.App) {
	router.Use(cors.New(cors.Config{
		AllowOrigins: "*", // comma string format
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	router.Use(Recover)
	router.Use(Security)
	if !configs.Conf.Debug {
		router.Use(csrf.New())
	}
	router.Use(Json)
}
