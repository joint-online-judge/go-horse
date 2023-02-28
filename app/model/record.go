package model

import (
	"time"

	"github.com/google/uuid"
)

const TableNameRecord = "records"

// Record mapped from table <records>
type Record struct {
	Id              uuid.UUID      `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"                         json:"id"`
	CreatedAt       time.Time      `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index" json:"createdAt"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index" json:"updatedAt"`
	ProblemSetId    *uuid.UUID     `gorm:"column:problem_set_id;type:uuid"                                                   json:"problemSetId"`
	ProblemSet      *ProblemSet    `gorm:"constraint:OnDelete:SET NULL,OnUpdate:NO ACTION"`
	ProblemId       *uuid.UUID     `gorm:"column:problem_id;type:uuid"                                                       json:"problemId"`
	Problem         *Problem       `gorm:"constraint:OnDelete:SET NULL,OnUpdate:NO ACTION"`
	Score           int32          `gorm:"column:score;not null"                                                             json:"score"`
	TimeMs          int32          `gorm:"column:time_ms;not null"                                                           json:"timeMs"`
	MemoryKb        int32          `gorm:"column:memory_kb;not null"                                                         json:"memoryKb"`
	CommitId        string         `gorm:"column:commit_id"                                                                  json:"commitId"`
	ProblemConfigId *uuid.UUID     `gorm:"column:problem_config_id;type:uuid"                                                json:"problemConfigId"`
	ProblemConfig   *ProblemConfig `gorm:"constraint:OnDelete:SET NULL,OnUpdate:NO ACTION"`
	CommitterId     *uuid.UUID     `gorm:"column:committer_id;type:uuid"                                                     json:"committerId"`
	Committer       *User          `gorm:"constraint:OnDelete:SET NULL,OnUpdate:NO ACTION"`
	JudgerId        *uuid.UUID     `gorm:"column:judger_id;type:uuid"                                                        json:"judgerId"`
	Judger          *User          `gorm:"constraint:OnDelete:SET NULL,OnUpdate:NO ACTION"`
	State           string         `gorm:"column:state;not null;default:processing"                                          json:"state"`
	Language        string         `gorm:"column:language;not null;default:''"                                               json:"language"`
	TaskId          string         `gorm:"column:task_id;type:uuid"                                                          json:"taskId"`
	Cases           string         `gorm:"column:cases;not null;type:json;default:'[]'"                                      json:"cases"`
	DomainId        uuid.UUID      `gorm:"column:domain_id;not null;type:uuid"                                               json:"domainId"`
	Domain          Domain         `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	JudgedAt        time.Time      `gorm:"column:judged_at"                                                                  json:"judgedAt"`
}

// TableName Record's table name
func (*Record) TableName() string {
	return TableNameRecord
}
