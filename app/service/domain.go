package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/schema"
)

type domainImpl struct {
	c *fiber.Ctx
}

func Domain(c *fiber.Ctx) *domainImpl {
	return &domainImpl{
		c: c,
	}
}

func (s *domainImpl) Domain() *model.Domain {
	return s.c.Locals("domain").(*model.Domain)
}

// GetCurrentDomain Assume we require current domain to be non-null when calling this method
func (s *domainImpl) GetCurrentDomain() (*model.Domain, error) {
	domain := s.c.Locals("domain")
	if domain == nil {
		return nil, schema.NewBizError(schema.DomainNotFoundError, "cannot find current domain")
	}
	domainModel := domain.(model.Domain)
	return &domainModel, nil
}
