package handlers

import (
	"context"

	"github.com/joint-online-judge/go-horse/schemas"
)

// List Problem Config Commits
// (GET /domains/{domain}/problems/{problem}/configs)
func (s *ApiV1) ListProblemConfigCommits(
	ctx context.Context,
	request schemas.ListProblemConfigCommitsRequestObject,
) (any, error) {
	return nil, nil
}

// Update Problem Config By Archive
// (POST /domains/{domain}/problems/{problem}/configs)
func (s *ApiV1) UpdateProblemConfigByArchive(
	ctx context.Context,
	request schemas.UpdateProblemConfigByArchiveRequestObject,
) (any, error) {
	return nil, nil
}

// Update Problem Config Json
// (POST /domains/{domain}/problems/{problem}/configs/json)
func (s *ApiV1) UpdateProblemConfigJson(
	ctx context.Context,
	request schemas.UpdateProblemConfigJsonRequestObject,
) (any, error) {
	return nil, nil
}

// Diff Problem Config Default Branch
// (GET /domains/{domain}/problems/{problem}/configs/latest/diff)
func (s *ApiV1) DiffProblemConfigDefaultBranch(
	ctx context.Context,
	request schemas.DiffProblemConfigDefaultBranchRequestObject,
) (any, error) {
	return nil, nil
}

// List Latest Problem Config Objects Under A Given Prefix
// (GET /domains/{domain}/problems/{problem}/configs/latest/ls)
func (s *ApiV1) ListLatestProblemConfigObjectsUnderAGivenPrefix(
	ctx context.Context,
	request schemas.ListLatestProblemConfigObjectsUnderAGivenPrefixRequestObject,
) (any, error) {
	return nil, nil
}

// Download Problem Config Archive
// (GET /domains/{domain}/problems/{problem}/configs/{config})
func (s *ApiV1) DownloadProblemConfigArchive(
	ctx context.Context,
	request schemas.DownloadProblemConfigArchiveRequestObject,
) (any, error) {
	return nil, nil
}

// Get Problem Config Json
// (GET /domains/{domain}/problems/{problem}/configs/{config}/json)
func (s *ApiV1) GetProblemConfigJson(
	ctx context.Context,
	request schemas.GetProblemConfigJsonRequestObject,
) (any, error) {
	return nil, nil
}
