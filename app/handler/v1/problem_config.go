package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/app/service"
	"github.com/joint-online-judge/go-horse/pkg/convert"
)

// List Problem Config Commits
// (GET /domains/{domain}/problems/{problem}/configs)
func (s *Api) ListProblemConfigCommits(
	c *fiber.Ctx,
	request schema.ListProblemConfigCommitsRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Update Problem Config By Archive
// (POST /domains/{domain}/problems/{problem}/configs)
func (s *Api) UpdateProblemConfigByArchive(
	c *fiber.Ctx,
	request schema.UpdateProblemConfigByArchiveRequestObject,
) (any, error) {
	return convert.WithErr[schema.ProblemConfigDetail](
		service.ProblemConfig(c).UpdateProblemConfigByArchive(),
	)
}

// Update Problem Config Json
// (POST /domains/{domain}/problems/{problem}/configs/json)
func (s *Api) UpdateProblemConfigJson(
	c *fiber.Ctx,
	request schema.UpdateProblemConfigJsonRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Diff Problem Config Default Branch
// (GET /domains/{domain}/problems/{problem}/configs/latest/diff)
func (s *Api) DiffProblemConfigDefaultBranch(
	c *fiber.Ctx,
	request schema.DiffProblemConfigDefaultBranchRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// List Latest Problem Config Objects Under A Given Prefix
// (GET /domains/{domain}/problems/{problem}/configs/latest/ls)
func (s *Api) ListLatestProblemConfigObjectsUnderAGivenPrefix(
	c *fiber.Ctx,
	request schema.ListLatestProblemConfigObjectsUnderAGivenPrefixRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Download Problem Config Archive
// (GET /domains/{domain}/problems/{problem}/configs/{config})
func (s *Api) DownloadProblemConfigArchive(
	c *fiber.Ctx,
	request schema.DownloadProblemConfigArchiveRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Get Problem Config Json
// (GET /domains/{domain}/problems/{problem}/configs/{config}/json)
func (s *Api) GetProblemConfigJson(
	c *fiber.Ctx,
	request schema.GetProblemConfigJsonRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}
