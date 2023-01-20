package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/joint-online-judge/go-horse/docs"
	"github.com/joint-online-judge/go-horse/handlers"
	"github.com/joint-online-judge/go-horse/middleware"
)

func Initalize(router *fiber.App) {
	// TODO: implement handler impl
	// var handlerImpl HandlerImpl
	// RegisterHandlers(router, &handlerImpl)
	router.Get("/swagger/*", swagger.HandlerDefault) // default

	router.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:3000/swagger/oauth2-redirect.html",
	}))

	router.Use(middleware.Security)

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello, World!")
	})

	router.Use(middleware.Json)

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
