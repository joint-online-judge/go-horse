package query

import (
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/pkg/errors"
)

var db *gorm.DB

func NewDB(newDB *gorm.DB) {
	db = newDB
}

func ListObjs[Schema any](statement *gorm.DB, pagination schema.Pagination) ([]Schema, int64, error) {
	var schema []Schema
	var count int64
	if err := statement.Count(&count).Error; err != nil {
		return nil, 0, errors.Wrapf(err, "failed get %T count", schema)
	}
	err := statement.Scopes(Paginate(pagination)).Scan(&schema).Error
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}
	return schema, count, nil
}

func Paginate(pagination schema.Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pagination.Offset == nil || *pagination.Offset < 0 {
			pagination.Offset = schema.Pointer(0)
		}
		if pagination.Limit == nil || *pagination.Limit > 100 {
			pagination.Limit = schema.Pointer(100)
		}
		if pagination.Ordering == nil {
			pagination.Ordering = schema.Pointer("-created_at")
		}
		statement := db.Offset(*pagination.Offset).Limit(*pagination.Limit)
		for _, item := range strings.Split(*pagination.Ordering, ",") {
			name := item
			desc := false
			if strings.HasPrefix(name, "-") {
				name = strings.TrimLeft(item, "-")
				desc = true
			}
			if name != "" {
				statement = statement.Order(clause.OrderByColumn{
					Column: clause.Column{Name: name},
					Desc:   desc,
				})
			}
		}
		return statement
	}
}
