package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameProblemConfig = "problem_configs"

// ProblemConfig mapped from table <problem_configs>
type ProblemConfig struct {
	CreatedAt     time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"updated_at"`
	ProblemID     uuid.UUID `gorm:"column:problem_id"                                                           json:"problem_id"`
	CommitterID   uuid.UUID `gorm:"column:committer_id"                                                         json:"committer_id"`
	ID            uuid.UUID `gorm:"column:id;primaryKey"                                                        json:"id"`
	CommitID      string    `gorm:"column:commit_id;not null;default:''"                                        json:"commit_id"`
	DataVersion   int32     `gorm:"column:data_version;not null;default:2"                                      json:"data_version"`
	CommitMessage string    `gorm:"column:commit_message;not null;default:''"                                   json:"commit_message"`
}

// TableName ProblemConfig's table name
func (*ProblemConfig) TableName() string {
	return TableNameProblemConfig
}

func (u *ProblemConfig) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
