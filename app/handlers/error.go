package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schemas"
)

func Error(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusOK).
		JSON(schemas.NewEmptyResp(schemas.InternalServerError))
}
