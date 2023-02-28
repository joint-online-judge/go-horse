package query

import (
	"github.com/joint-online-judge/go-horse/app/model"
	"gorm.io/gorm"
)

func ListProblems(
	db *gorm.DB,
	domain *model.Domain, problemSet *model.ProblemSet,
	includeHidden bool,
) *gorm.DB {
	statement := db.Model(model.Problem{})
	if domain != nil {
		statement = statement.Where("domain_id = ?", domain.Id)
	}
	if problemSet != nil {
		link := model.ProblemProblemSetLink{ProblemSetId: problemSet.Id}
		statement = statement.Where(
			"id in (?)", db.Model(&link).Select("problem_id").Where(&link),
		)
	}
	if !includeHidden {
		statement = statement.Not("hidden = ?", true)
	}
	// TODO: add latest record
	return statement
}
