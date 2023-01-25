package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameUserLatestRecord = "user_latest_records"

// UserLatestRecord mapped from table <user_latest_records>
type UserLatestRecord struct {
	CreatedAt    time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"updated_at"`
	UserID       uuid.UUID `gorm:"column:user_id;not null"                                                     json:"user_id"`
	ProblemID    uuid.UUID `gorm:"column:problem_id;not null"                                                  json:"problem_id"`
	ProblemSetID uuid.UUID `gorm:"column:problem_set_id"                                                       json:"problem_set_id"`
	RecordID     uuid.UUID `gorm:"column:record_id;not null"                                                   json:"record_id"`
	ID           uuid.UUID `gorm:"column:id;primaryKey"                                                        json:"id"`
}

// TableName UserLatestRecord's table name
func (*UserLatestRecord) TableName() string {
	return TableNameUserLatestRecord
}

func (u *UserLatestRecord) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
