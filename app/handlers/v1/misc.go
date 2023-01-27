package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schemas"
	"github.com/joint-online-judge/go-horse/pkg/configs"
)

// Jwt Decoded
// (GET /jwt_decoded)
func (s *Api) JwtDecoded(
	c *fiber.Ctx,
	request schemas.JwtDecodedRequestObject,
) (any, error) {
	return schemas.JWT(c), nil
}

// Test Error Report
// (GET /test/report)
func (s *Api) TestErrorReport(
	c *fiber.Ctx,
	request schemas.TestErrorReportRequestObject,
) (any, error) {
	panic("test error report")
	// return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Version
// (GET /version)
func (s *Api) Version(
	c *fiber.Ctx,
	request schemas.VersionRequestObject,
) (any, error) {
	return schemas.NewNonStandardResp(schemas.Version{
		Git:     configs.GitCommit,
		Version: configs.Version,
	}), nil
}
