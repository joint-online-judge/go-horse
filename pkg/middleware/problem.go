package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/app/service"
	"github.com/sirupsen/logrus"
)

func problemImpl(c *fiber.Ctx, checkProblemSet bool) error {
	problemUrl := c.Params("problem")
	logrus.Debugf("problemUrl: %v", problemUrl)
	if len(problemUrl) != 0 {
		domain, ok := c.Locals("domain").(*model.Domain)
		if !ok {
			return c.Status(fiber.StatusOK).
				JSON(schema.NewEmptyResp(schema.DomainNotFoundError,
					"invalid domain Id or url"))
		}
		var problemSet *model.ProblemSet
		if checkProblemSet {
			// TODO: check hidden
			problemSet, ok = c.Locals("problemSet").(*model.ProblemSet)
			if !ok {
				return c.Status(fiber.StatusOK).
					JSON(schema.NewEmptyResp(schema.ProblemSetNotFoundError,
						"invalid problemSet Id or url"))
			}
		}
		// TODO: check hidden
		problem, err := service.Problem(c).GetProblem(
			domain, problemSet, problemUrl,
		)
		if err != nil {
			return c.Status(fiber.StatusOK).
				JSON(schema.NewEmptyResp(schema.ProblemNotFoundError,
					"invalid problem Id or url"))
		}
		c.Locals("problem", &problem)
	}
	return c.Next()
}

func Problem(c *fiber.Ctx) error {
	return problemImpl(c, false)
}

func ProblemInProblemSet(c *fiber.Ctx) error {
	return problemImpl(c, true)
}
