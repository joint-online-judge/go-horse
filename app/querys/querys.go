package querys

import (
	"encoding/json"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/joint-online-judge/go-horse/app/schemas"
	"github.com/pkg/errors"
)

var DB *gorm.DB

func ConvertTo[DstType any](src any) (dst DstType, err error) {
	b, err := json.Marshal(src)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &dst)
	return
}

func ListObjs[Model, Schema any](pagination schemas.Pagination) ([]Schema, int64, error) {
	var schemas []Schema
	var model Model
	u := DB.Model(model)
	var count int64
	if err := u.Count(&count).Error; err != nil {
		return nil, 0, errors.Wrapf(err, "failed get %T count", schemas)
	}
	err := u.Scopes(Paginate(pagination)).Find(&schemas).Error
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}
	return schemas, count, nil
}

func GetObj[Dest any](modelPtr any) (dest Dest, err error) {
	err = DB.Where(modelPtr).First(&dest).Error
	return
}

func SaveObj(modelPtr any) error {
	return DB.Save(modelPtr).Error
}

func CreateObj[Schema any](modelPtr any) (schema Schema, err error) {
	if err = DB.Create(modelPtr).Error; err != nil {
		return
	}
	return ConvertTo[Schema](modelPtr)
}

func Paginate(pagination schemas.Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pagination.Offset == nil || *pagination.Offset < 0 {
			pagination.Offset = schemas.Pointer(0)
		}
		if pagination.Limit == nil || *pagination.Limit > 100 {
			pagination.Limit = schemas.Pointer(100)
		}
		if pagination.Ordering == nil {
			pagination.Ordering = schemas.Pointer("updated_at")
		}
		statement := db.Offset(*pagination.Offset).Limit(*pagination.Limit)
		for _, item := range strings.Split(*pagination.Ordering, ",") {
			name := item
			desc := false
			if strings.HasPrefix(name, "-") {
				name = strings.TrimLeft(item, "-")
				desc = true
			}
			statement = statement.Order(clause.OrderByColumn{
				Column: clause.Column{Name: name},
				Desc:   desc,
			})
		}
		return statement
	}
}
