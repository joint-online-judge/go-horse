package querys

import (
	"github.com/joint-online-judge/go-horse/app/models"
	"github.com/joint-online-judge/go-horse/app/schemas"
)

func GetUser() (schemas.User, error) {
	var user schemas.User
	u := DB.Model(models.User{})
	err := u.Limit(1).Find(&user).Error
	return user, err
}
