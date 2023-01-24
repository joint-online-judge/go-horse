package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/schemas"
)

// Admin List Domain Roles
// (GET /admin/domain_roles)
func (s *ApiV1) AdminListDomainRoles(
	c *fiber.Ctx,
	request schemas.AdminListDomainRolesRequestObject,
) (any, error) {
	return nil, nil
}

// Admin List Judgers
// (GET /admin/judgers)
func (s *ApiV1) AdminListJudgers(
	c *fiber.Ctx,
	request schemas.AdminListJudgersRequestObject,
) (any, error) {
	return nil, nil
}

// Admin Create Judger
// (POST /admin/judgers)
func (s *ApiV1) AdminCreateJudger(
	c *fiber.Ctx,
	request schemas.AdminCreateJudgerRequestObject,
) (any, error) {
	return nil, nil
}

// Admin List Users
// (GET /admin/users)
func (s *ApiV1) AdminListUsers(
	c *fiber.Ctx,
	request schemas.AdminListUsersRequestObject,
) (any, error) {
	return nil, nil
}

// Admin Get User
// (GET /admin/{uid})
func (s *ApiV1) AdminGetUser(
	c *fiber.Ctx,
	request schemas.AdminGetUserRequestObject,
) (any, error) {
	return nil, nil
}

// Admin List User Domains
// (GET /admin/{uid}/domains)
func (s *ApiV1) AdminListUserDomains(
	c *fiber.Ctx,
	request schemas.AdminListUserDomainsRequestObject,
) (any, error) {
	return nil, nil
}
