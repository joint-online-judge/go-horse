package model

import (
	"time"

	"github.com/google/uuid"
)

const TableNameDomainInvitation = "domain_invitations"

// DomainInvitation mapped from table <domain_invitations>
type DomainInvitation struct {
	Id        uuid.UUID `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"                         json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index" json:"updatedAt"`
	ExpireAt  time.Time `gorm:"column:expire_at"                                                                                                                                 json:"expireAt"`
	DomainId  uuid.UUID `gorm:"column:domain_id;not null;type:uuid;index:idx_domain_invitations_domain_id_code_key,unique;index:idx_domain_invitations_domain_id_url_key,unique" json:"domainId"`
	Domain    Domain    `gorm:"constraint:OnDelete:CASCADE,OnUpdate:NO ACTION"`
	Url       string    `gorm:"column:url;not null;index;index:idx_domain_invitations_domain_id_url_key,unique"                                                                  json:"url"`
	Code      string    `gorm:"column:code;not null;index;index:idx_domain_invitations_domain_id_code_key,unique"                                                                json:"code"`
	Role      string    `gorm:"column:role;not null"                                                                                                                             json:"role"`
}

// TableName DomainInvitation's table name
func (*DomainInvitation) TableName() string {
	return TableNameDomainInvitation
}
