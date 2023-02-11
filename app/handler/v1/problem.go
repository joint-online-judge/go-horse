package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schema"
)

// List Problems
// (GET /domains/{domain}/problems)
func (s *Api) ListProblems(
	c *fiber.Ctx,
	request schema.ListProblemsRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Create Problem
// (POST /domains/{domain}/problems)
func (s *Api) CreateProblem(
	c *fiber.Ctx,
	request schema.CreateProblemRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Clone Problem
// (POST /domains/{domain}/problems/clone)
func (s *Api) CloneProblem(
	c *fiber.Ctx,
	request schema.CloneProblemRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Delete Problem
// (DELETE /domains/{domain}/problems/{problem})
func (s *Api) DeleteProblem(
	c *fiber.Ctx,
	request schema.DeleteProblemRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Get Problem
// (GET /domains/{domain}/problems/{problem})
func (s *Api) GetProblem(
	c *fiber.Ctx,
	request schema.GetProblemRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Update Problem
// (PATCH /domains/{domain}/problems/{problem})
func (s *Api) UpdateProblem(
	c *fiber.Ctx,
	request schema.UpdateProblemRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Submit Solution To Problem
// (POST /domains/{domain}/problems/{problem})
func (s *Api) SubmitSolutionToProblem(
	c *fiber.Ctx,
	request schema.SubmitSolutionToProblemRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}
