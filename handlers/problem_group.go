package handlers

import (
	"context"

	"github.com/joint-online-judge/go-horse/schemas"
)

// List Problem Groups
// (GET /problem_groups)
func (s *ApiV1) ListProblemGroups(
	ctx context.Context,
	request schemas.ListProblemGroupsRequestObject,
) (any, error) {
	return nil, nil
}
