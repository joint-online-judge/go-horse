package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/model"
	"gorm.io/gorm"
)

type problemConfigImpl struct {
	c *fiber.Ctx
}

func ProblemConfig(c *fiber.Ctx) *problemConfigImpl {
	return &problemConfigImpl{
		c: c,
	}
}

func (s *problemConfigImpl) GetLatestproblemConfig(
	problem *model.Problem,
) (problemConfigModel model.ProblemConfig, err error) {
	problemConfigModel.ProblemId = &problem.Id
	result := db.Limit(1).Order("updated_at desc").Find(&problemConfigModel)
	if result.Error != nil {
		err = result.Error
	} else if result.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}
	return
}
