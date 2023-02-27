package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/query"
	"github.com/joint-online-judge/go-horse/app/schema"
)

type problemsetImpl struct {
	c *fiber.Ctx
}

func ProblemSet(c *fiber.Ctx) *problemsetImpl {
	return &problemsetImpl{
		c: c,
	}
}

func (s *problemsetImpl) ListProblemSets(
	params schema.ListProblemSetsParams,
) (resp schema.ListResp[schema.ProblemSet], err error) {
	domain, err := Domain(s.c).GetCurrentDomain()
	if err != nil {
		return
	}
	objs, count, err := query.ListProblemSets(
		domain, params.Pagination, false,
	)
	return schema.NewListResp(count, objs), err
}

func (s *problemsetImpl) CreateProblemSet(
	problemSetCreate schema.ProblemSetCreate,
) (resp model.ProblemSet, err error) {
	domain, err := Domain(s.c).GetCurrentDomain()
	if err != nil {
		return
	}
	user := Auth(s.c).JWTUser()
	return query.CreateProblemSet(problemSetCreate, domain, user)
}

func (s *problemsetImpl) GetCurrentProblemSet() (*model.ProblemSet, error) {
	problemSet := s.c.Locals("problemSet")
	if problemSet == nil {
		return nil, schema.NewBizError(
			schema.ProblemSetNotFoundError, "cannot find current problemSet",
		)
	}
	return problemSet.(*model.ProblemSet), nil
}
