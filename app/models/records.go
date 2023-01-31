package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameRecord = "records"

// Record mapped from table <records>
type Record struct {
	CreatedAt       time.Time     `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"created_at"`
	UpdatedAt       time.Time     `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"updated_at"`
	ProblemSetID    uuid.UUID     `gorm:"column:problem_set_id;type:uuid"                                             json:"problem_set_id"`
	ProblemSet      ProblemSet    `gorm:"constraint:OnDelete:SET NULL,OnUpdate:NO ACTION"`
	ProblemID       uuid.UUID     `gorm:"column:problem_id;type:uuid"                                                 json:"problem_id"`
	Problem         Problem       `gorm:"constraint:OnDelete:SET NULL,OnUpdate:NO ACTION"`
	ID              uuid.UUID     `gorm:"column:id;primaryKey;type:uuid"                                              json:"id"`
	Score           int32         `gorm:"column:score;not null"                                                       json:"score"`
	TimeMs          int32         `gorm:"column:time_ms;not null"                                                     json:"time_ms"`
	MemoryKb        int32         `gorm:"column:memory_kb;not null"                                                   json:"memory_kb"`
	CommitID        string        `gorm:"column:commit_id"                                                            json:"commit_id"`
	ProblemConfigID uuid.UUID     `gorm:"column:problem_config_id;type:uuid"                                          json:"problem_config_id"`
	ProblemConfig   ProblemConfig `gorm:"constraint:OnDelete:SET NULL,OnUpdate:NO ACTION"`
	CommitterID     uuid.UUID     `gorm:"column:committer_id;type:uuid"                                               json:"committer_id"`
	Committer       User          `gorm:"constraint:OnDelete:SET NULL,OnUpdate:NO ACTION"`
	JudgerID        uuid.UUID     `gorm:"column:judger_id;type:uuid"                                                  json:"judger_id"`
	Judger          User          `gorm:"constraint:OnDelete:SET NULL,OnUpdate:NO ACTION"`
	State           string        `gorm:"column:state;not null;default:processing"                                    json:"state"`
	Language        string        `gorm:"column:language;not null;default:''"                                         json:"language"`
	TaskID          string        `gorm:"column:task_id;type:uuid"                                                    json:"task_id"`
	Cases           string        `gorm:"column:cases;not null;type:json;default:'[]'"                                json:"cases"`
	DomainID        uuid.UUID     `gorm:"column:domain_id;not null;type:uuid"                                         json:"domain_id"`
	Domain          Domain        `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	JudgedAt        time.Time     `gorm:"column:judged_at"                                                            json:"judged_at"`
}

// TableName Record's table name
func (*Record) TableName() string {
	return TableNameRecord
}

func (u *Record) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
