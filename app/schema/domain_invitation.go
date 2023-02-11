package schema

import (
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// DomainInvitation defines model for DomainInvitation.
type DomainInvitation struct {
	// Code invitation code
	Code     string             `json:"code"`
	DomainId openapi_types.UUID `json:"domainId"`

	// ExpireAt expire time of invitation
	ExpireAt *time.Time         `json:"expireAt,omitempty"`
	Id       openapi_types.UUID `json:"id"`

	// Role domain role after invitation
	Role *string `json:"role,omitempty"`

	// Url (unique) url of the domain
	Url *string `json:"url,omitempty"`
}

// DomainInvitationCreate defines model for DomainInvitationCreate.
type DomainInvitationCreate struct {
	// Code invitation code
	Code     string     `json:"code"`
	ExpireAt *time.Time `json:"expireAt,omitempty"`

	// Role domain role after invitation
	Role *string `json:"role,omitempty"`

	// Url (unique) url of the domain
	Url *string `json:"url,omitempty"`
}

// DomainInvitationEdit defines model for DomainInvitationEdit.
type DomainInvitationEdit struct {
	// Code invitation code
	Code *string `json:"code,omitempty"`

	// ExpireAt expire time of invitation
	ExpireAt *time.Time `json:"expireAt,omitempty"`

	// Role domain role after invitation
	Role *string `json:"role,omitempty"`
}

// DomainInvitationList defines model for DomainInvitationList.
type DomainInvitationList struct {
	Count   *int                `json:"count,omitempty"`
	Results *[]DomainInvitation `json:"results,omitempty"`
}

// ListDomainInvitationsParams defines parameters for ListDomainInvitations.
type ListDomainInvitationsParams struct{ Pagination }

// JoinDomainByInvitationParams defines parameters for JoinDomainByInvitation.
type JoinDomainByInvitationParams struct {
	InvitationCode string `form:"invitationCode" json:"invitationCode"`
}
