package model

import (
	"time"

	"github.com/google/uuid"
)

const TableNameDomainUser = "domain_users"

// DomainUser mapped from table <domain_users>
type DomainUser struct {
	Id        uuid.UUID `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"                         json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index" json:"updatedAt"`
	DomainId  uuid.UUID `gorm:"column:domain_id;not null;type:uuid;index:idx_domain_users_domain_id_user_id_key,unique" json:"domainId"`
	Domain    Domain    `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	UserId    uuid.UUID `gorm:"column:user_id;not null;type:uuid;index:idx_domain_users_domain_id_user_id_key,unique"   json:"userId"`
	User      User      `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	Role      string    `gorm:"column:role;not null"                                                                    json:"role"`
}

// TableName DomainUser's table name
func (*DomainUser) TableName() string {
	return TableNameDomainUser
}
