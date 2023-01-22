package handlers

import (
	"context"

	"github.com/joint-online-judge/go-horse/types"
)

type ApiV1 struct{}

// Admin List Domain Roles
// (GET /admin/domain_roles)
func (s *ApiV1) V1AdminListDomainRoles(ctx context.Context, request types.V1AdminListDomainRolesRequestObject) (any, error) {
	return nil, nil
}

// Admin List Judgers
// (GET /admin/judgers)
func (s *ApiV1) V1AdminListJudgers(ctx context.Context, request types.V1AdminListJudgersRequestObject) (any, error) {
	return nil, nil
}

// Admin Create Judger
// (POST /admin/judgers)
func (s *ApiV1) V1AdminCreateJudger(ctx context.Context, request types.V1AdminCreateJudgerRequestObject) (any, error) {
	return nil, nil
}

// Admin List Users
// (GET /admin/users)
func (s *ApiV1) V1AdminListUsers(ctx context.Context, request types.V1AdminListUsersRequestObject) (any, error) {
	return nil, nil
}

// Admin Get User
// (GET /admin/{uid})
func (s *ApiV1) V1AdminGetUser(ctx context.Context, request types.V1AdminGetUserRequestObject) (any, error) {
	return nil, nil
}

// Admin List User Domains
// (GET /admin/{uid}/domains)
func (s *ApiV1) V1AdminListUserDomains(ctx context.Context, request types.V1AdminListUserDomainsRequestObject) (any, error) {
	return nil, nil
}

// Login
// (POST /auth/login)
func (s *ApiV1) V1Login(ctx context.Context, request types.V1LoginRequestObject) (any, error) {
	return nil, nil
}

// Logout
// (POST /auth/logout)
func (s *ApiV1) V1Logout(ctx context.Context, request types.V1LogoutRequestObject) (any, error) {
	return nil, nil
}

// List Oauth2
// (GET /auth/oauth2)
func (s *ApiV1) V1ListOauth2(ctx context.Context, request types.V1ListOauth2RequestObject) (any, error) {
	return nil, nil
}

// Oauth Authorize
// (GET /auth/oauth2/{oauth2}/authorize)
func (s *ApiV1) V1OauthAuthorize(ctx context.Context, request types.V1OauthAuthorizeRequestObject) (any, error) {
	return nil, nil
}

// Refresh
// (POST /auth/refresh)
func (s *ApiV1) V1Refresh(ctx context.Context, request types.V1RefreshRequestObject) (any, error) {
	return nil, nil
}

// Register
// (POST /auth/register)
func (s *ApiV1) V1Register(ctx context.Context, request types.V1RegisterRequestObject) (any, error) {
	return nil, nil
}

// Get Token
// (GET /auth/token)
func (s *ApiV1) V1GetToken(ctx context.Context, request types.V1GetTokenRequestObject) (any, error) {
	return nil, nil
}

// Search Domain Groups
// (GET /domains/groups)
func (s *ApiV1) V1SearchDomainGroups(ctx context.Context, request types.V1SearchDomainGroupsRequestObject) (any, error) {
	return nil, nil
}

// Delete Domain
// (DELETE /domains/{domain})
func (s *ApiV1) V1DeleteDomain(ctx context.Context, request types.V1DeleteDomainRequestObject) (any, error) {
	return nil, nil
}

// Get Domain
// (GET /domains/{domain})
func (s *ApiV1) V1GetDomain(ctx context.Context, request types.V1GetDomainRequestObject) (any, error) {
	return nil, nil
}

// Update Domain
// (PATCH /domains/{domain})
func (s *ApiV1) V1UpdateDomain(ctx context.Context, request types.V1UpdateDomainRequestObject) (any, error) {
	return nil, nil
}

// Search Domain Candidates
// (GET /domains/{domain}/candidates)
func (s *ApiV1) V1SearchDomainCandidates(ctx context.Context, request types.V1SearchDomainCandidatesRequestObject) (any, error) {
	return nil, nil
}

// List Domain Invitations
// (GET /domains/{domain}/invitations)
func (s *ApiV1) V1ListDomainInvitations(ctx context.Context, request types.V1ListDomainInvitationsRequestObject) (any, error) {
	return nil, nil
}

// Create Domain Invitation
// (POST /domains/{domain}/invitations)
func (s *ApiV1) V1CreateDomainInvitation(ctx context.Context, request types.V1CreateDomainInvitationRequestObject) (any, error) {
	return nil, nil
}

// Delete Domain Invitation
// (DELETE /domains/{domain}/invitations/{invitation})
func (s *ApiV1) V1DeleteDomainInvitation(ctx context.Context, request types.V1DeleteDomainInvitationRequestObject) (any, error) {
	return nil, nil
}

// Get Domain Invitation
// (GET /domains/{domain}/invitations/{invitation})
func (s *ApiV1) V1GetDomainInvitation(ctx context.Context, request types.V1GetDomainInvitationRequestObject) (any, error) {
	return nil, nil
}

// Update Domain Invitation
// (PATCH /domains/{domain}/invitations/{invitation})
func (s *ApiV1) V1UpdateDomainInvitation(ctx context.Context, request types.V1UpdateDomainInvitationRequestObject) (any, error) {
	return nil, nil
}

// Join Domain By Invitation
// (POST /domains/{domain}/join)
func (s *ApiV1) V1JoinDomainByInvitation(ctx context.Context, request types.V1JoinDomainByInvitationRequestObject) (any, error) {
	return nil, nil
}

// List Problem Sets
// (GET /domains/{domain}/problem_sets)
func (s *ApiV1) V1ListProblemSets(ctx context.Context, request types.V1ListProblemSetsRequestObject) (any, error) {
	return nil, nil
}

// Create Problem Set
// (POST /domains/{domain}/problem_sets)
func (s *ApiV1) V1CreateProblemSet(ctx context.Context, request types.V1CreateProblemSetRequestObject) (any, error) {
	return nil, nil
}

// Delete Problem Set
// (DELETE /domains/{domain}/problem_sets/{problemSet})
func (s *ApiV1) V1DeleteProblemSet(ctx context.Context, request types.V1DeleteProblemSetRequestObject) (any, error) {
	return nil, nil
}

// Get Problem Set
// (GET /domains/{domain}/problem_sets/{problemSet})
func (s *ApiV1) V1GetProblemSet(ctx context.Context, request types.V1GetProblemSetRequestObject) (any, error) {
	return nil, nil
}

// Update Problem Set
// (PATCH /domains/{domain}/problem_sets/{problemSet})
func (s *ApiV1) V1UpdateProblemSet(ctx context.Context, request types.V1UpdateProblemSetRequestObject) (any, error) {
	return nil, nil
}

// List Problems In Problem Set
// (GET /domains/{domain}/problem_sets/{problemSet}/problems)
func (s *ApiV1) V1ListProblemsInProblemSet(ctx context.Context, request types.V1ListProblemsInProblemSetRequestObject) (any, error) {
	return nil, nil
}

// Add Problem In Problem Set
// (POST /domains/{domain}/problem_sets/{problemSet}/problems)
func (s *ApiV1) V1AddProblemInProblemSet(ctx context.Context, request types.V1AddProblemInProblemSetRequestObject) (any, error) {
	return nil, nil
}

// Delete Problem In Problem Set
// (DELETE /domains/{domain}/problem_sets/{problemSet}/problems/{problem})
func (s *ApiV1) V1DeleteProblemInProblemSet(ctx context.Context, request types.V1DeleteProblemInProblemSetRequestObject) (any, error) {
	return nil, nil
}

// Get Problem In Problem Set
// (GET /domains/{domain}/problem_sets/{problemSet}/problems/{problem})
func (s *ApiV1) V1GetProblemInProblemSet(ctx context.Context, request types.V1GetProblemInProblemSetRequestObject) (any, error) {
	return nil, nil
}

// Update Problem In Problem Set
// (PATCH /domains/{domain}/problem_sets/{problemSet}/problems/{problem})
func (s *ApiV1) V1UpdateProblemInProblemSet(ctx context.Context, request types.V1UpdateProblemInProblemSetRequestObject) (any, error) {
	return nil, nil
}

// Submit Solution To Problem Set
// (POST /domains/{domain}/problem_sets/{problemSet}/problems/{problem}/submit)
func (s *ApiV1) V1SubmitSolutionToProblemSet(ctx context.Context, request types.V1SubmitSolutionToProblemSetRequestObject) (any, error) {
	return nil, nil
}

// List Problems
// (GET /domains/{domain}/problems)
func (s *ApiV1) V1ListProblems(ctx context.Context, request types.V1ListProblemsRequestObject) (any, error) {
	return nil, nil
}

// Create Problem
// (POST /domains/{domain}/problems)
func (s *ApiV1) V1CreateProblem(ctx context.Context, request types.V1CreateProblemRequestObject) (any, error) {
	return nil, nil
}

// Clone Problem
// (POST /domains/{domain}/problems/clone)
func (s *ApiV1) V1CloneProblem(ctx context.Context, request types.V1CloneProblemRequestObject) (any, error) {
	return nil, nil
}

// Delete Problem
// (DELETE /domains/{domain}/problems/{problem})
func (s *ApiV1) V1DeleteProblem(ctx context.Context, request types.V1DeleteProblemRequestObject) (any, error) {
	return nil, nil
}

// Get Problem
// (GET /domains/{domain}/problems/{problem})
func (s *ApiV1) V1GetProblem(ctx context.Context, request types.V1GetProblemRequestObject) (any, error) {
	return nil, nil
}

// Update Problem
// (PATCH /domains/{domain}/problems/{problem})
func (s *ApiV1) V1UpdateProblem(ctx context.Context, request types.V1UpdateProblemRequestObject) (any, error) {
	return nil, nil
}

// Submit Solution To Problem
// (POST /domains/{domain}/problems/{problem})
func (s *ApiV1) V1SubmitSolutionToProblem(ctx context.Context, request types.V1SubmitSolutionToProblemRequestObject) (any, error) {
	return nil, nil
}

// List Problem Config Commits
// (GET /domains/{domain}/problems/{problem}/configs)
func (s *ApiV1) V1ListProblemConfigCommits(ctx context.Context, request types.V1ListProblemConfigCommitsRequestObject) (any, error) {
	return nil, nil
}

// Update Problem Config By Archive
// (POST /domains/{domain}/problems/{problem}/configs)
func (s *ApiV1) V1UpdateProblemConfigByArchive(ctx context.Context, request types.V1UpdateProblemConfigByArchiveRequestObject) (any, error) {
	return nil, nil
}

// Update Problem Config Json
// (POST /domains/{domain}/problems/{problem}/configs/json)
func (s *ApiV1) V1UpdateProblemConfigJson(ctx context.Context, request types.V1UpdateProblemConfigJsonRequestObject) (any, error) {
	return nil, nil
}

// Diff Problem Config Default Branch
// (GET /domains/{domain}/problems/{problem}/configs/latest/diff)
func (s *ApiV1) V1DiffProblemConfigDefaultBranch(ctx context.Context, request types.V1DiffProblemConfigDefaultBranchRequestObject) (any, error) {
	return nil, nil
}

// List Latest Problem Config Objects Under A Given Prefix
// (GET /domains/{domain}/problems/{problem}/configs/latest/ls)
func (s *ApiV1) V1ListLatestProblemConfigObjectsUnderAGivenPrefix(ctx context.Context, request types.V1ListLatestProblemConfigObjectsUnderAGivenPrefixRequestObject) (any, error) {
	return nil, nil
}

// Download Problem Config Archive
// (GET /domains/{domain}/problems/{problem}/configs/{config})
func (s *ApiV1) V1DownloadProblemConfigArchive(ctx context.Context, request types.V1DownloadProblemConfigArchiveRequestObject) (any, error) {
	return nil, nil
}

// Get Problem Config Json
// (GET /domains/{domain}/problems/{problem}/configs/{config}/json)
func (s *ApiV1) V1GetProblemConfigJson(ctx context.Context, request types.V1GetProblemConfigJsonRequestObject) (any, error) {
	return nil, nil
}

// List Records In Domain
// (GET /domains/{domain}/records)
func (s *ApiV1) V1ListRecordsInDomain(ctx context.Context, request types.V1ListRecordsInDomainRequestObject) (any, error) {
	return nil, nil
}

// Get Record
// (GET /domains/{domain}/records/{record})
func (s *ApiV1) V1GetRecord(ctx context.Context, request types.V1GetRecordRequestObject) (any, error) {
	return nil, nil
}

// Submit Case By Judger
// (PUT /domains/{domain}/records/{record}/cases/{index}/judge)
func (s *ApiV1) V1SubmitCaseByJudger(ctx context.Context, request types.V1SubmitCaseByJudgerRequestObject) (any, error) {
	return nil, nil
}

// Submit Record By Judger
// (PUT /domains/{domain}/records/{record}/judge)
func (s *ApiV1) V1SubmitRecordByJudger(ctx context.Context, request types.V1SubmitRecordByJudgerRequestObject) (any, error) {
	return nil, nil
}

// Claim Record By Judger
// (POST /domains/{domain}/records/{record}/judge/claim)
func (s *ApiV1) V1ClaimRecordByJudger(ctx context.Context, request types.V1ClaimRecordByJudgerRequestObject) (any, error) {
	return nil, nil
}

// List Domain Roles
// (GET /domains/{domain}/roles)
func (s *ApiV1) V1ListDomainRoles(ctx context.Context, request types.V1ListDomainRolesRequestObject) (any, error) {
	return nil, nil
}

// Create Domain Role
// (POST /domains/{domain}/roles)
func (s *ApiV1) V1CreateDomainRole(ctx context.Context, request types.V1CreateDomainRoleRequestObject) (any, error) {
	return nil, nil
}

// Delete Domain Role
// (DELETE /domains/{domain}/roles/{role})
func (s *ApiV1) V1DeleteDomainRole(ctx context.Context, request types.V1DeleteDomainRoleRequestObject) (any, error) {
	return nil, nil
}

// Get Domain Role
// (GET /domains/{domain}/roles/{role})
func (s *ApiV1) V1GetDomainRole(ctx context.Context, request types.V1GetDomainRoleRequestObject) (any, error) {
	return nil, nil
}

// Update Domain Role
// (PATCH /domains/{domain}/roles/{role})
func (s *ApiV1) V1UpdateDomainRole(ctx context.Context, request types.V1UpdateDomainRoleRequestObject) (any, error) {
	return nil, nil
}

// Transfer Domain
// (POST /domains/{domain}/transfer)
func (s *ApiV1) V1TransferDomain(ctx context.Context, request types.V1TransferDomainRequestObject) (any, error) {
	return nil, nil
}

// List Domain Users
// (GET /domains/{domain}/users)
func (s *ApiV1) V1ListDomainUsers(ctx context.Context, request types.V1ListDomainUsersRequestObject) (any, error) {
	return nil, nil
}

// Add Domain User
// (POST /domains/{domain}/users)
func (s *ApiV1) V1AddDomainUser(ctx context.Context, request types.V1AddDomainUserRequestObject) (any, error) {
	return nil, nil
}

// Remove Domain User
// (DELETE /domains/{domain}/users/{user})
func (s *ApiV1) V1RemoveDomainUser(ctx context.Context, request types.V1RemoveDomainUserRequestObject) (any, error) {
	return nil, nil
}

// Get Domain User
// (GET /domains/{domain}/users/{user})
func (s *ApiV1) V1GetDomainUser(ctx context.Context, request types.V1GetDomainUserRequestObject) (any, error) {
	return nil, nil
}

// Update Domain User
// (PATCH /domains/{domain}/users/{user})
func (s *ApiV1) V1UpdateDomainUser(ctx context.Context, request types.V1UpdateDomainUserRequestObject) (any, error) {
	return nil, nil
}

// Get Domain User Permission
// (GET /domains/{domain}/users/{user}/permission)
func (s *ApiV1) V1GetDomainUserPermission(ctx context.Context, request types.V1GetDomainUserPermissionRequestObject) (any, error) {
	return nil, nil
}

// Jwt Decoded
// (GET /jwt_decoded)
func (s *ApiV1) V1JwtDecoded(ctx context.Context, request types.V1JwtDecodedRequestObject) (any, error) {
	return nil, nil
}

// List Problem Groups
// (GET /problem_groups)
func (s *ApiV1) V1ListProblemGroups(ctx context.Context, request types.V1ListProblemGroupsRequestObject) (any, error) {
	return nil, nil
}

// Test Error Report
// (GET /test/report)
func (s *ApiV1) V1TestErrorReport(ctx context.Context, request types.V1TestErrorReportRequestObject) (any, error) {
	return nil, nil
}

// Get Current User
// (GET /users/me)
func (s *ApiV1) V1GetCurrentUser(ctx context.Context, request types.V1GetCurrentUserRequestObject) (any, error) {
	return nil, nil
}

// Update Current User
// (PATCH /users/me)
func (s *ApiV1) V1UpdateCurrentUser(ctx context.Context, request types.V1UpdateCurrentUserRequestObject) (any, error) {
	return nil, nil
}

// Change Password
// (PATCH /users/me/password)
func (s *ApiV1) V1ChangePassword(ctx context.Context, request types.V1ChangePasswordRequestObject) (any, error) {
	return nil, nil
}

// Get User
// (GET /users/{uid})
func (s *ApiV1) V1GetUser(ctx context.Context, request types.V1GetUserRequestObject) (any, error) {
	return nil, nil
}

// Version
// (GET /version)
func (s *ApiV1) V1Version(ctx context.Context, request types.V1VersionRequestObject) (any, error) {
	return nil, nil
}
