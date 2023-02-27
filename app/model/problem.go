package model

import (
	"time"

	"github.com/google/uuid"
)

const TableNameProblem = "problems"

// Problem mapped from table <problems>
type Problem struct {
	Id             uuid.UUID    `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"                         json:"id"`
	CreatedAt      time.Time    `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index" json:"createdAt"`
	UpdatedAt      time.Time    `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index" json:"updatedAt"`
	DomainId       uuid.UUID    `gorm:"column:domain_id;not null;type:uuid;index:idx_problems_domain_id_url_key,unique" json:"domainId"`
	Domain         Domain       `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	OwnerId        uuid.UUID    `gorm:"column:owner_id;type:uuid"                                                       json:"ownerId"`
	Owner          User         `gorm:"constraint:OnDelete:SET NULL,OnUpdate:NO ACTION"`
	ProblemGroupId uuid.UUID    `gorm:"column:problem_group_id;type:uuid"                                               json:"problemGroupId"`
	ProblemGroup   ProblemGroup `gorm:"constraint:OnDelete:SET NULL,OnUpdate:NO ACTION"`
	Url            string       `gorm:"column:url;not null;index;index:idx_problems_domain_id_url_key,unique"           json:"url"`
	Title          string       `gorm:"column:title;not null"                                                           json:"title"`
	Content        string       `gorm:"column:content;not null"                                                         json:"content"`
	Hidden         bool         `gorm:"column:hidden;not null"                                                          json:"hidden"`
	NumSubmit      int32        `gorm:"column:num_submit;not null"                                                      json:"numSubmit"`
	NumAccept      int32        `gorm:"column:num_accept;not null"                                                      json:"numAccept"`
	Languages      string       `gorm:"column:languages;not null;type:json;default:'[]'"                                json:"languages"`
}

// TableName Problem's table name
func (*Problem) TableName() string {
	return TableNameProblem
}
