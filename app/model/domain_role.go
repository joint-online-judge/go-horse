package model

import "github.com/google/uuid"

const TableNameDomainRole = "domain_roles"

// DomainRole mapped from table <domain_roles>
type DomainRole struct {
	Base
	Permission string    `gorm:"column:permission;not null;type:json"                                                 json:"permission"`
	DomainID   uuid.UUID `gorm:"column:domain_id;not null;type:uuid;index:idx_domain_roles_domain_id_role_key,unique" json:"domainId"`
	Domain     Domain    `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	Role       string    `gorm:"column:role;not null;index:idx_domain_roles_domain_id_role_key,unique"                json:"role"`
}

// TableName DomainRole's table name
func (*DomainRole) TableName() string {
	return TableNameDomainRole
}
