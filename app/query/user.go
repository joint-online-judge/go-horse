package query

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/schema"
)

func GetUser(c *fiber.Ctx, user string) (u schema.User, err error) {
	var userId uuid.UUID
	if user == "me" {
		userId = schema.JWTUser(c).ID
	} else {
		userId, err = uuid.Parse(user)
		if err != nil {
			return
		}
	}
	model := model.User{ID: userId}
	u, err = GetObj[schema.User](&model)
	return
}
