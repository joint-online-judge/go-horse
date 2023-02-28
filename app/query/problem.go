package query

import (
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/model"
	"gorm.io/gorm"
)

func GetProblemInProblemSet(
	db *gorm.DB, domain *model.Domain, problemSet *model.ProblemSet,
	problem string,
) (problemModel model.Problem, err error) {
	if problemId, err := uuid.Parse(problem); err != nil {
		problemModel.Url = problem
	} else {
		problemModel.Id = problemId
	}
	problemModel.DomainId = domain.Id
	err = db.First(&problemModel).Error
	if err != nil {
		return
	}
	var link model.ProblemProblemSetLink
	link.ProblemId = problemModel.Id
	link.ProblemSetId = problemSet.Id
	err = db.First(&link).Error
	return
}

func ListProblems(
	db *gorm.DB, domain *model.Domain, includeHidden bool,
) *gorm.DB {
	statement := db.Model(model.Problem{}).
		Where("domain_id = ?", domain.Id)
	if !includeHidden {
		statement = statement.Not("hidden = ?", true)
	}
	// TODO: add latest record
	return statement
}
