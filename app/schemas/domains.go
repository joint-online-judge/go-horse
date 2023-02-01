package schemas

import (
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

type DefaultRole string

const (
	ROOT   DefaultRole = "root"
	ADMIN  DefaultRole = "admin"
	USER   DefaultRole = "user"
	GUEST  DefaultRole = "guest"
	JUDGER DefaultRole = "judger"
)

var DefaultRoles = []DefaultRole{ROOT, ADMIN, USER, GUEST, JUDGER}

// Domain defines model for Domain.
type Domain struct {
	// Bulletin bulletin of the domain
	Bulletin *string `json:"bulletin,omitempty"`

	// Gravatar gravatar url of the domain
	Gravatar *string `json:"gravatar,omitempty"`

	// Group group name of the domain
	Group *string `json:"group,omitempty"`

	// Hidden is the domain hidden
	Hidden *bool              `json:"hidden,omitempty"`
	Id     openapi_types.UUID `json:"id"`

	// Name displayed name of the domain
	Name    string              `json:"name"`
	OwnerId *openapi_types.UUID `json:"ownerId,omitempty"`

	// Url (unique) url of the domain
	Url *string `json:"url,omitempty"`
}

// DomainCreate defines model for DomainCreate.
type DomainCreate struct {
	// Bulletin bulletin of the domain
	Bulletin *string `json:"bulletin,omitempty"`

	// Gravatar gravatar url of the domain
	Gravatar *string `json:"gravatar,omitempty"`

	// Group group name of the domain
	Group *string `json:"group,omitempty"`

	// Hidden is the domain hidden
	Hidden *bool `json:"hidden,omitempty"`

	// Name displayed name of the domain
	Name string `json:"name" validate:"required"`

	// Url (unique) url of the domain
	Url *string `json:"url,omitempty" validate:"domain_url"`
}

func (t *DomainCreate) ApplyDefault() {
	t.Bulletin = Pointer("")
	t.Gravatar = Pointer("")
	t.Group = Pointer("")
	t.Hidden = Pointer(true)
	t.Url = Pointer("")
}

// DomainDetail defines model for DomainDetail.
type DomainDetail struct {
	// Bulletin bulletin of the domain
	Bulletin  *string    `json:"bulletin,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// Gravatar gravatar url of the domain
	Gravatar *string `json:"gravatar,omitempty"`

	// Group group name of the domain
	Group *string `json:"group,omitempty"`

	// Hidden is the domain hidden
	Hidden *bool              `json:"hidden,omitempty"`
	Id     openapi_types.UUID `json:"id"`

	// Name displayed name of the domain
	Name      string              `json:"name"`
	OwnerId   *openapi_types.UUID `json:"ownerId,omitempty"`
	Tag       *string             `json:"tag,omitempty"`
	UpdatedAt *time.Time          `json:"updatedAt,omitempty"`

	// Url (unique) url of the domain
	Url *string `json:"url,omitempty"`
}

// DomainEdit defines model for DomainEdit.
type DomainEdit struct {
	Bulletin *string `json:"bulletin,omitempty"`
	Gravatar *string `json:"gravatar,omitempty"`
	Group    *string `json:"group,omitempty"`
	Hidden   *bool   `json:"hidden,omitempty"`
	Name     *string `json:"name,omitempty"`
	Url      *string `json:"url,omitempty"`
}

// DomainList defines model for DomainList.
type DomainList struct {
	Count   *int      `json:"count,omitempty"`
	Results *[]Domain `json:"results,omitempty"`
}

// DomainTag defines model for DomainTag.
type DomainTag = string

// DomainTagList defines model for DomainTagList.
type DomainTagList struct {
	Count   *int         `json:"count,omitempty"`
	Results *[]DomainTag `json:"results,omitempty"`
}

// DomainTransfer defines model for DomainTransfer.
type DomainTransfer struct {
	// TargetUser 'me' or id of the user
	TargetUser string `json:"targetUser"`
}

// ListDomainsParams defines parameters for ListDomains.
type ListDomainsParams struct {
	Roles  *[]string `form:"roles,omitempty"  json:"roles,omitempty"`
	Groups *[]string `form:"groups,omitempty" json:"groups,omitempty"`
	Pagination
}

// SearchDomainGroupsParams defines parameters for SearchDomainGroups.
type SearchDomainGroupsParams struct {
	// Query search query
	Query string `form:"query" json:"query"`
}

// SearchDomainCandidatesParams defines parameters for SearchDomainCandidates.
type SearchDomainCandidatesParams struct {
	// Query search query
	Query string `form:"query" json:"query"`

	// Available fields: username
	Ordering *string `form:"ordering,omitempty" json:"ordering,omitempty"`
}
