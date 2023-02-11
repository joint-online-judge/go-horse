package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schema"
)

// List Problem Groups
// (GET /problem_groups)
func (s *Api) ListProblemGroups(
	c *fiber.Ctx,
	request schema.ListProblemGroupsRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}
