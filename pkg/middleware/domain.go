package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/query"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/platform/db"
	"github.com/sirupsen/logrus"
)

func Domain(c *fiber.Ctx) error {
	domainUrl := c.Params("domain")
	logrus.Debugf("domainUrl: %v", domainUrl)
	if len(domainUrl) != 0 {
		domain, err := query.GetDomain(db.DB, domainUrl)
		if err != nil {
			return c.Status(fiber.StatusOK).
				JSON(schema.NewEmptyResp(schema.DomainNotFoundError, "invalid domain Id or url"))
		}
		c.Locals("domain", &domain)
		c.Locals("domainId", domain.Id.String()) // for casbin only
	}
	return c.Next()
}
