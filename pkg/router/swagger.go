package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func RegisterSwagger(router fiber.Router) {
	docs := router.Group("/docs")
	docs.Get("/*", swagger.New(swagger.Config{
		URL:          "/api/v1/horse/static/openapi.json",
		DocExpansion: "none",
	}))
}
