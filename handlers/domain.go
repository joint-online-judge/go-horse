package handlers

import (
	"context"

	"github.com/joint-online-judge/go-horse/handlers/utils"
	"github.com/joint-online-judge/go-horse/models"
	"github.com/joint-online-judge/go-horse/schemas"
	log "github.com/sirupsen/logrus"
)

// List Domains
// (GET /domains)
func (s *ApiV1) ListDomains(
	ctx context.Context,
	request schemas.ListDomainsRequestObject,
) (any, error) {
	return utils.NewListResp[models.Domain, schemas.Domain]()
}

// Create Domain
// (POST /domains)
func (s *ApiV1) CreateDomain(
	ctx context.Context,
	request schemas.CreateDomainRequestObject,
) (any, error) {
	b := request.Body
	log.Infof("request.Body: %v", b)
	return nil, nil
}

// Get Domain
// (GET /domains/{domain})
func (s *ApiV1) GetDomain(
	ctx context.Context,
	request schemas.GetDomainRequestObject,
) (any, error) {
	log.Info("get domain")
	return nil, nil
}

// Search Domain Groups
// (GET /domains/groups)
func (s *ApiV1) SearchDomainGroups(
	ctx context.Context,
	request schemas.SearchDomainGroupsRequestObject,
) (any, error) {
	return nil, nil
}

// Delete Domain
// (DELETE /domains/{domain})
func (s *ApiV1) DeleteDomain(
	ctx context.Context,
	request schemas.DeleteDomainRequestObject,
) (any, error) {
	return nil, nil
}

// Update Domain
// (PATCH /domains/{domain})
func (s *ApiV1) UpdateDomain(
	ctx context.Context,
	request schemas.UpdateDomainRequestObject,
) (any, error) {
	return nil, nil
}

// Search Domain Candidates
// (GET /domains/{domain}/candidates)
func (s *ApiV1) SearchDomainCandidates(
	ctx context.Context,
	request schemas.SearchDomainCandidatesRequestObject,
) (any, error) {
	return nil, nil
}

// List Domain Invitations
// (GET /domains/{domain}/invitations)
func (s *ApiV1) ListDomainInvitations(
	ctx context.Context,
	request schemas.ListDomainInvitationsRequestObject,
) (any, error) {
	return nil, nil
}

// Create Domain Invitation
// (POST /domains/{domain}/invitations)
func (s *ApiV1) CreateDomainInvitation(
	ctx context.Context,
	request schemas.CreateDomainInvitationRequestObject,
) (any, error) {
	return nil, nil
}

// Delete Domain Invitation
// (DELETE /domains/{domain}/invitations/{invitation})
func (s *ApiV1) DeleteDomainInvitation(
	ctx context.Context,
	request schemas.DeleteDomainInvitationRequestObject,
) (any, error) {
	return nil, nil
}

// Get Domain Invitation
// (GET /domains/{domain}/invitations/{invitation})
func (s *ApiV1) GetDomainInvitation(
	ctx context.Context,
	request schemas.GetDomainInvitationRequestObject,
) (any, error) {
	return nil, nil
}

// Update Domain Invitation
// (PATCH /domains/{domain}/invitations/{invitation})
func (s *ApiV1) UpdateDomainInvitation(
	ctx context.Context,
	request schemas.UpdateDomainInvitationRequestObject,
) (any, error) {
	return nil, nil
}

// Join Domain By Invitation
// (POST /domains/{domain}/join)
func (s *ApiV1) JoinDomainByInvitation(
	ctx context.Context,
	request schemas.JoinDomainByInvitationRequestObject,
) (any, error) {
	return nil, nil
}

// List Domain Roles
// (GET /domains/{domain}/roles)
func (s *ApiV1) ListDomainRoles(
	ctx context.Context,
	request schemas.ListDomainRolesRequestObject,
) (any, error) {
	return nil, nil
}

// Create Domain Role
// (POST /domains/{domain}/roles)
func (s *ApiV1) CreateDomainRole(
	ctx context.Context,
	request schemas.CreateDomainRoleRequestObject,
) (any, error) {
	return nil, nil
}

// Delete Domain Role
// (DELETE /domains/{domain}/roles/{role})
func (s *ApiV1) DeleteDomainRole(
	ctx context.Context,
	request schemas.DeleteDomainRoleRequestObject,
) (any, error) {
	return nil, nil
}

// Get Domain Role
// (GET /domains/{domain}/roles/{role})
func (s *ApiV1) GetDomainRole(
	ctx context.Context,
	request schemas.GetDomainRoleRequestObject,
) (any, error) {
	return nil, nil
}

// Update Domain Role
// (PATCH /domains/{domain}/roles/{role})
func (s *ApiV1) UpdateDomainRole(
	ctx context.Context,
	request schemas.UpdateDomainRoleRequestObject,
) (any, error) {
	return nil, nil
}

// Transfer Domain
// (POST /domains/{domain}/transfer)
func (s *ApiV1) TransferDomain(
	ctx context.Context,
	request schemas.TransferDomainRequestObject,
) (any, error) {
	return nil, nil
}

// List Domain Users
// (GET /domains/{domain}/users)
func (s *ApiV1) ListDomainUsers(
	ctx context.Context,
	request schemas.ListDomainUsersRequestObject,
) (any, error) {
	return nil, nil
}

// Add Domain User
// (POST /domains/{domain}/users)
func (s *ApiV1) AddDomainUser(
	ctx context.Context,
	request schemas.AddDomainUserRequestObject,
) (any, error) {
	return nil, nil
}

// Remove Domain User
// (DELETE /domains/{domain}/users/{user})
func (s *ApiV1) RemoveDomainUser(
	ctx context.Context,
	request schemas.RemoveDomainUserRequestObject,
) (any, error) {
	return nil, nil
}

// Get Domain User
// (GET /domains/{domain}/users/{user})
func (s *ApiV1) GetDomainUser(
	ctx context.Context,
	request schemas.GetDomainUserRequestObject,
) (any, error) {
	return nil, nil
}

// Update Domain User
// (PATCH /domains/{domain}/users/{user})
func (s *ApiV1) UpdateDomainUser(
	ctx context.Context,
	request schemas.UpdateDomainUserRequestObject,
) (any, error) {
	return nil, nil
}

// Get Domain User Permission
// (GET /domains/{domain}/users/{user}/permission)
func (s *ApiV1) GetDomainUserPermission(
	ctx context.Context,
	request schemas.GetDomainUserPermissionRequestObject,
) (any, error) {
	return nil, nil
}
