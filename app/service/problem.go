package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/query"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/pkg/convert"
	"gorm.io/gorm"
)

type problemImpl struct {
	c *fiber.Ctx
}

func Problem(c *fiber.Ctx) *problemImpl {
	return &problemImpl{
		c: c,
	}
}

func (s *problemImpl) GetProblem(
	domain *model.Domain, problemSet *model.ProblemSet,
	problem string,
) (problemModel model.Problem, err error) {
	if problemId, err := uuid.Parse(problem); err != nil {
		problemModel.Url = problem
	} else {
		problemModel.Id = problemId
	}
	problemModel.DomainId = domain.Id
	if problemSet == nil {
		err = db.First(&problemModel).Error
	} else {
		err = db.First(&problemModel).
			Joins("JOIN problem_problem_set_links AS links ON "+
				"links.problem_id = problems.id").
			Where("links.problem_set_id = ?", problemSet.Id).
			Error
	}
	return
}

func (s *problemImpl) ListProblems(
	params schema.ListProblemsParams,
) (resp schema.ListResp[schema.ProblemWithLatestRecord], err error) {
	domain, err := Domain(s.c).GetCurrentDomain()
	if err != nil {
		return
	}
	objs, count, err := ListObjs[schema.ProblemWithLatestRecord](
		query.ListProblems(
			db, domain, nil, false,
		), params.Pagination,
	)
	return schema.NewListResp(count, objs), err
}

func (s *problemImpl) CreateProblem(
	problemCreate schema.ProblemCreate,
) (problem model.Problem, err error) {
	domain, err := Domain(s.c).GetCurrentDomain()
	if err != nil {
		return
	}
	user := Auth(s.c).JWTUser()
	owner := model.User{Id: user.Id}
	err = convert.Update(&problem, problemCreate)
	if err != nil {
		return
	}
	problem.Domain = *domain
	problem.Owner = &owner
	err = db.Transaction(func(tx *gorm.DB) error {
		problemGroup := model.ProblemGroup{}
		if err := tx.Create(&problemGroup).Error; err != nil {
			return err
		}
		problem.ProblemGroupId = &problemGroup.Id
		if err := tx.Save(&problem).Error; err != nil {
			return err
		}
		return nil
	})
	return
}

func (s *problemImpl) GetCurrentProblem() (*model.Problem, error) {
	problem := s.c.Locals("problem")
	if problem == nil {
		return nil, schema.NewBizError(
			schema.ProblemNotFoundError, "cannot find current problem",
		)
	}
	return problem.(*model.Problem), nil
}

func (s *problemImpl) AddProblemInProblemSet(
	problemSetAddProblem schema.ProblemSetAddProblem,
) (problemSet *model.ProblemSet, err error) {
	domain, err := Domain(s.c).GetCurrentDomain()
	if err != nil {
		return
	}
	problemSet, err = ProblemSet(s.c).GetCurrentProblemSet()
	if err != nil {
		return
	}
	// TODO: check hidden
	problem, err := s.GetProblem(
		domain, problemSet, problemSetAddProblem.Problem,
	)
	if err != nil {
		return
	}
	err = db.Create(&model.ProblemProblemSetLink{
		ProblemId: problem.Id, ProblemSetId: problemSet.Id,
		Position: int32(*problemSetAddProblem.Position),
	}).Error
	return
}
