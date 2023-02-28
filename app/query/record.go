package query

import (
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/model"
	"gorm.io/gorm"
)

func ListRecords(
	db *gorm.DB, domain *model.Domain, problemSetId *uuid.UUID,
	problemId *uuid.UUID, submitterId *uuid.UUID,
) *gorm.DB {
	statement := db.Model(model.Record{}).
		Where("domain_id = ?", domain.Id)
	if problemSetId != nil {
		statement = statement.Where("problem_set_id = ?", problemSetId)
	}
	if problemId != nil {
		statement = statement.Where("problem_id = ?", problemId)
	}
	if submitterId != nil {
		statement = statement.Where("committer_id = ?", submitterId)
	}
	return statement
}
