package handlers

import (
	"errors"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/types"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusOK
	// Retrieve the custom status code if it's a *types.BizError
	var e *types.BizError
	if errors.As(err, &e) {
		return ctx.Status(code).JSON(types.EmptyResp{BizError: *e, Data: nil})
	}
	b := make([]byte, 2048)
	n := runtime.Stack(b, false)
	s := string(b[:n])
	return ctx.Status(code).JSON(types.EmptyResp{BizError: types.BizError{
		ErrorCode: types.InternalServerError, ErrorMsg: &s}, Data: nil})
}
