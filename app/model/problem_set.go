package model

import (
	"time"

	"github.com/google/uuid"
)

const TableNameProblemSet = "problem_sets"

// ProblemSet mapped from table <problem_sets>
type ProblemSet struct {
	Id               uuid.UUID  `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"                           json:"id"`
	CreatedAt        time.Time  `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index"   json:"createdAt"`
	UpdatedAt        time.Time  `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index"   json:"updatedAt"`
	DomainId         uuid.UUID  `gorm:"column:domain_id;not null;type:uuid;index:idx_problem_sets_domain_id_url_key,unique" json:"domainId"`
	Domain           Domain     `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	OwnerId          *uuid.UUID `gorm:"column:owner_id;type:uuid"                                                           json:"ownerId"`
	Owner            *User      `gorm:"constraint:OnDelete:SET NULL,OnUpdate:NO ACTION"`
	Url              string     `gorm:"column:url;not null;index;index:idx_problem_sets_domain_id_url_key,unique"           json:"url"`
	Title            string     `gorm:"column:title;not null"                                                               json:"title"`
	Content          string     `gorm:"column:content;not null"                                                             json:"content"`
	Hidden           bool       `gorm:"column:hidden;not null"                                                              json:"hidden"`
	ScoreboardHidden bool       `gorm:"column:scoreboard_hidden;not null"                                                   json:"scoreboardHidden"`
	NumSubmit        int32      `gorm:"column:num_submit;not null"                                                          json:"numSubmit"`
	NumAccept        int32      `gorm:"column:num_accept;not null"                                                          json:"numAccept"`
	DueAt            time.Time  `gorm:"column:due_at"                                                                       json:"dueAt"`
	LockAt           time.Time  `gorm:"column:lock_at"                                                                      json:"lockAt"`
	UnlockAt         time.Time  `gorm:"column:unlock_at"                                                                    json:"unlockAt"`
}

// TableName ProblemSet's table name
func (*ProblemSet) TableName() string {
	return TableNameProblemSet
}
