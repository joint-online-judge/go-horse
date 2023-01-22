package handlers

import (
	"context"

	"github.com/joint-online-judge/go-horse/schemas"
)

// List Records In Domain
// (GET /domains/{domain}/records)
func (s *ApiV1) ListRecordsInDomain(
	ctx context.Context,
	request schemas.ListRecordsInDomainRequestObject,
) (any, error) {
	return nil, nil
}

// Get Record
// (GET /domains/{domain}/records/{record})
func (s *ApiV1) GetRecord(
	ctx context.Context,
	request schemas.GetRecordRequestObject,
) (any, error) {
	return nil, nil
}

// Submit Case By Judger
// (PUT /domains/{domain}/records/{record}/cases/{index}/judge)
func (s *ApiV1) SubmitCaseByJudger(
	ctx context.Context,
	request schemas.SubmitCaseByJudgerRequestObject,
) (any, error) {
	return nil, nil
}

// Submit Record By Judger
// (PUT /domains/{domain}/records/{record}/judge)
func (s *ApiV1) SubmitRecordByJudger(
	ctx context.Context,
	request schemas.SubmitRecordByJudgerRequestObject,
) (any, error) {
	return nil, nil
}

// Claim Record By Judger
// (POST /domains/{domain}/records/{record}/judge/claim)
func (s *ApiV1) ClaimRecordByJudger(
	ctx context.Context,
	request schemas.ClaimRecordByJudgerRequestObject,
) (any, error) {
	return nil, nil
}
