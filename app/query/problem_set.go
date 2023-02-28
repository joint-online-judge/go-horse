package query

import (
	"github.com/joint-online-judge/go-horse/app/model"
	"gorm.io/gorm"
)

func ListProblemSets(
	db *gorm.DB, domain *model.Domain, includeHidden bool,
) *gorm.DB {
	statement := db.Model(model.ProblemSet{}).
		Where("domain_id = ?", domain.Id)
	if !includeHidden {
		statement = statement.Not("hidden = ?", true)
	}
	return statement
}
