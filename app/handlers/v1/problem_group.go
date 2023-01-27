package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schemas"
)

// List Problem Groups
// (GET /problem_groups)
func (s *Api) ListProblemGroups(
	c *fiber.Ctx,
	request schemas.ListProblemGroupsRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}
