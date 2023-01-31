package models

import (
	"github.com/google/uuid"
)

const TableNameProblem = "problems"

// Problem mapped from table <problems>
type Problem struct {
	Base
	DomainID       uuid.UUID    `gorm:"column:domain_id;not null;type:uuid;index:idx_problems_domain_id_url_key,unique" json:"domainId"`
	Domain         Domain       `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	OwnerID        uuid.UUID    `gorm:"column:owner_id;type:uuid"                                                       json:"ownerId"`
	Owner          User         `gorm:"constraint:OnDelete:SET NULL,OnUpdate:NO ACTION"`
	ProblemGroupID uuid.UUID    `gorm:"column:problem_group_id;type:uuid"                                               json:"problemGroupId"`
	ProblemGroup   ProblemGroup `gorm:"constraint:OnDelete:SET NULL,OnUpdate:NO ACTION"`
	URL            string       `gorm:"column:url;not null;index;index:idx_problems_domain_id_url_key,unique"           json:"url"`
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
