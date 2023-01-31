package models

import (
	"time"

	"github.com/google/uuid"
)

const TableNameDomainInvitation = "domain_invitations"

// DomainInvitation mapped from table <domain_invitations>
type DomainInvitation struct {
	Base
	ExpireAt time.Time `gorm:"column:expire_at"                                                                                                                                 json:"expireAt"`
	DomainID uuid.UUID `gorm:"column:domain_id;not null;type:uuid;index:idx_domain_invitations_domain_id_code_key,unique;index:idx_domain_invitations_domain_id_url_key,unique" json:"domainId"`
	Domain   Domain    `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	URL      string    `gorm:"column:url;not null;index;index:idx_domain_invitations_domain_id_url_key,unique"                                                                  json:"url"`
	Code     string    `gorm:"column:code;not null;index;index:idx_domain_invitations_domain_id_code_key,unique"                                                                json:"code"`
	Role     string    `gorm:"column:role;not null"                                                                                                                             json:"role"`
}

// TableName DomainInvitation's table name
func (*DomainInvitation) TableName() string {
	return TableNameDomainInvitation
}
