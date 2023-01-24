package models

import (
	"time"
)

const TableNameDomainUser = "domain_users"

// DomainUser mapped from table <domain_users>
type DomainUser struct {
	CreatedAt time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"updated_at"`
	DomainID  string    `gorm:"column:domain_id;not null"                                                   json:"domain_id"`
	UserID    string    `gorm:"column:user_id;not null"                                                     json:"user_id"`
	ID        string    `gorm:"column:id;primaryKey"                                                        json:"id"`
	Role      string    `gorm:"column:role;not null"                                                        json:"role"`
}

// TableName DomainUser's table name
func (*DomainUser) TableName() string {
	return TableNameDomainUser
}
