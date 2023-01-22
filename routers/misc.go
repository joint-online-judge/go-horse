package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/schemas"
)

func RegisterMisc(router fiber.Router, wrapper schemas.ServerInterfaceWrapper) {
	router.Get("/jwt_decoded", wrapper.JwtDecoded)
	router.Get("/test/report", wrapper.TestErrorReport)
	router.Get("/version", wrapper.Version)
}
