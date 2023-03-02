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
	return convert.WithErr[schema.ProblemSet](
		service.ProblemSet(c).CreateProblemSet(*request.Body),
	)
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
	return convert.WithErr[schema.ProblemSet](
		service.ProblemSet(c).GetCurrentProblemSet(),
	)
}

// Update Problem Set
// (PATCH /domains/{domain}/problem_sets/{problemSet})
func (s *Api) UpdateProblemSet(
	c *fiber.Ctx,
	request schema.UpdateProblemSetRequestObject,
) (any, error) {
	return convert.WithErr[schema.ProblemSet](
		service.ProblemSet(c).UpdateProblemSet(*request.Body),
	)
}

// List Problems In Problem Set
// (GET /domains/{domain}/problem_sets/{problemSet}/problems)
func (s *Api) ListProblemsInProblemSet(
	c *fiber.Ctx,
	request schema.ListProblemsInProblemSetRequestObject,
) (any, error) {
	return service.ProblemSet(c).ListProblemsInProblemSet()
}

// Add Problem In Problem Set
// (POST /domains/{domain}/problem_sets/{problemSet}/problems)
func (s *Api) AddProblemInProblemSet(
	c *fiber.Ctx,
	request schema.AddProblemInProblemSetRequestObject,
) (any, error) {
	return convert.WithErr[schema.ProblemSet](
		service.Problem(c).AddProblemInProblemSet(*request.Body),
	)
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
	return convert.WithErr[schema.Problem](
		service.Problem(c).GetCurrentProblem(),
	)
}

// Update Problem In Problem Set
// (PATCH /domains/{domain}/problem_sets/{problemSet}/problems/{problem})
func (s *Api) UpdateProblemInProblemSet(
	c *fiber.Ctx,
	request schema.UpdateProblemInProblemSetRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Submit Solution To Problem In Problem Set
// (POST /domains/{domain}/problem_sets/{problemSet}/problems/{problem}/submit)
func (s *Api) SubmitSolutionToProblemInProblemSet(
	c *fiber.Ctx,
	request schema.SubmitSolutionToProblemInProblemSetRequestObject,
) (any, error) {
	return convert.WithErr[schema.Record](
		service.Record(c).SubmitInProblemSet(),
	)
}
