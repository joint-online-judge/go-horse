package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/app/validator"
)

func ResponseHandler(ctx *fiber.Ctx, response any, err error) error {
	code := fiber.StatusOK
	if err == nil {
		if response == nil {
			return ctx.Status(code).JSON(schema.NewEmptyResp(schema.Success))
		}
		switch typedResponse := response.(type) {
		// non standard response
		case schema.NonStandardResp:
			return ctx.Status(code).JSON(typedResponse.Data)
		default:
			// it should always be a struct
			validate_response, validate_err := validator.ValidateStruct(response)
			if validate_err == nil {
				return ctx.Status(code).JSON(schema.StandardResp[any]{
					BizError: schema.BizError{ErrorCode: schema.Success},
					Data:     response,
				})
			}
			err = validate_err
			response = validate_response
		}
	}
	// Retrieve the custom status code if it's a schema.BizError
	if e, ok := err.(schema.BizError); ok {
		if e.ErrorMsg == nil {
			return ctx.Status(code).
				JSON(schema.NewEmptyResp(e.ErrorCode))
		} else {
			return ctx.Status(code).
				JSON(schema.NewEmptyResp(e.ErrorCode, *e.ErrorMsg))
		}
	} else if errors.Is(err, fiber.ErrUnprocessableEntity) {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response)
	} else if errors.Is(err, fiber.ErrInternalServerError) {
		msg := response.(error).Error()
		return ctx.Status(code).
			JSON(schema.NewEmptyResp(schema.InternalServerError, msg))
	}
	return err
}
