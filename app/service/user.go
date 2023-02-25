package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/query"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/pkg/convert"
)

type userImpl struct {
	c *fiber.Ctx
}

func User(c *fiber.Ctx) *userImpl {
	return &userImpl{
		c: c,
	}
}

func (s *userImpl) GetUser(user string) (userModel *model.User, err error) {
	var userId uuid.UUID
	if user == "me" {
		userId = Auth(s.c).JWTUser().ID
	} else {
		userId, err = uuid.Parse(user)
		if err != nil {
			return
		}
	}
	userModel = &model.User{ID: userId}
	err = query.GetObj(&userModel)
	return
}

func (s *userImpl) GetCurrentUser() (*schema.UserDetail, error) {
	userModel, err := User(s.c).GetUser("me")
	if err != nil {
		return nil, schema.NewBizError(schema.UserNotFoundError)
	}
	userDetail, err := convert.To[schema.UserDetail](userModel)
	return &userDetail, err
}
