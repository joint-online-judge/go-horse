package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schemas"
)

// List Problem Sets
// (GET /domains/{domain}/problem_sets)
func (s *ApiV1) ListProblemSets(
	c *fiber.Ctx,
	request schemas.ListProblemSetsRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Create Problem Set
// (POST /domains/{domain}/problem_sets)
func (s *ApiV1) CreateProblemSet(
	c *fiber.Ctx,
	request schemas.CreateProblemSetRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Delete Problem Set
// (DELETE /domains/{domain}/problem_sets/{problemSet})
func (s *ApiV1) DeleteProblemSet(
	c *fiber.Ctx,
	request schemas.DeleteProblemSetRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Get Problem Set
// (GET /domains/{domain}/problem_sets/{problemSet})
func (s *ApiV1) GetProblemSet(
	c *fiber.Ctx,
	request schemas.GetProblemSetRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Update Problem Set
// (PATCH /domains/{domain}/problem_sets/{problemSet})
func (s *ApiV1) UpdateProblemSet(
	c *fiber.Ctx,
	request schemas.UpdateProblemSetRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// List Problems In Problem Set
// (GET /domains/{domain}/problem_sets/{problemSet}/problems)
func (s *ApiV1) ListProblemsInProblemSet(
	c *fiber.Ctx,
	request schemas.ListProblemsInProblemSetRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Add Problem In Problem Set
// (POST /domains/{domain}/problem_sets/{problemSet}/problems)
func (s *ApiV1) AddProblemInProblemSet(
	c *fiber.Ctx,
	request schemas.AddProblemInProblemSetRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Delete Problem In Problem Set
// (DELETE /domains/{domain}/problem_sets/{problemSet}/problems/{problem})
func (s *ApiV1) DeleteProblemInProblemSet(
	c *fiber.Ctx,
	request schemas.DeleteProblemInProblemSetRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Get Problem In Problem Set
// (GET /domains/{domain}/problem_sets/{problemSet}/problems/{problem})
func (s *ApiV1) GetProblemInProblemSet(
	c *fiber.Ctx,
	request schemas.GetProblemInProblemSetRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Update Problem In Problem Set
// (PATCH /domains/{domain}/problem_sets/{problemSet}/problems/{problem})
func (s *ApiV1) UpdateProblemInProblemSet(
	c *fiber.Ctx,
	request schemas.UpdateProblemInProblemSetRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Submit Solution To Problem Set
// (POST /domains/{domain}/problem_sets/{problemSet}/problems/{problem}/submit)
func (s *ApiV1) SubmitSolutionToProblemSet(
	c *fiber.Ctx,
	request schemas.SubmitSolutionToProblemSetRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}
