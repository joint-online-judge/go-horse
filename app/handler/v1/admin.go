package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schema"
)

// Admin List Domain Roles
// (GET /admin/domain_roles)
func (s *Api) AdminListDomainRoles(
	c *fiber.Ctx,
	request schema.AdminListDomainRolesRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Admin List Judgers
// (GET /admin/judgers)
func (s *Api) AdminListJudgers(
	c *fiber.Ctx,
	request schema.AdminListJudgersRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Admin Create Judger
// (POST /admin/judgers)
func (s *Api) AdminCreateJudger(
	c *fiber.Ctx,
	request schema.AdminCreateJudgerRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Admin List Users
// (GET /admin/users)
func (s *Api) AdminListUsers(
	c *fiber.Ctx,
	request schema.AdminListUsersRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Admin Get User
// (GET /admin/{uid})
func (s *Api) AdminGetUser(
	c *fiber.Ctx,
	request schema.AdminGetUserRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Admin List User Domains
// (GET /admin/{uid}/domains)
func (s *Api) AdminListUserDomains(
	c *fiber.Ctx,
	request schema.AdminListUserDomainsRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}
