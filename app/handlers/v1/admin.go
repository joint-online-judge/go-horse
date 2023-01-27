package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schemas"
)

// Admin List Domain Roles
// (GET /admin/domain_roles)
func (s *Api) AdminListDomainRoles(
	c *fiber.Ctx,
	request schemas.AdminListDomainRolesRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Admin List Judgers
// (GET /admin/judgers)
func (s *Api) AdminListJudgers(
	c *fiber.Ctx,
	request schemas.AdminListJudgersRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Admin Create Judger
// (POST /admin/judgers)
func (s *Api) AdminCreateJudger(
	c *fiber.Ctx,
	request schemas.AdminCreateJudgerRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Admin List Users
// (GET /admin/users)
func (s *Api) AdminListUsers(
	c *fiber.Ctx,
	request schemas.AdminListUsersRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Admin Get User
// (GET /admin/{uid})
func (s *Api) AdminGetUser(
	c *fiber.Ctx,
	request schemas.AdminGetUserRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Admin List User Domains
// (GET /admin/{uid}/domains)
func (s *Api) AdminListUserDomains(
	c *fiber.Ctx,
	request schemas.AdminListUserDomainsRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}
