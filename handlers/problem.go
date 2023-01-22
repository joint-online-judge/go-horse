package handlers

import (
	"context"

	"github.com/joint-online-judge/go-horse/schemas"
)

// List Problems
// (GET /domains/{domain}/problems)
func (s *ApiV1) ListProblems(
	ctx context.Context,
	request schemas.ListProblemsRequestObject,
) (any, error) {
	return nil, nil
}

// Create Problem
// (POST /domains/{domain}/problems)
func (s *ApiV1) CreateProblem(
	ctx context.Context,
	request schemas.CreateProblemRequestObject,
) (any, error) {
	return nil, nil
}

// Clone Problem
// (POST /domains/{domain}/problems/clone)
func (s *ApiV1) CloneProblem(
	ctx context.Context,
	request schemas.CloneProblemRequestObject,
) (any, error) {
	return nil, nil
}

// Delete Problem
// (DELETE /domains/{domain}/problems/{problem})
func (s *ApiV1) DeleteProblem(
	ctx context.Context,
	request schemas.DeleteProblemRequestObject,
) (any, error) {
	return nil, nil
}

// Get Problem
// (GET /domains/{domain}/problems/{problem})
func (s *ApiV1) GetProblem(
	ctx context.Context,
	request schemas.GetProblemRequestObject,
) (any, error) {
	return nil, nil
}

// Update Problem
// (PATCH /domains/{domain}/problems/{problem})
func (s *ApiV1) UpdateProblem(
	ctx context.Context,
	request schemas.UpdateProblemRequestObject,
) (any, error) {
	return nil, nil
}

// Submit Solution To Problem
// (POST /domains/{domain}/problems/{problem})
func (s *ApiV1) SubmitSolutionToProblem(
	ctx context.Context,
	request schemas.SubmitSolutionToProblemRequestObject,
) (any, error) {
	return nil, nil
}
