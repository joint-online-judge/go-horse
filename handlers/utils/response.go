package utils

import (
	"errors"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/db"
	"github.com/joint-online-judge/go-horse/types"
	log "github.com/sirupsen/logrus"
)

func ResponseHandler(ctx *fiber.Ctx, response any, err error) error {
	code := fiber.StatusOK
	if err == nil {
		if response == nil {
			return ctx.Status(code).JSON(types.EmptyResp{
				BizError: types.BizError{ErrorCode: types.Success},
				Data:     nil,
			})
		}
		validate_response, validate_err := ValidateStruct(response)
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
	if errors.Is(err, fiber.ErrInternalServerError) {
		msg := response.(error).Error()
		return ctx.Status(code).JSON(types.EmptyResp{BizError: types.BizError{
			ErrorCode: types.InternalServerError, ErrorMsg: &msg,
		}, Data: nil})
	}
	log.Infof("unknown error: %T, %v", err, err)
	b := make([]byte, 4096)
	n := runtime.Stack(b, false)
	s := string(b[:n])
	return ctx.Status(code).JSON(types.EmptyResp{BizError: types.BizError{
		ErrorCode: types.InternalServerError, ErrorMsg: &s,
	}, Data: nil})
}

func NewListResp[Model any, Types any]() (types.ListResp[Types], error) {
	objs, count, err := db.ListObjs[Model, Types]()
	return types.ListResp[Types]{Count: count, Results: objs}, err
}
