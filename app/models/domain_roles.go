package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameDomainRole = "domain_roles"

// DomainRole mapped from table <domain_roles>
type DomainRole struct {
	CreatedAt  time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"updated_at"`
	Permission string    `gorm:"column:permission;not null;type:json"                                        json:"permission"`
	DomainID   uuid.UUID `gorm:"column:domain_id;not null;type:uuid"                                         json:"domain_id"`
	Domain     Domain    `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	ID         uuid.UUID `gorm:"column:id;primaryKey;type:uuid"                                              json:"id"`
	Role       string    `gorm:"column:role;not null"                                                        json:"role"`
}

// TableName DomainRole's table name
func (*DomainRole) TableName() string {
	return TableNameDomainRole
}

func (u *DomainRole) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
