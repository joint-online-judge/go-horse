package schema

import "github.com/google/uuid"

// DomainUserAdd defines model for DomainUserAdd.
type DomainUserAdd struct {
	Role *string `json:"role,omitempty"`

	// User 'me' or id of the user
	User string `json:"user"`
}

// DomainUserPermission defines model for DomainUserPermission.
type DomainUserPermission struct {
	DomainId uuid.UUID `json:"domainId"`

	// Permission All permissions in a domain
	Permission DomainPermission `json:"permission"`
	Role       string           `json:"role"`
	UserId     uuid.UUID        `json:"userId"`
}

// DomainUserUpdate defines model for DomainUserUpdate.
type DomainUserUpdate struct {
	Role *string `json:"role,omitempty"`
}

// ListDomainUsersParams defines parameters for ListDomainUsers.
type ListDomainUsersParams struct{ Pagination }
