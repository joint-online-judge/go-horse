package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameProblem = "problems"

// Problem mapped from table <problems>
type Problem struct {
	CreatedAt      time.Time    `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"created_at"`
	UpdatedAt      time.Time    `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"updated_at"`
	DomainID       uuid.UUID    `gorm:"column:domain_id;not null;type:uuid"                                         json:"domain_id"`
	Domain         Domain       `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	OwnerID        uuid.UUID    `gorm:"column:owner_id;type:uuid"                                                   json:"owner_id"`
	Owner          User         `gorm:"constraint:OnDelete:SET NULL,OnUpdate:NO ACTION"`
	ProblemGroupID uuid.UUID    `gorm:"column:problem_group_id;type:uuid"                                           json:"problem_group_id"`
	ProblemGroup   ProblemGroup `gorm:"constraint:OnDelete:SET NULL,OnUpdate:NO ACTION"`
	ID             uuid.UUID    `gorm:"column:id;primaryKey;type:uuid"                                              json:"id"`
	URL            string       `gorm:"column:url;not null"                                                         json:"url"`
	Title          string       `gorm:"column:title;not null"                                                       json:"title"`
	Content        string       `gorm:"column:content;not null"                                                     json:"content"`
	Hidden         bool         `gorm:"column:hidden;not null"                                                      json:"hidden"`
	NumSubmit      int32        `gorm:"column:num_submit;not null"                                                  json:"num_submit"`
	NumAccept      int32        `gorm:"column:num_accept;not null"                                                  json:"num_accept"`
	Languages      string       `gorm:"column:languages;not null;type:json;default:'[]'"                            json:"languages"`
}

// TableName Problem's table name
func (*Problem) TableName() string {
	return TableNameProblem
}

func (u *Problem) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
