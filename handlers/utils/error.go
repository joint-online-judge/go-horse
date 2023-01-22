package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/types"
)

// this will only handle error from panic
func Panic(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusOK).JSON(types.EmptyResp{BizError: types.BizError{
		ErrorCode: types.InternalServerError, ErrorMsg: nil}, Data: nil})
}
