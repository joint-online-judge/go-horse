package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schemas"
	"github.com/joint-online-judge/go-horse/pkg/configs"
)

// Jwt Decoded
// (GET /jwt_decoded)
func (s *ApiV1) JwtDecoded(
	c *fiber.Ctx,
	request schemas.JwtDecodedRequestObject,
) (any, error) {
	return schemas.JWT(c), nil
}

// Test Error Report
// (GET /test/report)
func (s *ApiV1) TestErrorReport(
	c *fiber.Ctx,
	request schemas.TestErrorReportRequestObject,
) (any, error) {
	panic("test error report")
	// return nil, nil
}

// Version
// (GET /version)
func (s *ApiV1) Version(c *fiber.Ctx, request schemas.VersionRequestObject) (any, error) {
	return schemas.NonStandardResp{Data: schemas.Version{
		Git:     configs.GitCommit,
		Version: configs.Version,
	}}, nil
}
