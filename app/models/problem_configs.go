package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameProblemConfig = "problem_configs"

// ProblemConfig mapped from table <problem_configs>
type ProblemConfig struct {
	CreatedAt     time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index" json:"createdAt"`
	UpdatedAt     time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index" json:"updatedAt"`
	ProblemID     uuid.UUID `gorm:"column:problem_id;type:uuid"                                                       json:"problemId"`
	Problem       Problem   `gorm:"constraint:OnDelete:SET NULL,OnUpdate:NO ACTION"`
	CommitterID   uuid.UUID `gorm:"column:committer_id;type:uuid"                                                     json:"committerId"`
	Committer     User      `gorm:"constraint:OnDelete:SET NULL,OnUpdate:NO ACTION"`
	ID            uuid.UUID `gorm:"column:id;primaryKey;type:uuid"                                                    json:"id"`
	CommitID      string    `gorm:"column:commit_id;not null;default:''"                                              json:"commitId"`
	DataVersion   int32     `gorm:"column:data_version;not null;default:2"                                            json:"dataVersion"`
	CommitMessage string    `gorm:"column:commit_message;not null;default:''"                                         json:"commitMessage"`
}

// TableName ProblemConfig's table name
func (*ProblemConfig) TableName() string {
	return TableNameProblemConfig
}

func (u *ProblemConfig) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
