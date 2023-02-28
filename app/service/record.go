package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func (s *recordImpl) GetRecord(
	domain *model.Domain, record string,
) (recordModel model.Record, err error) {
	recordId, err := uuid.Parse(record)
	if err != nil {
		return
	}
	recordModel.Id = recordId
	recordModel.DomainId = domain.Id
	err = db.First(&recordModel).Error
	return
}

func (s *recordImpl) ListRecords(
	params schema.ListRecordsInDomainParams,
) (resp schema.ListResp[schema.RecordListDetail], err error) {
	domain, err := Domain(s.c).GetCurrentDomain()
	if err != nil {
		return
	}
	objs, count, err := ListObjs[schema.RecordListDetail](
		query.ListRecords(
			db, domain, params.ProblemSet, params.Problem,
			params.SubmitterId,
		), params.Pagination,
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
