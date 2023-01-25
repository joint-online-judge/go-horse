package schemas

import openapi_types "github.com/deepmap/oapi-codegen/pkg/types"

// DomainUserAdd defines model for DomainUserAdd.
type DomainUserAdd struct {
	Role *string `json:"role,omitempty"`

	// User 'me' or id of the user
	User string `json:"user"`
}

// DomainUserPermission defines model for DomainUserPermission.
type DomainUserPermission struct {
	DomainId openapi_types.UUID `json:"domainId"`

	// Permission All permissions in a domain
	Permission DomainPermission   `json:"permission"`
	Role       string             `json:"role"`
	UserId     openapi_types.UUID `json:"userId"`
}

// DomainUserUpdate defines model for DomainUserUpdate.
type DomainUserUpdate struct {
	Role *string `json:"role,omitempty"`
}

// ListDomainUsersParams defines parameters for ListDomainUsers.
type ListDomainUsersParams struct {
	// Ordering Comma separated list of ordering the results.
	// You may specify reverse orderings by prefixing the field name with '-'.
	//
	// Available fields: created_at,updated_at
	Ordering *string `form:"ordering,omitempty" json:"ordering,omitempty"`
	Offset   *int    `form:"offset,omitempty"   json:"offset,omitempty"`
	Limit    *int    `form:"limit,omitempty"    json:"limit,omitempty"`
}
