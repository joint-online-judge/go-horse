package query

import (
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/schema"
	"gorm.io/gorm"
)

func GetProblemSet(db *gorm.DB, domain *model.Domain, problemSet string) (
	problemSetModel model.ProblemSet, err error,
) {
	if problemSetId, err := uuid.Parse(problemSet); err != nil {
		problemSetModel.Url = problemSet
	} else {
		problemSetModel.Id = problemSetId
	}
	problemSetModel.DomainId = domain.Id
	err = db.First(&problemSetModel).Error
	return
}

func ListProblemSets(
	db *gorm.DB, domain *model.Domain, pagination schema.Pagination, includeHidden bool,
) ([]schema.ProblemSet, int64, error) {
	statement := db.Model(model.ProblemSet{}).
		Where("domain_id = ?", domain.Id)
	if !includeHidden {
		statement = statement.Not("hidden = ?", true)
	}
	return ListObjs[schema.ProblemSet](
		statement, pagination,
	)
}

func ListProblemsInProblemSet(
	db *gorm.DB, problemSet *model.ProblemSet, problem *model.Problem,
	pagination schema.Pagination,
) ([]schema.ProblemSet, int64, error) {
	link := model.ProblemProblemSetLink{ProblemSetId: problemSet.Id}
	statement := db.Model(model.Problem{}).Where(
		"id in (?)", db.Model(&link).Select("problem_id").Where(&link),
	)
	// TODO: add latest record
	return ListObjs[schema.ProblemSet](
		statement, pagination,
	)
}
