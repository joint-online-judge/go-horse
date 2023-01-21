package middleware

import (
	"errors"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/types"
	log "github.com/sirupsen/logrus"
)

func Response(ctx *fiber.Ctx, response any, err error) error {
	code := fiber.StatusOK
	if err == nil {
		validate_response, validate_err := validateImpl(response)
		if validate_err == nil {
			return ctx.Status(code).JSON(types.StandardResp[any]{
				BizError: types.BizError{ErrorCode: types.Success},
				Data:     response,
			})
		}
		err = validate_err
		response = validate_response
	}
	// Retrieve the custom status code if it's a *types.BizError
	var e *types.BizError
	if errors.As(err, &e) {
		return ctx.Status(code).JSON(types.EmptyResp{BizError: *e, Data: nil})
	}
	if errors.Is(err, fiber.ErrUnprocessableEntity) {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response)
	}
	log.Info(err)
	b := make([]byte, 4096)
	n := runtime.Stack(b, false)
	s := string(b[:n])
	return ctx.Status(code).JSON(types.EmptyResp{BizError: types.BizError{
		ErrorCode: types.InternalServerError, ErrorMsg: &s}, Data: nil})
}
