package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schemas"
)

// List Records In Domain
// (GET /domains/{domain}/records)
func (s *Api) ListRecordsInDomain(
	c *fiber.Ctx,
	request schemas.ListRecordsInDomainRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Get Record
// (GET /domains/{domain}/records/{record})
func (s *Api) GetRecord(
	c *fiber.Ctx,
	request schemas.GetRecordRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Submit Case By Judger
// (PUT /domains/{domain}/records/{record}/cases/{index}/judge)
func (s *Api) SubmitCaseByJudger(
	c *fiber.Ctx,
	request schemas.SubmitCaseByJudgerRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Submit Record By Judger
// (PUT /domains/{domain}/records/{record}/judge)
func (s *Api) SubmitRecordByJudger(
	c *fiber.Ctx,
	request schemas.SubmitRecordByJudgerRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Claim Record By Judger
// (POST /domains/{domain}/records/{record}/judge/claim)
func (s *Api) ClaimRecordByJudger(
	c *fiber.Ctx,
	request schemas.ClaimRecordByJudgerRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}
