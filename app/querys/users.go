package querys

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/models"
	"github.com/joint-online-judge/go-horse/app/schemas"
)

func GetUser(c *fiber.Ctx, user string) (schema schemas.User, err error) {
	var userId uuid.UUID
	if user == "me" {
		userId = schemas.JWTUser(c).Id
	} else {
		userId, err = uuid.Parse(user)
		if err != nil {
			return
		}
	}
	model := models.User{Base: models.Base{ID: userId}}
	schema, err = GetObj[schemas.User](&model)
	return
}
