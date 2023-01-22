package models

import (
	"time"
)

const TableNameProblemSet = "problem_sets"

// ProblemSet mapped from table <problem_sets>
type ProblemSet struct {
	CreatedAt        time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"updated_at"`
	DomainID         string    `gorm:"column:domain_id;not null"                                                   json:"domain_id"`
	OwnerID          string    `gorm:"column:owner_id"                                                             json:"owner_id"`
	ID               string    `gorm:"column:id;primaryKey"                                                        json:"id"`
	URL              string    `gorm:"column:url;not null"                                                         json:"url"`
	Title            string    `gorm:"column:title;not null"                                                       json:"title"`
	Content          string    `gorm:"column:content;not null"                                                     json:"content"`
	Hidden           bool      `gorm:"column:hidden;not null"                                                      json:"hidden"`
	ScoreboardHidden bool      `gorm:"column:scoreboard_hidden;not null"                                           json:"scoreboard_hidden"`
	NumSubmit        int32     `gorm:"column:num_submit;not null"                                                  json:"num_submit"`
	NumAccept        int32     `gorm:"column:num_accept;not null"                                                  json:"num_accept"`
	DueAt            time.Time `gorm:"column:due_at"                                                               json:"due_at"`
	LockAt           time.Time `gorm:"column:lock_at"                                                              json:"lock_at"`
	UnlockAt         time.Time `gorm:"column:unlock_at"                                                            json:"unlock_at"`
}

// TableName ProblemSet's table name
func (*ProblemSet) TableName() string {
	return TableNameProblemSet
}
