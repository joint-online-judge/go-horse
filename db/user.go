package db

import (
	"github.com/joint-online-judge/go-horse/models"
	"github.com/joint-online-judge/go-horse/schemas"
)

func GetUser() (schemas.User, error) {
	var user schemas.User
	u := DB.Model(models.User{})
	err := u.First(&user).Error
	return user, err
}
