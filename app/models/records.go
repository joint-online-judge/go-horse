package models

import (
	"time"
)

const TableNameRecord = "records"

// Record mapped from table <records>
type Record struct {
	CreatedAt       time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"updated_at"`
	ProblemSetID    string    `gorm:"column:problem_set_id"                                                       json:"problem_set_id"`
	ProblemID       string    `gorm:"column:problem_id"                                                           json:"problem_id"`
	ID              string    `gorm:"column:id;primaryKey"                                                        json:"id"`
	Score           int32     `gorm:"column:score;not null"                                                       json:"score"`
	TimeMs          int32     `gorm:"column:time_ms;not null"                                                     json:"time_ms"`
	MemoryKb        int32     `gorm:"column:memory_kb;not null"                                                   json:"memory_kb"`
	CommitID        string    `gorm:"column:commit_id"                                                            json:"commit_id"`
	ProblemConfigID string    `gorm:"column:problem_config_id"                                                    json:"problem_config_id"`
	CommitterID     string    `gorm:"column:committer_id"                                                         json:"committer_id"`
	JudgerID        string    `gorm:"column:judger_id"                                                            json:"judger_id"`
	State           string    `gorm:"column:state;not null;default:processing"                                    json:"state"`
	Language        string    `gorm:"column:language;not null;default:''::character varying"                      json:"language"`
	TaskID          string    `gorm:"column:task_id"                                                              json:"task_id"`
	Cases           string    `gorm:"column:cases;not null;default:'[]'::json"                                    json:"cases"`
	DomainID        string    `gorm:"column:domain_id;not null"                                                   json:"domain_id"`
	JudgedAt        time.Time `gorm:"column:judged_at"                                                            json:"judged_at"`
}

// TableName Record's table name
func (*Record) TableName() string {
	return TableNameRecord
}
