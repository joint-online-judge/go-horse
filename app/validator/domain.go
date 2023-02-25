package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func isDomainUrl(fl validator.FieldLevel) bool {
	return fl.Field().String() != "domains"
}

func init() {
	if err := validate.RegisterValidation("domain_url", isDomainUrl); err != nil {
		logrus.Fatalf("failed to register validation domain_url, %+v", err)
	}
}
