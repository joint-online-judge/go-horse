package query

import (
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/schema"
	"gorm.io/gorm"
)

func GetRecord(db *gorm.DB, domain *model.Domain, record string) (
	recordModel model.Record, err error,
) {
	recordId, err := uuid.Parse(record)
	if err != nil {
		return
	}
	recordModel.Id = recordId
	recordModel.DomainId = domain.Id
	err = db.First(&recordModel).Error
	return
}

func ListRecords(
	db *gorm.DB, domain *model.Domain, pagination schema.Pagination,
) ([]schema.RecordListDetail, int64, error) {
	statement := db.Model(model.Record{}).
		Where("domain_id = ?", domain.Id)
	return ListObjs[schema.RecordListDetail](
		statement, pagination,
	)
}
