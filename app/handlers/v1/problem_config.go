package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schemas"
)

// List Problem Config Commits
// (GET /domains/{domain}/problems/{problem}/configs)
func (s *Api) ListProblemConfigCommits(
	c *fiber.Ctx,
	request schemas.ListProblemConfigCommitsRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Update Problem Config By Archive
// (POST /domains/{domain}/problems/{problem}/configs)
func (s *Api) UpdateProblemConfigByArchive(
	c *fiber.Ctx,
	request schemas.UpdateProblemConfigByArchiveRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Update Problem Config Json
// (POST /domains/{domain}/problems/{problem}/configs/json)
func (s *Api) UpdateProblemConfigJson(
	c *fiber.Ctx,
	request schemas.UpdateProblemConfigJsonRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Diff Problem Config Default Branch
// (GET /domains/{domain}/problems/{problem}/configs/latest/diff)
func (s *Api) DiffProblemConfigDefaultBranch(
	c *fiber.Ctx,
	request schemas.DiffProblemConfigDefaultBranchRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// List Latest Problem Config Objects Under A Given Prefix
// (GET /domains/{domain}/problems/{problem}/configs/latest/ls)
func (s *Api) ListLatestProblemConfigObjectsUnderAGivenPrefix(
	c *fiber.Ctx,
	request schemas.ListLatestProblemConfigObjectsUnderAGivenPrefixRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Download Problem Config Archive
// (GET /domains/{domain}/problems/{problem}/configs/{config})
func (s *Api) DownloadProblemConfigArchive(
	c *fiber.Ctx,
	request schemas.DownloadProblemConfigArchiveRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Get Problem Config Json
// (GET /domains/{domain}/problems/{problem}/configs/{config}/json)
func (s *Api) GetProblemConfigJson(
	c *fiber.Ctx,
	request schemas.GetProblemConfigJsonRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}
