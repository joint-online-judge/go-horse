package v1

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/pkg/config"
)

// Jwt Decoded
// (GET /jwt_decoded)
func (s *Api) JwtDecoded(
	c *fiber.Ctx,
	request schema.JwtDecodedRequestObject,
) (any, error) {
	return schema.JWT(c), nil
}

// Test Error Report
// (GET /test/report)
func (s *Api) TestErrorReport(
	c *fiber.Ctx,
	request schema.TestErrorReportRequestObject,
) (any, error) {
	panic("test error report")
	// return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Version
// (GET /version)
func (s *Api) Version(
	c *fiber.Ctx,
	request schema.VersionRequestObject,
) (any, error) {
	return schema.NewNonStandardResp(schema.Version{
		Git:     fmt.Sprintf("%s %s", config.GitCommit, config.BuiltAt),
		Version: config.Version,
	}), nil
}
