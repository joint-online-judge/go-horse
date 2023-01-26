package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameRecord = "records"

// Record mapped from table <records>
type Record struct {
	CreatedAt       time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"updated_at"`
	ProblemSetID    uuid.UUID `gorm:"column:problem_set_id"                                                       json:"problem_set_id"`
	ProblemID       uuid.UUID `gorm:"column:problem_id"                                                           json:"problem_id"`
	ID              uuid.UUID `gorm:"column:id;primaryKey"                                                        json:"id"`
	Score           int32     `gorm:"column:score;not null"                                                       json:"score"`
	TimeMs          int32     `gorm:"column:time_ms;not null"                                                     json:"time_ms"`
	MemoryKb        int32     `gorm:"column:memory_kb;not null"                                                   json:"memory_kb"`
	CommitID        string    `gorm:"column:commit_id"                                                            json:"commit_id"`
	ProblemConfigID uuid.UUID `gorm:"column:problem_config_id"                                                    json:"problem_config_id"`
	CommitterID     uuid.UUID `gorm:"column:committer_id"                                                         json:"committer_id"`
	JudgerID        uuid.UUID `gorm:"column:judger_id"                                                            json:"judger_id"`
	State           string    `gorm:"column:state;not null;default:processing"                                    json:"state"`
	Language        string    `gorm:"column:language;not null;default:''"                                         json:"language"`
	TaskID          string    `gorm:"column:task_id"                                                              json:"task_id"`
	Cases           string    `gorm:"column:cases;not null;default:'[]'::json"                                    json:"cases"`
	DomainID        uuid.UUID `gorm:"column:domain_id;not null"                                                   json:"domain_id"`
	JudgedAt        time.Time `gorm:"column:judged_at"                                                            json:"judged_at"`
}

// TableName Record's table name
func (*Record) TableName() string {
	return TableNameRecord
}

func (u *Record) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
