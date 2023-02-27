package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/app/service"
	"github.com/joint-online-judge/go-horse/pkg/convert"
)

// List Problem Sets
// (GET /domains/{domain}/problem_sets)
func (s *Api) ListProblemSets(
	c *fiber.Ctx,
	request schema.ListProblemSetsRequestObject,
) (any, error) {
	return service.ProblemSet(c).ListProblemSets(request.Params)
}

// Create Problem Set
// (POST /domains/{domain}/problem_sets)
func (s *Api) CreateProblemSet(
	c *fiber.Ctx,
	request schema.CreateProblemSetRequestObject,
) (any, error) {
	problemSet, err := service.ProblemSet(c).CreateProblemSet(*request.Body)
	if err != nil {
		return nil, err
	}
	return convert.To[schema.ProblemSet](problemSet)
}

// Delete Problem Set
// (DELETE /domains/{domain}/problem_sets/{problemSet})
func (s *Api) DeleteProblemSet(
	c *fiber.Ctx,
	request schema.DeleteProblemSetRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Get Problem Set
// (GET /domains/{domain}/problem_sets/{problemSet})
func (s *Api) GetProblemSet(
	c *fiber.Ctx,
	request schema.GetProblemSetRequestObject,
) (any, error) {
	problemSet, err := service.ProblemSet(c).GetCurrentProblemSet()
	if err != nil {
		return nil, err
	}
	return convert.To[schema.ProblemSet](problemSet)
}

// Update Problem Set
// (PATCH /domains/{domain}/problem_sets/{problemSet})
func (s *Api) UpdateProblemSet(
	c *fiber.Ctx,
	request schema.UpdateProblemSetRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// List Problems In Problem Set
// (GET /domains/{domain}/problem_sets/{problemSet}/problems)
func (s *Api) ListProblemsInProblemSet(
	c *fiber.Ctx,
	request schema.ListProblemsInProblemSetRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Add Problem In Problem Set
// (POST /domains/{domain}/problem_sets/{problemSet}/problems)
func (s *Api) AddProblemInProblemSet(
	c *fiber.Ctx,
	request schema.AddProblemInProblemSetRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Delete Problem In Problem Set
// (DELETE /domains/{domain}/problem_sets/{problemSet}/problems/{problem})
func (s *Api) DeleteProblemInProblemSet(
	c *fiber.Ctx,
	request schema.DeleteProblemInProblemSetRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Get Problem In Problem Set
// (GET /domains/{domain}/problem_sets/{problemSet}/problems/{problem})
func (s *Api) GetProblemInProblemSet(
	c *fiber.Ctx,
	request schema.GetProblemInProblemSetRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Update Problem In Problem Set
// (PATCH /domains/{domain}/problem_sets/{problemSet}/problems/{problem})
func (s *Api) UpdateProblemInProblemSet(
	c *fiber.Ctx,
	request schema.UpdateProblemInProblemSetRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Submit Solution To Problem Set
// (POST /domains/{domain}/problem_sets/{problemSet}/problems/{problem}/submit)
func (s *Api) SubmitSolutionToProblemSet(
	c *fiber.Ctx,
	request schema.SubmitSolutionToProblemSetRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}
