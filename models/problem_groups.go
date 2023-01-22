package models

import (
	"time"
)

const TableNameProblemGroup = "problem_groups"

// ProblemGroup mapped from table <problem_groups>
type ProblemGroup struct {
	CreatedAt time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"updated_at"`
	ID        string    `gorm:"column:id;primaryKey"                                                        json:"id"`
}

// TableName ProblemGroup's table name
func (*ProblemGroup) TableName() string {
	return TableNameProblemGroup
}
