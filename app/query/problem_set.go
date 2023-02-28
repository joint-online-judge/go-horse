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

func ListProblemsInProblemSet(
	db *gorm.DB, problemSet *model.ProblemSet, problem *model.Problem,
) *gorm.DB {
	link := model.ProblemProblemSetLink{ProblemSetId: problemSet.Id}
	statement := db.Model(model.Problem{}).Where(
		"id in (?)", db.Model(&link).Select("problem_id").Where(&link),
	)
	// TODO: add latest record
	return statement
}
