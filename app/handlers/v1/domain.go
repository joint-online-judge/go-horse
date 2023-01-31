package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/models"
	"github.com/joint-online-judge/go-horse/app/querys"
	"github.com/joint-online-judge/go-horse/app/schemas"
	"github.com/sirupsen/logrus"
)

// List Domains
// (GET /domains)
func (s *Api) ListDomains(
	c *fiber.Ctx,
	request schemas.ListDomainsRequestObject,
) (any, error) {
	objs, count, err := querys.ListObjs[models.Domain, schemas.Domain]()
	return schemas.NewListResp(count, objs), err
}

// Create Domain
// (POST /domains)
func (s *Api) CreateDomain(
	c *fiber.Ctx,
	request schemas.CreateDomainRequestObject,
) (any, error) {
	// TODO: verify input values
	domain := request.Body
	user := schemas.JWTUser(c)
	domainModel := models.Domain{
		Owner:    models.User{Base: models.Base{ID: user.Id}},
		URL:      *domain.Url,
		Name:     domain.Name,
		Gravatar: *domain.Gravatar,
		Bulletin: *domain.Bulletin,
		Hidden:   *domain.Hidden,
		Group_:   *domain.Group,
	}
	return querys.CreateObj[schemas.Domain](&domainModel)
}

// Get Domain
// (GET /domains/{domain})
func (s *Api) GetDomain(
	c *fiber.Ctx,
	request schemas.GetDomainRequestObject,
) (any, error) {
	_, domain, err := querys.GetDomain[schemas.Domain](request.Domain)
	return domain, err
}

// Search Domain Groups
// (GET /domains/groups)
func (s *Api) SearchDomainGroups(
	c *fiber.Ctx,
	request schemas.SearchDomainGroupsRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Delete Domain
// (DELETE /domains/{domain})
func (s *Api) DeleteDomain(
	c *fiber.Ctx,
	request schemas.DeleteDomainRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Update Domain
// (PATCH /domains/{domain})
func (s *Api) UpdateDomain(
	c *fiber.Ctx,
	request schemas.UpdateDomainRequestObject,
) (any, error) {
	domainEdit := request.Body
	domainModel, _, err := querys.GetDomain[schemas.Domain](request.Domain)
	if err != nil {
		return nil, err
	}
	if err := models.Update(&domainModel, domainEdit); err != nil {
		return nil, err
	}
	logrus.Infof("update domain to: %+v", domainModel)
	if err = querys.SaveObj(&domainModel); err != nil {
		return nil, err
	}
	return querys.ConvertTo[schemas.Domain](domainModel)
}

// Search Domain Candidates
// (GET /domains/{domain}/candidates)
func (s *Api) SearchDomainCandidates(
	c *fiber.Ctx,
	request schemas.SearchDomainCandidatesRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// List Domain Invitations
// (GET /domains/{domain}/invitations)
func (s *Api) ListDomainInvitations(
	c *fiber.Ctx,
	request schemas.ListDomainInvitationsRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Create Domain Invitation
// (POST /domains/{domain}/invitations)
func (s *Api) CreateDomainInvitation(
	c *fiber.Ctx,
	request schemas.CreateDomainInvitationRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Delete Domain Invitation
// (DELETE /domains/{domain}/invitations/{invitation})
func (s *Api) DeleteDomainInvitation(
	c *fiber.Ctx,
	request schemas.DeleteDomainInvitationRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Get Domain Invitation
// (GET /domains/{domain}/invitations/{invitation})
func (s *Api) GetDomainInvitation(
	c *fiber.Ctx,
	request schemas.GetDomainInvitationRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Update Domain Invitation
// (PATCH /domains/{domain}/invitations/{invitation})
func (s *Api) UpdateDomainInvitation(
	c *fiber.Ctx,
	request schemas.UpdateDomainInvitationRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Join Domain By Invitation
// (POST /domains/{domain}/join)
func (s *Api) JoinDomainByInvitation(
	c *fiber.Ctx,
	request schemas.JoinDomainByInvitationRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// List Domain Roles
// (GET /domains/{domain}/roles)
func (s *Api) ListDomainRoles(
	c *fiber.Ctx,
	request schemas.ListDomainRolesRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Create Domain Role
// (POST /domains/{domain}/roles)
func (s *Api) CreateDomainRole(
	c *fiber.Ctx,
	request schemas.CreateDomainRoleRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Delete Domain Role
// (DELETE /domains/{domain}/roles/{role})
func (s *Api) DeleteDomainRole(
	c *fiber.Ctx,
	request schemas.DeleteDomainRoleRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Get Domain Role
// (GET /domains/{domain}/roles/{role})
func (s *Api) GetDomainRole(
	c *fiber.Ctx,
	request schemas.GetDomainRoleRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Update Domain Role
// (PATCH /domains/{domain}/roles/{role})
func (s *Api) UpdateDomainRole(
	c *fiber.Ctx,
	request schemas.UpdateDomainRoleRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Transfer Domain
// (POST /domains/{domain}/transfer)
func (s *Api) TransferDomain(
	c *fiber.Ctx,
	request schemas.TransferDomainRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// List Domain Users
// (GET /domains/{domain}/users)
func (s *Api) ListDomainUsers(
	c *fiber.Ctx,
	request schemas.ListDomainUsersRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Add Domain User
// (POST /domains/{domain}/users)
func (s *Api) AddDomainUser(
	c *fiber.Ctx,
	request schemas.AddDomainUserRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Remove Domain User
// (DELETE /domains/{domain}/users/{user})
func (s *Api) RemoveDomainUser(
	c *fiber.Ctx,
	request schemas.RemoveDomainUserRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Get Domain User
// (GET /domains/{domain}/users/{user})
func (s *Api) GetDomainUser(
	c *fiber.Ctx,
	request schemas.GetDomainUserRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Update Domain User
// (PATCH /domains/{domain}/users/{user})
func (s *Api) UpdateDomainUser(
	c *fiber.Ctx,
	request schemas.UpdateDomainUserRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Get Domain User Permission
// (GET /domains/{domain}/users/{user}/permission)
func (s *Api) GetDomainUserPermission(
	c *fiber.Ctx,
	request schemas.GetDomainUserPermissionRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}
