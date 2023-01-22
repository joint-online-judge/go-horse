package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/handlers"
	"github.com/joint-online-judge/go-horse/handlers/utils"
	"github.com/joint-online-judge/go-horse/middleware"
	"github.com/joint-online-judge/go-horse/types"
)

func Initalize(router *fiber.App) {
	router.Use(middleware.Recover)
	router.Use(middleware.Security)

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello, World!")
	})

	router.Use(middleware.Json)

	var strictHandlerImpl handlers.ApiV1
	middlewares := []types.StrictMiddlewareFunc{utils.Validate}
	response := utils.Response
	RegisterStrictHandlers(router, &strictHandlerImpl, middlewares, response)
}
