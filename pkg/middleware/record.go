package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/query"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/platform/db"
	"github.com/sirupsen/logrus"
)

func Record(c *fiber.Ctx) error {
	recordUrl := c.Params("record")
	logrus.Debugf("recordUrl: %v", recordUrl)
	if len(recordUrl) != 0 {
		domain, ok := c.Locals("domain").(*model.Domain)
		if !ok {
			return c.Status(fiber.StatusOK).
				JSON(schema.NewEmptyResp(schema.DomainNotFoundError, "invalid domain Id or url"))
		}
		record, err := query.GetRecord(db.DB, domain, recordUrl)
		if err != nil {
			return c.Status(fiber.StatusOK).
				JSON(schema.NewEmptyResp(schema.RecordNotFoundError, "invalid record Id or url"))
		}
		c.Locals("record", &record)
	}
	return c.Next()
}
