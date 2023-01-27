package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schemas"
)

func ResponseHandler(ctx *fiber.Ctx, response any, err error) error {
	code := fiber.StatusOK
	if err == nil {
		if response == nil {
			return ctx.Status(code).JSON(schemas.NewEmptyResp(schemas.Success))
		}
		switch typedResponse := response.(type) {
		// non standard response
		case schemas.NonStandardResp:
			return ctx.Status(code).JSON(typedResponse.Data)
		default:
			// it should always be a struct
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
	// Retrieve the custom status code if it's a schemas.BizError
	if e, ok := err.(schemas.BizError); ok {
		return ctx.Status(code).
			JSON(schemas.NewEmptyResp(e.ErrorCode, *e.ErrorMsg))
	} else if errors.Is(err, fiber.ErrUnprocessableEntity) {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response)
	} else if errors.Is(err, fiber.ErrInternalServerError) {
		msg := response.(error).Error()
		return ctx.Status(code).
			JSON(schemas.NewEmptyResp(schemas.InternalServerError, msg))
	}
	return err
}
