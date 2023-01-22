package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/schemas"
)

func RegisterAuth(router fiber.Router, wrapper schemas.ServerInterfaceWrapper) {
	auth := router.Group("/auth")
	auth.Post("/login", wrapper.Login)
	auth.Post("/logout", wrapper.Logout)
	auth.Get("/oauth2", wrapper.ListOauth2)
	auth.Get("/oauth2/:oauth2/authorize", wrapper.OauthAuthorize)
	auth.Post("/refresh", wrapper.Refresh)
	auth.Post("/register", wrapper.Register)
	auth.Get("/token", wrapper.GetToken)
}
