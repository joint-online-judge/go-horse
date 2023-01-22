package handlers

import (
	"context"

	"github.com/joint-online-judge/go-horse/schemas"
)

// Admin List Domain Roles
// (GET /admin/domain_roles)
func (s *ApiV1) AdminListDomainRoles(
	ctx context.Context,
	request schemas.AdminListDomainRolesRequestObject,
) (any, error) {
	return nil, nil
}

// Admin List Judgers
// (GET /admin/judgers)
func (s *ApiV1) AdminListJudgers(
	ctx context.Context,
	request schemas.AdminListJudgersRequestObject,
) (any, error) {
	return nil, nil
}

// Admin Create Judger
// (POST /admin/judgers)
func (s *ApiV1) AdminCreateJudger(
	ctx context.Context,
	request schemas.AdminCreateJudgerRequestObject,
) (any, error) {
	return nil, nil
}

// Admin List Users
// (GET /admin/users)
func (s *ApiV1) AdminListUsers(
	ctx context.Context,
	request schemas.AdminListUsersRequestObject,
) (any, error) {
	return nil, nil
}

// Admin Get User
// (GET /admin/{uid})
func (s *ApiV1) AdminGetUser(
	ctx context.Context,
	request schemas.AdminGetUserRequestObject,
) (any, error) {
	return nil, nil
}

// Admin List User Domains
// (GET /admin/{uid}/domains)
func (s *ApiV1) AdminListUserDomains(
	ctx context.Context,
	request schemas.AdminListUserDomainsRequestObject,
) (any, error) {
	return nil, nil
}
