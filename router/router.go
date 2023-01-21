package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/handlers"
	"github.com/joint-online-judge/go-horse/middleware"
	"github.com/joint-online-judge/go-horse/types"
)

func Initalize(router *fiber.App) {
	router.Use(middleware.Security)

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello, World!")
	})

	router.Use(middleware.Json)

	var strictHandlerImpl handlers.ApiV1
	middlewares := []types.StrictMiddlewareFunc{middleware.Validate}
	response := middleware.Response
	RegisterStrictHandlers(router, &strictHandlerImpl, middlewares, response)

	users := router.Group("/users")
	users.Post("/", handlers.CreateUser)
	users.Delete("/", middleware.Authenticated, handlers.DeleteUser)
	users.Put("/", middleware.Authenticated, handlers.ChangePassword)
	users.Post("/me", middleware.Authenticated, handlers.GetUserInfo)
	users.Post("/login", handlers.Login)
	users.Delete("/logout", handlers.Logout)

	products := router.Group("/products", middleware.Authenticated)
	products.Post("/", handlers.CreateProduct)
	products.Post("/all", handlers.GetProducts)
	products.Delete("/:id", handlers.DeleteProduct)
	products.Post("/:id", handlers.GetProductById)
	products.Put("/:id", handlers.UpdateProduct)

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})

}
