package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/query"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/pkg/convert"
)

type problemSetImpl struct {
	c *fiber.Ctx
}

func ProblemSet(c *fiber.Ctx) *problemSetImpl {
	return &problemSetImpl{
		c: c,
	}
}

func (s *problemSetImpl) GetProblemSet(
	domain *model.Domain, problemSet string,
) (problemSetModel model.ProblemSet, err error) {
	if problemSetId, err := uuid.Parse(problemSet); err != nil {
		problemSetModel.Url = problemSet
	} else {
		problemSetModel.Id = problemSetId
	}
	problemSetModel.DomainId = domain.Id
	err = db.First(&problemSetModel).Error
	return
}

func (s *problemSetImpl) ListProblemSets(
	params schema.ListProblemSetsParams,
) (resp schema.ListResp[schema.ProblemSet], err error) {
	domain, err := Domain(s.c).GetCurrentDomain()
	if err != nil {
		return
	}
	objs, count, err := ListObjs[schema.ProblemSet](
		query.ListProblemSets(db, domain, false), params.Pagination,
	)
	return schema.NewListResp(count, objs), err
}

func (s *problemSetImpl) CreateProblemSet(
	problemSetCreate schema.ProblemSetCreate,
) (problemSet model.ProblemSet, err error) {
	domain, err := Domain(s.c).GetCurrentDomain()
	if err != nil {
		return
	}
	user := Auth(s.c).JWTUser()
	owner := model.User{Id: user.Id}
	err = convert.Update(&problemSet, problemSetCreate)
	if err != nil {
		return
	}
	problemSet.Domain = *domain
	problemSet.Owner = &owner
	err = db.Save(&problemSet).Error
	return
}

func (s *problemSetImpl) GetCurrentProblemSet() (*model.ProblemSet, error) {
	problemSet := s.c.Locals("problemSet")
	if problemSet == nil {
		return nil, schema.NewBizError(
			schema.ProblemSetNotFoundError, "cannot find current problemSet",
		)
	}
	return problemSet.(*model.ProblemSet), nil
}

func (s *problemSetImpl) ListProblemsInProblemSet() (
	resp schema.ListResp[schema.ProblemSet], err error,
) {
	problemSet, err := s.GetCurrentProblemSet()
	if err != nil {
		return
	}
	objs, count, err := ListObjs[schema.ProblemSet](
		query.ListProblems(
			db, nil, problemSet, false,
		), schema.Pagination{},
	)
	return schema.NewListResp(count, objs), err
}

func (s *problemSetImpl) UpdateProblemSet(
	problemSetEdit schema.ProblemSetEdit,
) (problemSet *model.ProblemSet, err error) {
	if problemSet, err = s.GetCurrentProblemSet(); err != nil {
		return
	}
	if err = convert.Update(problemSet, problemSetEdit); err != nil {
		return
	}
	err = db.Save(problemSet).Error
	return
}
