package models

import "github.com/google/uuid"

const TableNameDomain = "domains"

// Domain mapped from table <domains>
type Domain struct {
	Base
	OwnerID  uuid.UUID `gorm:"column:owner_id;type:uuid"                       json:"ownerId"`
	Owner    User      `gorm:"constraint:OnDelete:SET NULL,OnUpdate:NO ACTION"`
	URL      string    `gorm:"column:url;not null;index:unique"                json:"url"      validate:"domain_url"`
	Name     string    `gorm:"column:name;not null"                            json:"name"`
	Gravatar string    `gorm:"column:gravatar;not null"                        json:"gravatar"`
	Bulletin string    `gorm:"column:bulletin;not null"                        json:"bulletin"`
	Hidden   bool      `gorm:"column:hidden;not null"                          json:"hidden"`
	Group_   string    `gorm:"column:group;not null;default:''"                json:"group"`
	Tag      string    `gorm:"column:tag"                                      json:"tag"`
}

// TableName Domain's table name
func (*Domain) TableName() string {
	return TableNameDomain
}
