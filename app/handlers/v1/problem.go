package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schemas"
)

// List Problems
// (GET /domains/{domain}/problems)
func (s *Api) ListProblems(
	c *fiber.Ctx,
	request schemas.ListProblemsRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Create Problem
// (POST /domains/{domain}/problems)
func (s *Api) CreateProblem(
	c *fiber.Ctx,
	request schemas.CreateProblemRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Clone Problem
// (POST /domains/{domain}/problems/clone)
func (s *Api) CloneProblem(
	c *fiber.Ctx,
	request schemas.CloneProblemRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Delete Problem
// (DELETE /domains/{domain}/problems/{problem})
func (s *Api) DeleteProblem(
	c *fiber.Ctx,
	request schemas.DeleteProblemRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Get Problem
// (GET /domains/{domain}/problems/{problem})
func (s *Api) GetProblem(
	c *fiber.Ctx,
	request schemas.GetProblemRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Update Problem
// (PATCH /domains/{domain}/problems/{problem})
func (s *Api) UpdateProblem(
	c *fiber.Ctx,
	request schemas.UpdateProblemRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Submit Solution To Problem
// (POST /domains/{domain}/problems/{problem})
func (s *Api) SubmitSolutionToProblem(
	c *fiber.Ctx,
	request schemas.SubmitSolutionToProblemRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}