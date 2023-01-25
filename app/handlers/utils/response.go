package utils

import (
	"errors"
	"reflect"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schemas"
	log "github.com/sirupsen/logrus"
)

func ResponseHandler(ctx *fiber.Ctx, response any, err error) error {
	code := fiber.StatusOK
	if err == nil {
		if response == nil {
			return ctx.Status(code).JSON(schemas.EmptyResp{
				BizError: schemas.BizError{ErrorCode: schemas.Success},
				Data:     nil,
			})
		}
		switch response.(type) {
		// non standard response
		case schemas.NonStandardResp:
			resp := response
			return ctx.Status(code).JSON(resp)
		default:
			if reflect.ValueOf(response).Kind() == reflect.Struct {
				validate_response, validate_err := ValidateStruct(response)
				if validate_err == nil {
					return ctx.Status(code).JSON(schemas.StandardResp[any]{
						BizError: schemas.BizError{ErrorCode: schemas.Success},
						Data:     response,
					})
				}
				err = validate_err
				response = validate_response
			}
		}
	}
	// Retrieve the custom status code if it's a schemas.BizError
	if e, ok := err.(schemas.BizError); ok {
		return ctx.Status(code).JSON(schemas.EmptyResp{BizError: e, Data: nil})
	}
	if errors.Is(err, fiber.ErrUnprocessableEntity) {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response)
	}
	if errors.Is(err, fiber.ErrInternalServerError) {
		msg := response.(error).Error()
		return ctx.Status(code).
			JSON(schemas.EmptyResp{BizError: schemas.BizError{
				ErrorCode: schemas.InternalServerError, ErrorMsg: &msg,
			}, Data: nil})
	}
	b := make([]byte, 4096)
	n := runtime.Stack(b, false)
	s := string(b[:n])
	log.Infof("unknown error: %T, %v, %s", err, err, s)
	return ctx.Status(code).JSON(schemas.EmptyResp{BizError: schemas.BizError{
		ErrorCode: schemas.InternalServerError, ErrorMsg: &s,
	}, Data: nil})
}
