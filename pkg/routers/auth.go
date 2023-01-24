package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schemas"
	"github.com/joint-online-judge/go-horse/pkg/middlewares"
)

func RegisterAuth(router fiber.Router, wrapper schemas.ServerInterfaceWrapper) {
	auth := router.Group("/auth")
	auth.Post("/login", wrapper.Login)
	auth.Get("/oauth2", wrapper.ListOauth2)
	auth.Get("/oauth2/:oauth2/authorize", wrapper.OauthAuthorize)
	auth.Post("/refresh", wrapper.Refresh)
	auth.Get("/token", wrapper.GetToken)
	auth.Post("/logout", middlewares.JWT(), wrapper.Logout)
	auth.Post("/register", middlewares.JWT(), wrapper.Register)
}
