package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/pkg/middleware"
)

func RegisterMisc(router fiber.Router, wrapper schema.ServerInterfaceWrapper) {
	router.Get("/test/report", wrapper.TestErrorReport)
	router.Get("/version", wrapper.Version)
	router.Get("/jwt_decoded", middleware.JWT(), wrapper.JwtDecoded)
}
