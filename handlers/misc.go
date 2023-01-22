package handlers

import (
	"context"

	"github.com/joint-online-judge/go-horse/schemas"
)

// Jwt Decoded
// (GET /jwt_decoded)
func (s *ApiV1) JwtDecoded(
	ctx context.Context,
	request schemas.JwtDecodedRequestObject,
) (any, error) {
	return nil, nil
}

// Test Error Report
// (GET /test/report)
func (s *ApiV1) TestErrorReport(
	ctx context.Context,
	request schemas.TestErrorReportRequestObject,
) (any, error) {
	return nil, nil
}

// Version
// (GET /version)
func (s *ApiV1) Version(ctx context.Context, request schemas.VersionRequestObject) (any, error) {
	return nil, nil
}
