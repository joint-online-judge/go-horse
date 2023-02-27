package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/query"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/platform/db"
	"github.com/sirupsen/logrus"
)

func Problem(c *fiber.Ctx) error {
	problemUrl := c.Params("problem")
	logrus.Debugf("problemUrl: %v", problemUrl)
	if len(problemUrl) != 0 {
		domain, ok := c.Locals("domain").(*model.Domain)
		if !ok {
			return c.Status(fiber.StatusOK).
				JSON(schema.NewEmptyResp(schema.DomainNotFoundError, "invalid domain Id or url"))
		}
		problem, err := query.GetProblem(db.DB, domain, problemUrl)
		if err != nil {
			return c.Status(fiber.StatusOK).
				JSON(schema.NewEmptyResp(schema.ProblemNotFoundError, "invalid problem Id or url"))
		}
		c.Locals("problem", &problem)
	}
	return c.Next()
}
