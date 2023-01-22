package handlers

import (
	"context"

	"github.com/joint-online-judge/go-horse/schemas"
)

// List Problem Sets
// (GET /domains/{domain}/problem_sets)
func (s *ApiV1) ListProblemSets(
	ctx context.Context,
	request schemas.ListProblemSetsRequestObject,
) (any, error) {
	return nil, nil
}

// Create Problem Set
// (POST /domains/{domain}/problem_sets)
func (s *ApiV1) CreateProblemSet(
	ctx context.Context,
	request schemas.CreateProblemSetRequestObject,
) (any, error) {
	return nil, nil
}

// Delete Problem Set
// (DELETE /domains/{domain}/problem_sets/{problemSet})
func (s *ApiV1) DeleteProblemSet(
	ctx context.Context,
	request schemas.DeleteProblemSetRequestObject,
) (any, error) {
	return nil, nil
}

// Get Problem Set
// (GET /domains/{domain}/problem_sets/{problemSet})
func (s *ApiV1) GetProblemSet(
	ctx context.Context,
	request schemas.GetProblemSetRequestObject,
) (any, error) {
	return nil, nil
}

// Update Problem Set
// (PATCH /domains/{domain}/problem_sets/{problemSet})
func (s *ApiV1) UpdateProblemSet(
	ctx context.Context,
	request schemas.UpdateProblemSetRequestObject,
) (any, error) {
	return nil, nil
}

// List Problems In Problem Set
// (GET /domains/{domain}/problem_sets/{problemSet}/problems)
func (s *ApiV1) ListProblemsInProblemSet(
	ctx context.Context,
	request schemas.ListProblemsInProblemSetRequestObject,
) (any, error) {
	return nil, nil
}

// Add Problem In Problem Set
// (POST /domains/{domain}/problem_sets/{problemSet}/problems)
func (s *ApiV1) AddProblemInProblemSet(
	ctx context.Context,
	request schemas.AddProblemInProblemSetRequestObject,
) (any, error) {
	return nil, nil
}

// Delete Problem In Problem Set
// (DELETE /domains/{domain}/problem_sets/{problemSet}/problems/{problem})
func (s *ApiV1) DeleteProblemInProblemSet(
	ctx context.Context,
	request schemas.DeleteProblemInProblemSetRequestObject,
) (any, error) {
	return nil, nil
}

// Get Problem In Problem Set
// (GET /domains/{domain}/problem_sets/{problemSet}/problems/{problem})
func (s *ApiV1) GetProblemInProblemSet(
	ctx context.Context,
	request schemas.GetProblemInProblemSetRequestObject,
) (any, error) {
	return nil, nil
}

// Update Problem In Problem Set
// (PATCH /domains/{domain}/problem_sets/{problemSet}/problems/{problem})
func (s *ApiV1) UpdateProblemInProblemSet(
	ctx context.Context,
	request schemas.UpdateProblemInProblemSetRequestObject,
) (any, error) {
	return nil, nil
}

// Submit Solution To Problem Set
// (POST /domains/{domain}/problem_sets/{problemSet}/problems/{problem}/submit)
func (s *ApiV1) SubmitSolutionToProblemSet(
	ctx context.Context,
	request schemas.SubmitSolutionToProblemSetRequestObject,
) (any, error) {
	return nil, nil
}
