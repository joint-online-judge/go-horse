package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/query"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/pkg/convert"
	"github.com/sirupsen/logrus"
)

type userImpl struct {
	c *fiber.Ctx
}

func User(c *fiber.Ctx) *userImpl {
	return &userImpl{
		c: c,
	}
}

func (s *userImpl) GetUser(user string) (userModel model.User, err error) {
	var userId uuid.UUID
	if user == "me" {
		userId = Auth(s.c).JWTUser().ID
	} else {
		userId, err = uuid.Parse(user)
		if err != nil {
			return
		}
	}
	userModel = model.User{ID: userId}
	err = query.GetObj(&userModel)
	return
}

func (s *userImpl) GetCurrentUser() (model.User, error) {
	return s.GetUser("me")
}

func (s *userImpl) UpdateCurrentUser(userEdit schema.UserEdit) (
	u model.User, err error,
) {
	u, err = s.GetCurrentUser()
	if err != nil {
		return
	}
	if err = convert.Update(&u, userEdit); err != nil {
		return
	}
	logrus.Infof("update user to: %+v", u)
	err = query.SaveObj(&u)
	return
}

func (s *userImpl) ChangePassword(userResetPassword schema.UserResetPassword) (
	u model.User, err error,
) {
	u, err = s.GetCurrentUser()
	if err != nil {
		return
	}
	if (userResetPassword.CurrentPassword == nil && u.HashedPassword != "") ||
		!schema.VerifyPassword(
			*userResetPassword.CurrentPassword,
			u.HashedPassword,
		) {
		err = schema.NewBizError(
			schema.UsernamePasswordError,
			"incorrect password",
		)
		return
	}
	u.HashedPassword, err = schema.HashPassword(userResetPassword.NewPassword)
	if err != nil {
		return
	}
	err = query.SaveObj(&u)
	return
}
