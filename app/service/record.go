package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/query"
	"github.com/joint-online-judge/go-horse/app/schema"
)

type recordImpl struct {
	c *fiber.Ctx
}

func Record(c *fiber.Ctx) *recordImpl {
	return &recordImpl{
		c: c,
	}
}

func (s *recordImpl) ListRecords(
	params schema.ListRecordsInDomainParams,
) (resp schema.ListResp[schema.RecordListDetail], err error) {
	domain, err := Domain(s.c).GetCurrentDomain()
	if err != nil {
		return
	}
	// TODO: more filter options
	objs, count, err := query.ListRecords(
		db, domain, params.Pagination,
	)
	return schema.NewListResp(count, objs), err
}

func (s *recordImpl) GetCurrentRecord() (*model.Record, error) {
	record := s.c.Locals("record")
	if record == nil {
		return nil, schema.NewBizError(
			schema.RecordNotFoundError, "cannot find current record",
		)
	}
	return record.(*model.Record), nil
}
