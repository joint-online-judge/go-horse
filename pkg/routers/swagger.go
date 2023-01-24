package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func RegisterSwagger(router fiber.Router) {
	docs := router.Group("/docs")
	docs.Get("/*", swagger.HandlerDefault) // get one user by ID
}
