package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/schemas"
)

// this will only handle error from panic
func Panic(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusOK).JSON(schemas.EmptyResp{BizError: schemas.BizError{
		ErrorCode: schemas.InternalServerError, ErrorMsg: nil,
	}, Data: nil})
}
