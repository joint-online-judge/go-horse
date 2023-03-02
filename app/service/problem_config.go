package service

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/platform/storage"
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
) (problemConfig model.ProblemConfig, err error) {
	problemConfig.ProblemId = &problem.Id
	result := db.Limit(1).Order("updated_at desc").Find(&problemConfig)
	if result.Error != nil {
		err = result.Error
	} else if result.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}
	return
}

func (s *problemConfigImpl) UpdateProblemConfigByArchive() (
	problemConfig model.ProblemConfig, err error,
) {
	problem, err := Problem(s.c).GetCurrentProblem()
	if err != nil {
		return
	}
	user, err := User(s.c).GetCurrentUser()
	if err != nil {
		return
	}
	var form *multipart.Form
	if form, err = s.c.MultipartForm(); err != nil {
		return
	}
	formFiles, ok := form.File["file"]
	if !ok || len(formFiles) != 1 {
		err = fiber.ErrUnprocessableEntity
		return
	}
	problemConfig = model.ProblemConfig{
		ProblemId:   &problem.Id,
		CommitterId: &user.Id,
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&problemConfig).Error; err != nil {
			return err
		}
		return storage.PutProblemConfig("", formFiles[0])
	})
	return
}
