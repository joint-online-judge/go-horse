package model

import (
	"time"

	"github.com/google/uuid"
)

const TableNameProblemGroup = "problem_groups"

// ProblemGroup mapped from table <problem_groups>
type ProblemGroup struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"                         json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index" json:"updatedAt"`
}

// TableName ProblemGroup's table name
func (*ProblemGroup) TableName() string {
	return TableNameProblemGroup
}
