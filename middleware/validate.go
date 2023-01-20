package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

var validate = validator.New()

func isDomainUrl(fl validator.FieldLevel) bool {
	return fl.Field().String() != "domains"
}

func init() {
	validate.RegisterValidation("domain_url", isDomainUrl)
}

func ValidateBody(c *fiber.Ctx) error {
	p := c.Locals("body")
	if err := validate.Struct(p); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Errorf("validation error: %v, %v, %v", err.StructNamespace(), err.Tag(), err.Param())
		}
		return fiber.ErrUnprocessableEntity
	}
	return c.Next()
}
