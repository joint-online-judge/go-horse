package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameDomain = "domains"

// Domain mapped from table <domains>
type Domain struct {
	CreatedAt time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"updated_at"`
	OwnerID   uuid.UUID `gorm:"column:owner_id"                                                             json:"owner_id"`
	URL       string    `gorm:"column:url;not null"                                                         json:"url"        validate:"domain_url"`
	Name      string    `gorm:"column:name;not null"                                                        json:"name"`
	Gravatar  string    `gorm:"column:gravatar;not null"                                                    json:"gravatar"`
	Bulletin  string    `gorm:"column:bulletin;not null"                                                    json:"bulletin"`
	Hidden    bool      `gorm:"column:hidden;not null"                                                      json:"hidden"`
	ID        uuid.UUID `gorm:"column:id;primaryKey"                                                        json:"id"`
	Group_    string    `gorm:"column:group;not null;default:''::character varying"                         json:"group"`
	Tag       string    `gorm:"column:tag"                                                                  json:"tag"`
}

// TableName Domain's table name
func (*Domain) TableName() string {
	return TableNameDomain
}

func (u *Domain) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
