package db

import (
	"github.com/pkg/errors"
)

func ListObjs[Model any, Types any]() ([]Types, int64, error) {
	var types []Types
	var model Model
	u := DB.Model(model)
	var count int64
	if err := u.Count(&count).Error; err != nil {
		return nil, 0, errors.Wrapf(err, "failed get %T count", types)
	}
	err := u.Find(&types).Error
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}
	return types, count, nil
}
