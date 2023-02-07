package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func RegisterSwagger(router fiber.Router) {
	docs := router.Group("/docs")
	docs.Get("/*", swagger.New(swagger.Config{
		URL:          "https://raw.githubusercontent.com/joint-online-judge/horse/openapi/openapi.json",
		DocExpansion: "none",
	}))
}
