package models

import "github.com/google/uuid"

const TableNameUserLatestRecord = "user_latest_records"

// UserLatestRecord mapped from table <user_latest_records>
type UserLatestRecord struct {
	Base
	UserID       uuid.UUID  `gorm:"column:user_id;not null;type:uuid;index:idx_user_latest_records_user_id_problem_id_problem_set_id_recor_key,unique"    json:"userId"`
	User         User       `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	ProblemID    uuid.UUID  `gorm:"column:problem_id;not null;type:uuid;index:idx_user_latest_records_user_id_problem_id_problem_set_id_recor_key,unique" json:"problemId"`
	Problem      Problem    `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	ProblemSetID uuid.UUID  `gorm:"column:problem_set_id;type:uuid;index:idx_user_latest_records_user_id_problem_id_problem_set_id_recor_key,unique"      json:"problemSetId"`
	ProblemSet   ProblemSet `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	RecordID     uuid.UUID  `gorm:"column:record_id;not null;type:uuid;index:idx_user_latest_records_user_id_problem_id_problem_set_id_recor_key,unique"  json:"recordId"`
	Record       Record     `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
}

// TableName UserLatestRecord's table name
func (*UserLatestRecord) TableName() string {
	return TableNameUserLatestRecord
}
