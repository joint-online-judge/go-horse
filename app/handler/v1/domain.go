package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/query"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/sirupsen/logrus"
)

// List Domains
// (GET /domains)
func (s *Api) ListDomains(
	c *fiber.Ctx,
	request schema.ListDomainsRequestObject,
) (any, error) {
	// TODO: filter by domain users
	objs, count, err := query.ListObjsByType[
		model.Domain, schema.Domain,
	](request.Params.Pagination)
	return schema.NewListResp(count, objs), err
}

// Create Domain
// (POST /domains)
func (s *Api) CreateDomain(
	c *fiber.Ctx,
	request schema.CreateDomainRequestObject,
) (any, error) {
	// TODO: verify input values
	domain := request.Body
	user := schema.JWTUser(c)
	return query.CreateDomain(domain, user)
}

// Get Domain
// (GET /domains/{domain})
func (s *Api) GetDomain(
	c *fiber.Ctx,
	request schema.GetDomainRequestObject,
) (any, error) {
	domainModel, err := query.GetDomain(request.Domain)
	if err != nil {
		return nil, err
	}
	return query.ConvertTo[schema.Domain](domainModel)
}

// Search Domain Groups
// (GET /domains/groups)
func (s *Api) SearchDomainGroups(
	c *fiber.Ctx,
	request schema.SearchDomainGroupsRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Delete Domain
// (DELETE /domains/{domain})
func (s *Api) DeleteDomain(
	c *fiber.Ctx,
	request schema.DeleteDomainRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Update Domain
// (PATCH /domains/{domain})
func (s *Api) UpdateDomain(
	c *fiber.Ctx,
	request schema.UpdateDomainRequestObject,
) (any, error) {
	domainEdit := request.Body
	domainModel, err := query.GetDomain(request.Domain)
	if err != nil {
		return nil, err
	}
	if err := model.Update(&domainModel, domainEdit); err != nil {
		return nil, err
	}
	logrus.Infof("update domain to: %+v", domainModel)
	if err = query.SaveObj(&domainModel); err != nil {
		return nil, err
	}
	return query.ConvertTo[schema.Domain](domainModel)
}

// Search Domain Candidates
// (GET /domains/{domain}/candidates)
func (s *Api) SearchDomainCandidates(
	c *fiber.Ctx,
	request schema.SearchDomainCandidatesRequestObject,
) (any, error) {
	domainModel, err := query.GetDomain(request.Domain)
	if err != nil {
		return nil, err
	}
	pagination := schema.Pagination{
		Ordering: request.Params.Ordering,
		Offset:   nil,
		Limit:    nil,
	}
	return query.SearchDomainCandidates(
		domainModel.ID, request.Params.Query, pagination,
	)
}

// List Domain Invitations
// (GET /domains/{domain}/invitations)
func (s *Api) ListDomainInvitations(
	c *fiber.Ctx,
	request schema.ListDomainInvitationsRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Create Domain Invitation
// (POST /domains/{domain}/invitations)
func (s *Api) CreateDomainInvitation(
	c *fiber.Ctx,
	request schema.CreateDomainInvitationRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Delete Domain Invitation
// (DELETE /domains/{domain}/invitations/{invitation})
func (s *Api) DeleteDomainInvitation(
	c *fiber.Ctx,
	request schema.DeleteDomainInvitationRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Get Domain Invitation
// (GET /domains/{domain}/invitations/{invitation})
func (s *Api) GetDomainInvitation(
	c *fiber.Ctx,
	request schema.GetDomainInvitationRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Update Domain Invitation
// (PATCH /domains/{domain}/invitations/{invitation})
func (s *Api) UpdateDomainInvitation(
	c *fiber.Ctx,
	request schema.UpdateDomainInvitationRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Join Domain By Invitation
// (POST /domains/{domain}/join)
func (s *Api) JoinDomainByInvitation(
	c *fiber.Ctx,
	request schema.JoinDomainByInvitationRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// List Domain Roles
// (GET /domains/{domain}/roles)
func (s *Api) ListDomainRoles(
	c *fiber.Ctx,
	request schema.ListDomainRolesRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Create Domain Role
// (POST /domains/{domain}/roles)
func (s *Api) CreateDomainRole(
	c *fiber.Ctx,
	request schema.CreateDomainRoleRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Delete Domain Role
// (DELETE /domains/{domain}/roles/{role})
func (s *Api) DeleteDomainRole(
	c *fiber.Ctx,
	request schema.DeleteDomainRoleRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Get Domain Role
// (GET /domains/{domain}/roles/{role})
func (s *Api) GetDomainRole(
	c *fiber.Ctx,
	request schema.GetDomainRoleRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Update Domain Role
// (PATCH /domains/{domain}/roles/{role})
func (s *Api) UpdateDomainRole(
	c *fiber.Ctx,
	request schema.UpdateDomainRoleRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Transfer Domain
// (POST /domains/{domain}/transfer)
func (s *Api) TransferDomain(
	c *fiber.Ctx,
	request schema.TransferDomainRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// List Domain Users
// (GET /domains/{domain}/users)
func (s *Api) ListDomainUsers(
	c *fiber.Ctx,
	request schema.ListDomainUsersRequestObject,
) (any, error) {
	domainId, err := query.GetDomainId(request.Domain)
	if err != nil {
		return nil, err
	}
	objs, count, err := query.ListDomainUsers(
		domainId, request.Params.Pagination,
	)
	return schema.NewListResp(count, objs), err
}

// Add Domain User
// (POST /domains/{domain}/users)
func (s *Api) AddDomainUser(
	c *fiber.Ctx,
	request schema.AddDomainUserRequestObject,
) (any, error) {
	domainId, err := query.GetDomainId(request.Domain)
	if err != nil {
		return nil, err
	}
	user, err := query.GetUser(c, request.Body.User)
	if err != nil {
		return nil, err
	}
	return query.AddDomainUser(domainId, user, *request.Body.Role)
}

// Remove Domain User
// (DELETE /domains/{domain}/users/{user})
func (s *Api) RemoveDomainUser(
	c *fiber.Ctx,
	request schema.RemoveDomainUserRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Get Domain User
// (GET /domains/{domain}/users/{user})
func (s *Api) GetDomainUser(
	c *fiber.Ctx,
	request schema.GetDomainUserRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Update Domain User
// (PATCH /domains/{domain}/users/{user})
func (s *Api) UpdateDomainUser(
	c *fiber.Ctx,
	request schema.UpdateDomainUserRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Get Domain User Permission
// (GET /domains/{domain}/users/{user}/permission)
func (s *Api) GetDomainUserPermission(
	c *fiber.Ctx,
	request schema.GetDomainUserPermissionRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}
