package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameDomainInvitation = "domain_invitations"

// DomainInvitation mapped from table <domain_invitations>
type DomainInvitation struct {
	CreatedAt time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP)" json:"updated_at"`
	ExpireAt  time.Time `gorm:"column:expire_at"                                                            json:"expire_at"`
	DomainID  uuid.UUID `gorm:"column:domain_id;not null"                                                   json:"domain_id"`
	URL       string    `gorm:"column:url;not null"                                                         json:"url"`
	Code      string    `gorm:"column:code;not null"                                                        json:"code"`
	Role      string    `gorm:"column:role;not null"                                                        json:"role"`
	ID        uuid.UUID `gorm:"column:id;primaryKey"                                                        json:"id"`
}

// TableName DomainInvitation's table name
func (*DomainInvitation) TableName() string {
	return TableNameDomainInvitation
}

func (u *DomainInvitation) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
