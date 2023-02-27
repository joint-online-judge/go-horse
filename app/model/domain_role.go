package model

import (
	"time"

	"github.com/google/uuid"
)

const TableNameDomainRole = "domain_roles"

// DomainRole mapped from table <domain_roles>
type DomainRole struct {
	Id         uuid.UUID `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"                         json:"id"`
	CreatedAt  time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index" json:"updatedAt"`
	Permission string    `gorm:"column:permission;not null;type:json"                                                 json:"permission"`
	DomainId   uuid.UUID `gorm:"column:domain_id;not null;type:uuid;index:idx_domain_roles_domain_id_role_key,unique" json:"domainId"`
	Domain     Domain    `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	Role       string    `gorm:"column:role;not null;index:idx_domain_roles_domain_id_role_key,unique"                json:"role"`
}

// TableName DomainRole's table name
func (*DomainRole) TableName() string {
	return TableNameDomainRole
}
