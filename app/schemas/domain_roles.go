package schemas

import (
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// DomainPermission All permissions in a domain
type DomainPermission struct {
	General    *GeneralPermission    `json:"general,omitempty"`
	Problem    *ProblemPermission    `json:"problem,omitempty"`
	ProblemSet *ProblemSetPermission `json:"problemSet,omitempty"`
	Record     *RecordPermission     `json:"record,omitempty"`
}

// DomainRole defines model for DomainRole.
type DomainRole struct {
	DomainId openapi_types.UUID `json:"domainId"`
	Id       openapi_types.UUID `json:"id"`

	// Permission All permissions in a domain
	Permission DomainPermission `json:"permission"`
	Role       string           `json:"role"`
}

// DomainRoleCreate defines model for DomainRoleCreate.
type DomainRoleCreate struct {
	// Permission All permissions in a domain
	Permission DomainPermission `json:"permission"`
	Role       string           `json:"role"`
}

// DomainRoleDetail defines model for DomainRoleDetail.
type DomainRoleDetail struct {
	CreatedAt *time.Time         `json:"createdAt,omitempty"`
	DomainId  openapi_types.UUID `json:"domainId"`
	Id        openapi_types.UUID `json:"id"`

	// Permission All permissions in a domain
	Permission DomainPermission `json:"permission"`
	Role       string           `json:"role"`
	UpdatedAt  *time.Time       `json:"updatedAt,omitempty"`
}

// DomainRoleEdit defines model for DomainRoleEdit.
type DomainRoleEdit struct {
	// Permission The permission which needs to be updated
	Permission *map[string]any `json:"permission,omitempty"`
}

// DomainRoleList defines model for DomainRoleList.
type DomainRoleList struct {
	Count   *int          `json:"count,omitempty"`
	Results *[]DomainRole `json:"results,omitempty"`
}

// GeneralPermission defines model for GeneralPermission.
type GeneralPermission struct {
	Edit           *bool `json:"edit,omitempty"`
	EditPermission *bool `json:"editPermission,omitempty"`
	UnlimitedQuota *bool `json:"unlimitedQuota,omitempty"`
	View           *bool `json:"view,omitempty"`
	ViewModBadge   *bool `json:"viewModBadge,omitempty"`
}

// ListDomainRolesParams defines parameters for ListDomainRoles.
type ListDomainRolesParams struct {
	// Available fields: created_at,updated_at
	Ordering *string `form:"ordering,omitempty" json:"ordering,omitempty"`
}
