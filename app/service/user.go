package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/query"
	"github.com/joint-online-judge/go-horse/app/schema"
)

type userImpl struct {
	c *fiber.Ctx
}

func User(c *fiber.Ctx) *userImpl {
	return &userImpl{
		c: c,
	}
}

func (s *userImpl) GetUser(user string) (u schema.User, err error) {
	var userId uuid.UUID
	if user == "me" {
		userId = Auth(s.c).JWTUser().ID
	} else {
		userId, err = uuid.Parse(user)
		if err != nil {
			return
		}
	}
	userModel := model.User{ID: userId}
	u, err = query.GetObj[schema.User](&userModel)
	return
}
