package service

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/query"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

func (s *recordImpl) submitImpl(body *multipart.Reader, inProblemSet bool) (
	record model.Record, err error,
) {
	var form *multipart.Form
	if form, err = s.c.MultipartForm(); err != nil {
		return
	}
	logrus.Debugf("form: %+v", form)
	var problemSet *model.ProblemSet
	userLatestRecord := model.UserLatestRecord{}
	domain, err := Domain(s.c).GetCurrentDomain()
	if err != nil {
		return
	}
	record.DomainId = domain.Id
	problem, err := Problem(s.c).GetCurrentProblem()
	if err != nil {
		return
	}
	record.ProblemId = &problem.Id
	userLatestRecord.ProblemId = problem.Id
	if inProblemSet {
		problemSet, err = ProblemSet(s.c).GetCurrentProblemSet()
		if err != nil {
			return
		}
		record.ProblemSetId = &problemSet.Id
		userLatestRecord.ProblemSetId = &problemSet.Id
	}
	problemConfig, err := ProblemConfig(s.c).GetLatestproblemConfig(problem)
	if err != nil {
		return
	}
	record.ProblemConfigId = &problemConfig.Id
	user := Auth(s.c).JWTUser()
	record.CommitterId = &user.Id
	userLatestRecord.UserId = user.Id
	// TODO: check language
	languages, ok := form.Value["language"]
	if !ok || len(languages) != 1 {
		err = fiber.ErrUnprocessableEntity
		return
	}
	record.Language = languages[0]
	// TODO: upload file
	problemSet.NumSubmit += 1
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := db.Save(&record).Error; err != nil {
			return err
		}
		if err := db.Save(problemSet).Error; err != nil {
			return err
		}
		// TODO: save latest record in cache
		result := db.Limit(1).Find(&userLatestRecord)
		if result.RowsAffected == 1 {
			if err := result.Update("record_id", record.Id).Error; err != nil {
				return err
			}
		} else {
			userLatestRecord.RecordId = record.Id
			if err := db.Save(&userLatestRecord).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return
}

func (s *recordImpl) SubmitInProblemSet(
	body *multipart.Reader,
) (model.Record, error) {
	return s.submitImpl(body, true)
}

func (s *recordImpl) Submit(body *multipart.Reader) (model.Record, error) {
	return s.submitImpl(body, false)
}
