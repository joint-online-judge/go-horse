package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/app/service"
	"github.com/joint-online-judge/go-horse/pkg/convert"
)

// List Records In Domain
// (GET /domains/{domain}/records)
func (s *Api) ListRecordsInDomain(
	c *fiber.Ctx,
	request schema.ListRecordsInDomainRequestObject,
) (any, error) {
	return service.Record(c).ListRecords(request.Params)
}

// Get Record
// (GET /domains/{domain}/records/{record})
func (s *Api) GetRecord(
	c *fiber.Ctx,
	request schema.GetRecordRequestObject,
) (any, error) {
	return convert.WithErr[schema.Record](
		service.Record(c).GetCurrentRecord(),
	)
}

// Submit Case By Judger
// (PUT /domains/{domain}/records/{record}/cases/{index}/judge)
func (s *Api) SubmitCaseByJudger(
	c *fiber.Ctx,
	request schema.SubmitCaseByJudgerRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Submit Record By Judger
// (PUT /domains/{domain}/records/{record}/judge)
func (s *Api) SubmitRecordByJudger(
	c *fiber.Ctx,
	request schema.SubmitRecordByJudgerRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Claim Record By Judger
// (POST /domains/{domain}/records/{record}/judge/claim)
func (s *Api) ClaimRecordByJudger(
	c *fiber.Ctx,
	request schema.ClaimRecordByJudgerRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}
