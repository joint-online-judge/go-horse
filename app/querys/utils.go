package querys

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func ConvertTo[DstType any](src any) (dst DstType, err error) {
	b, err := json.Marshal(src)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &dst)
	return
}

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

func GetObj[Model, Dest any](model *Model) (dest Dest, err error) {
	err = DB.Where(model).First(&dest).Error
	return
}

func SaveObj(model any) error {
	return DB.Save(model).Error
}

func CreateObj[Schema any](model any) (schema Schema, err error) {
	if err = DB.Create(model).Error; err != nil {
		return
	}
	return ConvertTo[Schema](model)
}
