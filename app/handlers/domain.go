package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/models"
	"github.com/joint-online-judge/go-horse/app/querys"
	"github.com/joint-online-judge/go-horse/app/schemas"
	log "github.com/sirupsen/logrus"
)

// List Domains
// (GET /domains)
func (s *ApiV1) ListDomains(
	c *fiber.Ctx,
	request schemas.ListDomainsRequestObject,
) (any, error) {
	objs, count, err := querys.ListObjs[models.Domain, schemas.Domain]()
	return schemas.NewListResp(count, objs), err
}

// Create Domain
// (POST /domains)
func (s *ApiV1) CreateDomain(
	c *fiber.Ctx,
	request schemas.CreateDomainRequestObject,
) (any, error) {
	// TODO: verify input values
	domain := request.Body
	user := schemas.JWTUser(c)
	domainModel := models.Domain{
		OwnerID:  user.Id,
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
func (s *ApiV1) GetDomain(
	c *fiber.Ctx,
	request schemas.GetDomainRequestObject,
) (any, error) {
	log.Info("get domain")
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Search Domain Groups
// (GET /domains/groups)
func (s *ApiV1) SearchDomainGroups(
	c *fiber.Ctx,
	request schemas.SearchDomainGroupsRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Delete Domain
// (DELETE /domains/{domain})
func (s *ApiV1) DeleteDomain(
	c *fiber.Ctx,
	request schemas.DeleteDomainRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Update Domain
// (PATCH /domains/{domain})
func (s *ApiV1) UpdateDomain(
	c *fiber.Ctx,
	request schemas.UpdateDomainRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Search Domain Candidates
// (GET /domains/{domain}/candidates)
func (s *ApiV1) SearchDomainCandidates(
	c *fiber.Ctx,
	request schemas.SearchDomainCandidatesRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// List Domain Invitations
// (GET /domains/{domain}/invitations)
func (s *ApiV1) ListDomainInvitations(
	c *fiber.Ctx,
	request schemas.ListDomainInvitationsRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Create Domain Invitation
// (POST /domains/{domain}/invitations)
func (s *ApiV1) CreateDomainInvitation(
	c *fiber.Ctx,
	request schemas.CreateDomainInvitationRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Delete Domain Invitation
// (DELETE /domains/{domain}/invitations/{invitation})
func (s *ApiV1) DeleteDomainInvitation(
	c *fiber.Ctx,
	request schemas.DeleteDomainInvitationRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Get Domain Invitation
// (GET /domains/{domain}/invitations/{invitation})
func (s *ApiV1) GetDomainInvitation(
	c *fiber.Ctx,
	request schemas.GetDomainInvitationRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Update Domain Invitation
// (PATCH /domains/{domain}/invitations/{invitation})
func (s *ApiV1) UpdateDomainInvitation(
	c *fiber.Ctx,
	request schemas.UpdateDomainInvitationRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Join Domain By Invitation
// (POST /domains/{domain}/join)
func (s *ApiV1) JoinDomainByInvitation(
	c *fiber.Ctx,
	request schemas.JoinDomainByInvitationRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// List Domain Roles
// (GET /domains/{domain}/roles)
func (s *ApiV1) ListDomainRoles(
	c *fiber.Ctx,
	request schemas.ListDomainRolesRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Create Domain Role
// (POST /domains/{domain}/roles)
func (s *ApiV1) CreateDomainRole(
	c *fiber.Ctx,
	request schemas.CreateDomainRoleRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Delete Domain Role
// (DELETE /domains/{domain}/roles/{role})
func (s *ApiV1) DeleteDomainRole(
	c *fiber.Ctx,
	request schemas.DeleteDomainRoleRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Get Domain Role
// (GET /domains/{domain}/roles/{role})
func (s *ApiV1) GetDomainRole(
	c *fiber.Ctx,
	request schemas.GetDomainRoleRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Update Domain Role
// (PATCH /domains/{domain}/roles/{role})
func (s *ApiV1) UpdateDomainRole(
	c *fiber.Ctx,
	request schemas.UpdateDomainRoleRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Transfer Domain
// (POST /domains/{domain}/transfer)
func (s *ApiV1) TransferDomain(
	c *fiber.Ctx,
	request schemas.TransferDomainRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// List Domain Users
// (GET /domains/{domain}/users)
func (s *ApiV1) ListDomainUsers(
	c *fiber.Ctx,
	request schemas.ListDomainUsersRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Add Domain User
// (POST /domains/{domain}/users)
func (s *ApiV1) AddDomainUser(
	c *fiber.Ctx,
	request schemas.AddDomainUserRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Remove Domain User
// (DELETE /domains/{domain}/users/{user})
func (s *ApiV1) RemoveDomainUser(
	c *fiber.Ctx,
	request schemas.RemoveDomainUserRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Get Domain User
// (GET /domains/{domain}/users/{user})
func (s *ApiV1) GetDomainUser(
	c *fiber.Ctx,
	request schemas.GetDomainUserRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Update Domain User
// (PATCH /domains/{domain}/users/{user})
func (s *ApiV1) UpdateDomainUser(
	c *fiber.Ctx,
	request schemas.UpdateDomainUserRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Get Domain User Permission
// (GET /domains/{domain}/users/{user}/permission)
func (s *ApiV1) GetDomainUserPermission(
	c *fiber.Ctx,
	request schemas.GetDomainUserPermissionRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}
