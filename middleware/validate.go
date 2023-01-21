package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/types"
	log "github.com/sirupsen/logrus"
)

var validate = validator.New()

func isDomainUrl(fl validator.FieldLevel) bool {
	return fl.Field().String() != "domains"
}

func init() {
	validate.RegisterValidation("domain_url", isDomainUrl)
}

func validateImpl(request any) (any, error) {
	var validationError []types.ValidationError
	log.Infof("validating %T", request)
	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Errorf("validation error: %v, %v, %v", err.StructNamespace(), err.Tag(), err.Param())
			validationError = append(validationError, types.ValidationError{Msg: err.Error(), Type: err.Type().Name()})
		}
		return validationError, fiber.ErrUnprocessableEntity
	}
	return nil, nil
}

func Validate(f types.StrictHandlerFunc, operationID string) types.StrictHandlerFunc {
	return func(ctx *fiber.Ctx, request any) (any, error) {
		response, err := validateImpl(request)
		if err != nil {
			return response, err
		}
		return f(ctx, request)
	}
}
