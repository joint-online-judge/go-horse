package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/schemas"
)

// List Records In Domain
// (GET /domains/{domain}/records)
func (s *ApiV1) ListRecordsInDomain(
	c *fiber.Ctx,
	request schemas.ListRecordsInDomainRequestObject,
) (any, error) {
	return nil, nil
}

// Get Record
// (GET /domains/{domain}/records/{record})
func (s *ApiV1) GetRecord(
	c *fiber.Ctx,
	request schemas.GetRecordRequestObject,
) (any, error) {
	return nil, nil
}

// Submit Case By Judger
// (PUT /domains/{domain}/records/{record}/cases/{index}/judge)
func (s *ApiV1) SubmitCaseByJudger(
	c *fiber.Ctx,
	request schemas.SubmitCaseByJudgerRequestObject,
) (any, error) {
	return nil, nil
}

// Submit Record By Judger
// (PUT /domains/{domain}/records/{record}/judge)
func (s *ApiV1) SubmitRecordByJudger(
	c *fiber.Ctx,
	request schemas.SubmitRecordByJudgerRequestObject,
) (any, error) {
	return nil, nil
}

// Claim Record By Judger
// (POST /domains/{domain}/records/{record}/judge/claim)
func (s *ApiV1) ClaimRecordByJudger(
	c *fiber.Ctx,
	request schemas.ClaimRecordByJudgerRequestObject,
) (any, error) {
	return nil, nil
}
