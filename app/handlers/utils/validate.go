package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schemas"
	"github.com/joint-online-judge/go-horse/pkg/logger"
)

var validate = validator.New()

func isDomainUrl(fl validator.FieldLevel) bool {
	return fl.Field().String() != "domains"
}

func init() {
	if err := validate.RegisterValidation("domain_url", isDomainUrl); err != nil {
		logger.Fatalf("failed to register validation domain_url, %+v", err)
	}
}

func ValidateStruct(object any) (any, error) {
	logger.Infof("validating %T as struct, %v", object, object)
	var validationError []schemas.ValidationError
	if err := validate.Struct(object); err != nil {
		vierr, ok := err.(*validator.InvalidValidationError)
		if ok {
			logger.Errorf("invalid validation error: %v", vierr)
			return vierr, fiber.ErrInternalServerError
		}
		for _, e := range err.(validator.ValidationErrors) {
			logger.Errorf(
				"validation error: %v, %v, %v",
				e.StructNamespace(),
				e.Tag(),
				e.Param(),
			)
			validationError = append(
				validationError,
				schemas.ValidationError{
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
	f schemas.StrictHandlerFunc,
	operationID string,
) schemas.StrictHandlerFunc {
	return func(ctx *fiber.Ctx, request any) (any, error) {
		if response, err := ValidateStruct(request); err != nil {
			return response, err
		}
		return f(ctx, request)
	}
}
