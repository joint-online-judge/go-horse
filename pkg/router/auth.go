package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/pkg/middleware"
)

func RegisterAuth(router fiber.Router, wrapper schema.ServerInterfaceWrapper) {
	auth := router.Group("/auth")
	auth.Post("/login", wrapper.Login)
	auth.Post("/register", wrapper.Register)
	auth.Get("/oauth2", wrapper.ListOauth2)
	auth.Get("/oauth2/:oauth2/authorize", wrapper.OauthAuthorize)
	auth.Post("/logout", middleware.JWT(), wrapper.Logout)
	auth.Post("/refresh", middleware.JWT(), wrapper.Refresh)
	auth.Get("/token", middleware.JWT(), wrapper.GetToken)
}
