// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDomainRole = "domain_roles"

// DomainRole mapped from table <domain_roles>
type DomainRole struct {
	CreatedAt  time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"updated_at"`
	Permission string    `gorm:"column:permission;not null" json:"permission"`
	DomainID   string    `gorm:"column:domain_id;not null" json:"domain_id"`
	ID         string    `gorm:"column:id;primaryKey" json:"id"`
	Role       string    `gorm:"column:role;not null" json:"role"`
}

// TableName DomainRole's table name
func (*DomainRole) TableName() string {
	return TableNameDomainRole
}