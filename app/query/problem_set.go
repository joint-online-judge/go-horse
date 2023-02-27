package query

import (
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/schema"
)

func ListProblemSets(
	domain *model.Domain, pagination schema.Pagination, includeHidden bool,
) ([]schema.ProblemSet, int64, error) {
	statement := db.Model(model.ProblemSet{}).
		Where("problem_sets.domain_id = ?", domain.Id)
	if !includeHidden {
		statement = statement.Not("problem_sets.hidden = ?", true)
	}
	return ListObjs[schema.ProblemSet](
		statement, pagination,
	)
}

func GetProblemSet(domain *model.Domain, problemSet string) (
	problemSetModel model.ProblemSet, err error,
) {
	var query model.ProblemSet
	if problemSetId, err := uuid.Parse(problemSet); err != nil {
		query.Url = problemSet
	} else {
		query.Id = problemSetId
	}
	query.DomainId = domain.Id
	problemSetModel, err = GetObjTo[model.ProblemSet](&query)
	if err != nil {
		return
	}
	return
}
