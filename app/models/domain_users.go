package models

import "github.com/google/uuid"

const TableNameDomainUser = "domain_users"

// DomainUser mapped from table <domain_users>
type DomainUser struct {
	Base
	DomainID uuid.UUID `gorm:"column:domain_id;not null;type:uuid;index:idx_domain_users_domain_id_user_id_key,unique" json:"domainId"`
	Domain   Domain    `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	UserID   uuid.UUID `gorm:"column:user_id;not null;type:uuid;index:idx_domain_users_domain_id_user_id_key,unique"   json:"userId"`
	User     User      `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	Role     string    `gorm:"column:role;not null"                                                                    json:"role"`
}

// TableName DomainUser's table name
func (*DomainUser) TableName() string {
	return TableNameDomainUser
}
