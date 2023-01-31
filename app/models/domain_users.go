package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameDomainUser = "domain_users"

// DomainUser mapped from table <domain_users>
type DomainUser struct {
	CreatedAt time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index"       json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index"       json:"updated_at"`
	DomainID  uuid.UUID `gorm:"column:domain_id;not null;type:uuid;index:idx_domain_users_domain_id_user_id_key,unique" json:"domain_id"`
	Domain    Domain    `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	UserID    uuid.UUID `gorm:"column:user_id;not null;type:uuid;index:idx_domain_users_domain_id_user_id_key,unique"   json:"user_id"`
	User      User      `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	ID        uuid.UUID `gorm:"column:id;primaryKey;type:uuid"                                                          json:"id"`
	Role      string    `gorm:"column:role;not null"                                                                    json:"role"`
}

// TableName DomainUser's table name
func (*DomainUser) TableName() string {
	return TableNameDomainUser
}

func (u *DomainUser) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
