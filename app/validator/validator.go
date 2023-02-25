package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/sirupsen/logrus"
)

var validate = validator.New()

func ValidateStruct(object any) (any, error) {
	logrus.Infof("validating %T as struct, %v", object, object)
	var validationError []schema.ValidationError
	if err := validate.Struct(object); err != nil {
		vierr, ok := err.(*validator.InvalidValidationError)
		if ok {
			logrus.Errorf("invalid validation error: %v", vierr)
			return vierr, fiber.ErrInternalServerError
		}
		for _, e := range err.(validator.ValidationErrors) {
			logrus.Errorf(
				"validation error: %v, %v, %v",
				e.StructNamespace(),
				e.Tag(),
				e.Param(),
			)
			validationError = append(
				validationError,
				schema.ValidationError{
					Msg:  e.Error(),
					Type: e.Type().Name(),
				},
			)
		}
		return validationError, fiber.ErrUnprocessableEntity
	}
	return nil, nil
}

func ValidateRequest(
	f schema.StrictHandlerFunc,
	operationID string,
) schema.StrictHandlerFunc {
	return func(ctx *fiber.Ctx, request any) (any, error) {
		if response, err := ValidateStruct(request); err != nil {
			return response, err
		}
		return f(ctx, request)
	}
}
