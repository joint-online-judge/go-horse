package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/middlewares"
	"github.com/joint-online-judge/go-horse/schemas"
)

func RegisterMisc(router fiber.Router, wrapper schemas.ServerInterfaceWrapper) {
	router.Get("/test/report", wrapper.TestErrorReport)
	router.Get("/version", wrapper.Version)
	router.Get("/jwt_decoded", middlewares.JWT(), wrapper.JwtDecoded)
}
