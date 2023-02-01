package schemas

// AdminListDomainRolesParams defines parameters for AdminListDomainRoles.
type AdminListDomainRolesParams struct{ Pagination }

// AdminListJudgersParams defines parameters for AdminListJudgers.
type AdminListJudgersParams struct{ Pagination }

// AdminListUsersParams defines parameters for AdminListUsers.
type AdminListUsersParams struct{ Pagination }

// AdminListUserDomainsParams defines parameters for AdminListUserDomains.
type AdminListUserDomainsParams struct {
	Role   *[]string `form:"role,omitempty"   json:"role,omitempty"`
	Groups *[]string `form:"groups,omitempty" json:"groups,omitempty"`
	Pagination
}
