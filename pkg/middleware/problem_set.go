package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/app/service"
	"github.com/sirupsen/logrus"
)

func ProblemSet(c *fiber.Ctx) error {
	problemSetUrl := c.Params("problemSet")
	logrus.Debugf("problemSetUrl: %v", problemSetUrl)
	if len(problemSetUrl) != 0 {
		domain, ok := c.Locals("domain").(*model.Domain)
		if !ok {
			return c.Status(fiber.StatusOK).
				JSON(schema.NewEmptyResp(schema.DomainNotFoundError,
					"invalid domain Id or url"))
		}
		problemSet, err := service.ProblemSet(c).GetProblemSet(
			domain, problemSetUrl,
		)
		if err != nil {
			return c.Status(fiber.StatusOK).
				JSON(schema.NewEmptyResp(schema.ProblemSetNotFoundError,
					"invalid problemSet Id or url"))
		}
		c.Locals("problemSet", &problemSet)
	}
	return c.Next()
}
