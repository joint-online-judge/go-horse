package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schemas"
)

// List Problems
// (GET /domains/{domain}/problems)
func (s *ApiV1) ListProblems(
	c *fiber.Ctx,
	request schemas.ListProblemsRequestObject,
) (any, error) {
	return nil, nil
}

// Create Problem
// (POST /domains/{domain}/problems)
func (s *ApiV1) CreateProblem(
	c *fiber.Ctx,
	request schemas.CreateProblemRequestObject,
) (any, error) {
	return nil, nil
}

// Clone Problem
// (POST /domains/{domain}/problems/clone)
func (s *ApiV1) CloneProblem(
	c *fiber.Ctx,
	request schemas.CloneProblemRequestObject,
) (any, error) {
	return nil, nil
}

// Delete Problem
// (DELETE /domains/{domain}/problems/{problem})
func (s *ApiV1) DeleteProblem(
	c *fiber.Ctx,
	request schemas.DeleteProblemRequestObject,
) (any, error) {
	return nil, nil
}

// Get Problem
// (GET /domains/{domain}/problems/{problem})
func (s *ApiV1) GetProblem(
	c *fiber.Ctx,
	request schemas.GetProblemRequestObject,
) (any, error) {
	return nil, nil
}

// Update Problem
// (PATCH /domains/{domain}/problems/{problem})
func (s *ApiV1) UpdateProblem(
	c *fiber.Ctx,
	request schemas.UpdateProblemRequestObject,
) (any, error) {
	return nil, nil
}

// Submit Solution To Problem
// (POST /domains/{domain}/problems/{problem})
func (s *ApiV1) SubmitSolutionToProblem(
	c *fiber.Ctx,
	request schemas.SubmitSolutionToProblemRequestObject,
) (any, error) {
	return nil, nil
}
