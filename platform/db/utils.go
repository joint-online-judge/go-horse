package db

import (
	"github.com/pkg/errors"
)

func ListObjs[Model, Schema any]() ([]Schema, int64, error) {
	var schemas []Schema
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

func GetObj[Model, Schema any](model *Model) (Schema, error) {
	var schema Schema
	err := DB.Where(model).First(&schema).Error
	return schema, err
}

func SaveObj(model any) error {
	return DB.Save(model).Error
}
