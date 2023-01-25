package schemas

// AdminListDomainRolesParams defines parameters for AdminListDomainRoles.
type AdminListDomainRolesParams struct {
	// Ordering Comma separated list of ordering the results.
	// You may specify reverse orderings by prefixing the field name with '-'.
	//
	// Available fields: created_at,updated_at
	Ordering *string `form:"ordering,omitempty" json:"ordering,omitempty"`
	Offset   *int    `form:"offset,omitempty"   json:"offset,omitempty"`
	Limit    *int    `form:"limit,omitempty"    json:"limit,omitempty"`
}

// AdminListJudgersParams defines parameters for AdminListJudgers.
type AdminListJudgersParams struct {
	// Ordering Comma separated list of ordering the results.
	// You may specify reverse orderings by prefixing the field name with '-'.
	//
	// Available fields: created_at,updated_at
	Ordering *string `form:"ordering,omitempty" json:"ordering,omitempty"`
	Offset   *int    `form:"offset,omitempty"   json:"offset,omitempty"`
	Limit    *int    `form:"limit,omitempty"    json:"limit,omitempty"`
}

// AdminListUsersParams defines parameters for AdminListUsers.
type AdminListUsersParams struct {
	// Ordering Comma separated list of ordering the results.
	// You may specify reverse orderings by prefixing the field name with '-'.
	//
	// Available fields: created_at,updated_at
	Ordering *string `form:"ordering,omitempty" json:"ordering,omitempty"`
	Offset   *int    `form:"offset,omitempty"   json:"offset,omitempty"`
	Limit    *int    `form:"limit,omitempty"    json:"limit,omitempty"`
}

// AdminListUserDomainsParams defines parameters for AdminListUserDomains.
type AdminListUserDomainsParams struct {
	Role   *[]string `form:"role,omitempty"   json:"role,omitempty"`
	Groups *[]string `form:"groups,omitempty" json:"groups,omitempty"`

	// Ordering Comma separated list of ordering the results.
	// You may specify reverse orderings by prefixing the field name with '-'.
	//
	// Available fields: created_at,updated_at
	Ordering *string `form:"ordering,omitempty" json:"ordering,omitempty"`
	Offset   *int    `form:"offset,omitempty"   json:"offset,omitempty"`
	Limit    *int    `form:"limit,omitempty"    json:"limit,omitempty"`
}
