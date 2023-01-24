package db

import (
	"github.com/pkg/errors"
)

func ListObjs[Model any, Schemas any]() ([]Schemas, int64, error) {
	var schemas []Schemas
	var model Model
	u := DB.Model(model)
	var count int64
	if err := u.Count(&count).Error; err != nil {
		return nil, 0, errors.Wrapf(err, "failed get %T count", schemas)
	}
	err := u.Find(&schemas).Error
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}
	return schemas, count, nil
}
