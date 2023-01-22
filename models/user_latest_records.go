package models

import (
	"time"
)

const TableNameUserLatestRecord = "user_latest_records"

// UserLatestRecord mapped from table <user_latest_records>
type UserLatestRecord struct {
	CreatedAt    time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"updated_at"`
	UserID       string    `gorm:"column:user_id;not null"                                                     json:"user_id"`
	ProblemID    string    `gorm:"column:problem_id;not null"                                                  json:"problem_id"`
	ProblemSetID string    `gorm:"column:problem_set_id"                                                       json:"problem_set_id"`
	RecordID     string    `gorm:"column:record_id;not null"                                                   json:"record_id"`
	ID           string    `gorm:"column:id;primaryKey"                                                        json:"id"`
}

// TableName UserLatestRecord's table name
func (*UserLatestRecord) TableName() string {
	return TableNameUserLatestRecord
}
