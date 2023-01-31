package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameDomainInvitation = "domain_invitations"

// DomainInvitation mapped from table <domain_invitations>
type DomainInvitation struct {
	CreatedAt time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index"                                                                json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index"                                                                json:"updated_at"`
	ExpireAt  time.Time `gorm:"column:expire_at"                                                                                                                                 json:"expire_at"`
	DomainID  uuid.UUID `gorm:"column:domain_id;not null;type:uuid;index:idx_domain_invitations_domain_id_code_key,unique;index:idx_domain_invitations_domain_id_url_key,unique" json:"domain_id"`
	Domain    Domain    `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	URL       string    `gorm:"column:url;not null;index;index:idx_domain_invitations_domain_id_url_key,unique"                                                                  json:"url"`
	Code      string    `gorm:"column:code;not null;index;index:idx_domain_invitations_domain_id_code_key,unique"                                                                json:"code"`
	Role      string    `gorm:"column:role;not null"                                                                                                                             json:"role"`
	ID        uuid.UUID `gorm:"column:id;primaryKey;type:uuid"                                                                                                                   json:"id"`
}

// TableName DomainInvitation's table name
func (*DomainInvitation) TableName() string {
	return TableNameDomainInvitation
}

func (u *DomainInvitation) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
