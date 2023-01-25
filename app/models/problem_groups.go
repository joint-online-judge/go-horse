package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameProblemGroup = "problem_groups"

// ProblemGroup mapped from table <problem_groups>
type ProblemGroup struct {
	CreatedAt time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"updated_at"`
	ID        uuid.UUID `gorm:"column:id;primaryKey"                                                        json:"id"`
}

// TableName ProblemGroup's table name
func (*ProblemGroup) TableName() string {
	return TableNameProblemGroup
}

func (u *ProblemGroup) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
