package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schemas"
	"github.com/joint-online-judge/go-horse/pkg/middlewares"
)

func RegisterAuth(router fiber.Router, wrapper schemas.ServerInterfaceWrapper) {
	auth := router.Group("/auth")
	auth.Post("/login", wrapper.Login)
	auth.Post("/register", wrapper.Register)
	auth.Get("/oauth2", wrapper.ListOauth2)
	auth.Get("/oauth2/:oauth2/authorize", wrapper.OauthAuthorize)
	auth.Post("/logout", middlewares.JWT(), wrapper.Logout)
	auth.Post("/refresh", middlewares.JWT(), wrapper.Refresh)
	auth.Get("/token", middlewares.JWT(), wrapper.GetToken)
}
