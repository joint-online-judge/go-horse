package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/query"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/app/service"
	"github.com/joint-online-judge/go-horse/platform/db"
	"github.com/sirupsen/logrus"
)

func ProblemSet(c *fiber.Ctx) error {
	problemSetUrl := c.Params("problemSet")
	logrus.Debugf("problemSetUrl: %v", problemSetUrl)
	if len(problemSetUrl) != 0 {
		domain, ok := c.Locals("domain").(*model.Domain)
		if !ok {
			return c.Status(fiber.StatusOK).
				JSON(schema.NewEmptyResp(schema.DomainNotFoundError, "invalid domain Id or url"))
		}
		problemSet, err := service.ProblemSet(c).GetProblemSet(
			domain, problemSetUrl,
		)
		if err != nil {
			return c.Status(fiber.StatusOK).
				JSON(schema.NewEmptyResp(schema.ProblemSetNotFoundError, "invalid problemSet Id or url"))
		}
		c.Locals("problemSet", &problemSet)
	}
	return c.Next()
}

func ProblemInProblemSet(c *fiber.Ctx) error {
	problemUrl := c.Params("problem")
	logrus.Debugf("problemUrl: %v", problemUrl)
	if len(problemUrl) != 0 {
		domain, ok := c.Locals("domain").(*model.Domain)
		if !ok {
			return c.Status(fiber.StatusOK).
				JSON(schema.NewEmptyResp(schema.DomainNotFoundError, "invalid domain Id or url"))
		}
		problemSet, ok := c.Locals("problemSet").(*model.ProblemSet)
		if !ok {
			return c.Status(fiber.StatusOK).
				JSON(schema.NewEmptyResp(schema.ProblemSetNotFoundError, "invalid problemSet Id or url"))
		}
		problem, err := query.GetProblemInProblemSet(
			db.DB, domain, problemSet, problemUrl,
		)
		if err != nil {
			return c.Status(fiber.StatusOK).
				JSON(schema.NewEmptyResp(schema.ProblemNotFoundError, "invalid problem Id or url"))
		}
		c.Locals("problem", &problem)
	}
	return c.Next()
}
