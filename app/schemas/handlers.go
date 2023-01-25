// Package schemas provides primitives to interact with the openapi HTTP API.
package schemas

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/url"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/gofiber/fiber/v2"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Admin List Domain Roles
	// (GET /admin/domain_roles)
	AdminListDomainRoles(c *fiber.Ctx, params AdminListDomainRolesParams) error
	// Admin List Judgers
	// (GET /admin/judgers)
	AdminListJudgers(c *fiber.Ctx, params AdminListJudgersParams) error
	// Admin Create Judger
	// (POST /admin/judgers)
	AdminCreateJudger(c *fiber.Ctx) error
	// Admin List Users
	// (GET /admin/users)
	AdminListUsers(c *fiber.Ctx, params AdminListUsersParams) error
	// Admin Get User
	// (GET /admin/{uid})
	AdminGetUser(c *fiber.Ctx, uid string) error
	// Admin List User Domains
	// (GET /admin/{uid}/domains)
	AdminListUserDomains(c *fiber.Ctx, uid string, params AdminListUserDomainsParams) error
	// Login
	// (POST /auth/login)
	Login(c *fiber.Ctx, params LoginParams) error
	// Logout
	// (POST /auth/logout)
	Logout(c *fiber.Ctx, params LogoutParams) error
	// List Oauth2
	// (GET /auth/oauth2)
	ListOauth2(c *fiber.Ctx) error
	// Oauth Authorize
	// (GET /auth/oauth2/{oauth2}/authorize)
	OauthAuthorize(c *fiber.Ctx, oauth2 string, params OauthAuthorizeParams) error
	// Refresh
	// (POST /auth/refresh)
	Refresh(c *fiber.Ctx, params RefreshParams) error
	// Register
	// (POST /auth/register)
	Register(c *fiber.Ctx, params RegisterParams) error
	// Get Token
	// (GET /auth/token)
	GetToken(c *fiber.Ctx, params GetTokenParams) error
	// List Domains
	// (GET /domains)
	ListDomains(c *fiber.Ctx, params ListDomainsParams) error
	// Create Domain
	// (POST /domains)
	CreateDomain(c *fiber.Ctx) error
	// Search Domain Groups
	// (GET /domains/groups)
	SearchDomainGroups(c *fiber.Ctx, params SearchDomainGroupsParams) error
	// Delete Domain
	// (DELETE /domains/{domain})
	DeleteDomain(c *fiber.Ctx, domain string) error
	// Get Domain
	// (GET /domains/{domain})
	GetDomain(c *fiber.Ctx, domain string) error
	// Update Domain
	// (PATCH /domains/{domain})
	UpdateDomain(c *fiber.Ctx, domain string) error
	// Search Domain Candidates
	// (GET /domains/{domain}/candidates)
	SearchDomainCandidates(c *fiber.Ctx, domain string, params SearchDomainCandidatesParams) error
	// List Domain Invitations
	// (GET /domains/{domain}/invitations)
	ListDomainInvitations(c *fiber.Ctx, domain string, params ListDomainInvitationsParams) error
	// Create Domain Invitation
	// (POST /domains/{domain}/invitations)
	CreateDomainInvitation(c *fiber.Ctx, domain string) error
	// Delete Domain Invitation
	// (DELETE /domains/{domain}/invitations/{invitation})
	DeleteDomainInvitation(c *fiber.Ctx, domain string, invitation string) error
	// Get Domain Invitation
	// (GET /domains/{domain}/invitations/{invitation})
	GetDomainInvitation(c *fiber.Ctx, domain string, invitation string) error
	// Update Domain Invitation
	// (PATCH /domains/{domain}/invitations/{invitation})
	UpdateDomainInvitation(c *fiber.Ctx, domain string, invitation string) error
	// Join Domain By Invitation
	// (POST /domains/{domain}/join)
	JoinDomainByInvitation(c *fiber.Ctx, domain string, params JoinDomainByInvitationParams) error
	// List Problem Sets
	// (GET /domains/{domain}/problem_sets)
	ListProblemSets(c *fiber.Ctx, domain string, params ListProblemSetsParams) error
	// Create Problem Set
	// (POST /domains/{domain}/problem_sets)
	CreateProblemSet(c *fiber.Ctx, domain string) error
	// Delete Problem Set
	// (DELETE /domains/{domain}/problem_sets/{problemSet})
	DeleteProblemSet(c *fiber.Ctx, domain string, problemSet string) error
	// Get Problem Set
	// (GET /domains/{domain}/problem_sets/{problemSet})
	GetProblemSet(c *fiber.Ctx, domain string, problemSet string) error
	// Update Problem Set
	// (PATCH /domains/{domain}/problem_sets/{problemSet})
	UpdateProblemSet(c *fiber.Ctx, domain string, problemSet string) error
	// List Problems In Problem Set
	// (GET /domains/{domain}/problem_sets/{problemSet}/problems)
	ListProblemsInProblemSet(c *fiber.Ctx, domain string, problemSet string) error
	// Add Problem In Problem Set
	// (POST /domains/{domain}/problem_sets/{problemSet}/problems)
	AddProblemInProblemSet(c *fiber.Ctx, domain string, problemSet string) error
	// Delete Problem In Problem Set
	// (DELETE /domains/{domain}/problem_sets/{problemSet}/problems/{problem})
	DeleteProblemInProblemSet(c *fiber.Ctx, domain string, problemSet string, problem string) error
	// Get Problem In Problem Set
	// (GET /domains/{domain}/problem_sets/{problemSet}/problems/{problem})
	GetProblemInProblemSet(c *fiber.Ctx, domain string, problemSet string, problem string) error
	// Update Problem In Problem Set
	// (PATCH /domains/{domain}/problem_sets/{problemSet}/problems/{problem})
	UpdateProblemInProblemSet(c *fiber.Ctx, domain string, problemSet string, problem string) error
	// Submit Solution To Problem Set
	// (POST /domains/{domain}/problem_sets/{problemSet}/problems/{problem}/submit)
	SubmitSolutionToProblemSet(c *fiber.Ctx, domain string, problemSet string, problem string) error
	// List Problems
	// (GET /domains/{domain}/problems)
	ListProblems(c *fiber.Ctx, domain string, params ListProblemsParams) error
	// Create Problem
	// (POST /domains/{domain}/problems)
	CreateProblem(c *fiber.Ctx, domain string) error
	// Clone Problem
	// (POST /domains/{domain}/problems/clone)
	CloneProblem(c *fiber.Ctx, domain string) error
	// Delete Problem
	// (DELETE /domains/{domain}/problems/{problem})
	DeleteProblem(c *fiber.Ctx, domain string, problem string) error
	// Get Problem
	// (GET /domains/{domain}/problems/{problem})
	GetProblem(c *fiber.Ctx, domain string, problem string) error
	// Update Problem
	// (PATCH /domains/{domain}/problems/{problem})
	UpdateProblem(c *fiber.Ctx, domain string, problem string) error
	// Submit Solution To Problem
	// (POST /domains/{domain}/problems/{problem})
	SubmitSolutionToProblem(c *fiber.Ctx, domain string, problem string) error
	// List Problem Config Commits
	// (GET /domains/{domain}/problems/{problem}/configs)
	ListProblemConfigCommits(
		c *fiber.Ctx,
		domain string,
		problem string,
		params ListProblemConfigCommitsParams,
	) error
	// Update Problem Config By Archive
	// (POST /domains/{domain}/problems/{problem}/configs)
	UpdateProblemConfigByArchive(
		c *fiber.Ctx,
		domain string,
		problem string,
		params UpdateProblemConfigByArchiveParams,
	) error
	// Update Problem Config Json
	// (POST /domains/{domain}/problems/{problem}/configs/json)
	UpdateProblemConfigJson(c *fiber.Ctx, domain string, problem string) error
	// Diff Problem Config Default Branch
	// (GET /domains/{domain}/problems/{problem}/configs/latest/diff)
	DiffProblemConfigDefaultBranch(
		c *fiber.Ctx,
		domain string,
		problem string,
		params DiffProblemConfigDefaultBranchParams,
	) error
	// List Latest Problem Config Objects Under A Given Prefix
	// (GET /domains/{domain}/problems/{problem}/configs/latest/ls)
	ListLatestProblemConfigObjectsUnderAGivenPrefix(
		c *fiber.Ctx,
		domain string,
		problem string,
		params ListLatestProblemConfigObjectsUnderAGivenPrefixParams,
	) error
	// Download Problem Config Archive
	// (GET /domains/{domain}/problems/{problem}/configs/{config})
	DownloadProblemConfigArchive(
		c *fiber.Ctx,
		domain string,
		problem string,
		config string,
		params DownloadProblemConfigArchiveParams,
	) error
	// Get Problem Config Json
	// (GET /domains/{domain}/problems/{problem}/configs/{config}/json)
	GetProblemConfigJson(c *fiber.Ctx, domain string, problem string, config string) error
	// List Records In Domain
	// (GET /domains/{domain}/records)
	ListRecordsInDomain(c *fiber.Ctx, domain string, params ListRecordsInDomainParams) error
	// Get Record
	// (GET /domains/{domain}/records/{record})
	GetRecord(c *fiber.Ctx, domain string, record openapi_types.UUID) error
	// Submit Case By Judger
	// (PUT /domains/{domain}/records/{record}/cases/{index}/judge)
	SubmitCaseByJudger(c *fiber.Ctx, domain string, record string, index int) error
	// Submit Record By Judger
	// (PUT /domains/{domain}/records/{record}/judge)
	SubmitRecordByJudger(c *fiber.Ctx, domain string, record string) error
	// Claim Record By Judger
	// (POST /domains/{domain}/records/{record}/judge/claim)
	ClaimRecordByJudger(c *fiber.Ctx, domain string, record string) error
	// List Domain Roles
	// (GET /domains/{domain}/roles)
	ListDomainRoles(c *fiber.Ctx, domain string, params ListDomainRolesParams) error
	// Create Domain Role
	// (POST /domains/{domain}/roles)
	CreateDomainRole(c *fiber.Ctx, domain string) error
	// Delete Domain Role
	// (DELETE /domains/{domain}/roles/{role})
	DeleteDomainRole(c *fiber.Ctx, domain string, role string) error
	// Get Domain Role
	// (GET /domains/{domain}/roles/{role})
	GetDomainRole(c *fiber.Ctx, domain string, role string) error
	// Update Domain Role
	// (PATCH /domains/{domain}/roles/{role})
	UpdateDomainRole(c *fiber.Ctx, domain string, role string) error
	// Transfer Domain
	// (POST /domains/{domain}/transfer)
	TransferDomain(c *fiber.Ctx, domain string) error
	// List Domain Users
	// (GET /domains/{domain}/users)
	ListDomainUsers(c *fiber.Ctx, domain string, params ListDomainUsersParams) error
	// Add Domain User
	// (POST /domains/{domain}/users)
	AddDomainUser(c *fiber.Ctx, domain string) error
	// Remove Domain User
	// (DELETE /domains/{domain}/users/{user})
	RemoveDomainUser(c *fiber.Ctx, domain string, user string) error
	// Get Domain User
	// (GET /domains/{domain}/users/{user})
	GetDomainUser(c *fiber.Ctx, domain string, user string) error
	// Update Domain User
	// (PATCH /domains/{domain}/users/{user})
	UpdateDomainUser(c *fiber.Ctx, domain string, user string) error
	// Get Domain User Permission
	// (GET /domains/{domain}/users/{user}/permission)
	GetDomainUserPermission(c *fiber.Ctx, domain string, user string) error
	// Jwt Decoded
	// (GET /jwt_decoded)
	JwtDecoded(c *fiber.Ctx) error
	// List Problem Groups
	// (GET /problem_groups)
	ListProblemGroups(c *fiber.Ctx, params ListProblemGroupsParams) error
	// Test Error Report
	// (GET /test/report)
	TestErrorReport(c *fiber.Ctx) error
	// Get Current User
	// (GET /users/me)
	GetCurrentUser(c *fiber.Ctx) error
	// Update Current User
	// (PATCH /users/me)
	UpdateCurrentUser(c *fiber.Ctx) error
	// Change Password
	// (PATCH /users/me/password)
	ChangePassword(c *fiber.Ctx) error
	// Get User
	// (GET /users/{uid})
	GetUser(c *fiber.Ctx, uid string) error
	// Version
	// (GET /version)
	Version(c *fiber.Ctx) error
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

type MiddlewareFunc fiber.Handler

// AdminListDomainRoles operation middleware
func (siw *ServerInterfaceWrapper) AdminListDomainRoles(c *fiber.Ctx) error {
	var err error

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params AdminListDomainRolesParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "ordering" -------------

	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter ordering: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter offset: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter limit: %w", err).Error(),
		)
	}

	return siw.Handler.AdminListDomainRoles(c, params)
}

// AdminListJudgers operation middleware
func (siw *ServerInterfaceWrapper) AdminListJudgers(c *fiber.Ctx) error {
	var err error

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params AdminListJudgersParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "ordering" -------------

	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter ordering: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter offset: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter limit: %w", err).Error(),
		)
	}

	return siw.Handler.AdminListJudgers(c, params)
}

// AdminCreateJudger operation middleware
func (siw *ServerInterfaceWrapper) AdminCreateJudger(c *fiber.Ctx) error {
	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.AdminCreateJudger(c)
}

// AdminListUsers operation middleware
func (siw *ServerInterfaceWrapper) AdminListUsers(c *fiber.Ctx) error {
	var err error

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params AdminListUsersParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "ordering" -------------

	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter ordering: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter offset: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter limit: %w", err).Error(),
		)
	}

	return siw.Handler.AdminListUsers(c, params)
}

// AdminGetUser operation middleware
func (siw *ServerInterfaceWrapper) AdminGetUser(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "uid" -------------
	var uid string

	err = runtime.BindStyledParameter("simple", false, "uid", c.Params("uid"), &uid)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter uid: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.AdminGetUser(c, uid)
}

// AdminListUserDomains operation middleware
func (siw *ServerInterfaceWrapper) AdminListUserDomains(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "uid" -------------
	var uid string

	err = runtime.BindStyledParameter("simple", false, "uid", c.Params("uid"), &uid)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter uid: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params AdminListUserDomainsParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "role" -------------

	err = runtime.BindQueryParameter("form", true, false, "role", query, &params.Role)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter role: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "groups" -------------

	err = runtime.BindQueryParameter("form", true, false, "groups", query, &params.Groups)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter groups: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "ordering" -------------

	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter ordering: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter offset: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter limit: %w", err).Error(),
		)
	}

	return siw.Handler.AdminListUserDomains(c, uid, params)
}

// Login operation middleware
func (siw *ServerInterfaceWrapper) Login(c *fiber.Ctx) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params LoginParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "cookie" -------------

	err = runtime.BindQueryParameter("form", true, false, "cookie", query, &params.Cookie)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter cookie: %w", err).Error(),
		)
	}

	// ------------- Required query parameter "responseType" -------------

	if paramValue := c.Query("responseType"); paramValue != "" {
	} else {
		err = fmt.Errorf("query argument responseType is required, but not found")
		return fiber.NewError(
			fiber.StatusBadRequest,
			err.Error(),
		)
	}

	err = runtime.BindQueryParameter(
		"form",
		true,
		true,
		"responseType",
		query,
		&params.ResponseType,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter responseType: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "redirectUrl" -------------

	err = runtime.BindQueryParameter("form", true, false, "redirectUrl", query, &params.RedirectUrl)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter redirectUrl: %w", err).Error(),
		)
	}

	return siw.Handler.Login(c, params)
}

// Logout operation middleware
func (siw *ServerInterfaceWrapper) Logout(c *fiber.Ctx) error {
	var err error

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params LogoutParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "cookie" -------------

	err = runtime.BindQueryParameter("form", true, false, "cookie", query, &params.Cookie)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter cookie: %w", err).Error(),
		)
	}

	// ------------- Required query parameter "responseType" -------------

	if paramValue := c.Query("responseType"); paramValue != "" {
	} else {
		err = fmt.Errorf("query argument responseType is required, but not found")
		return fiber.NewError(
			fiber.StatusBadRequest,
			err.Error(),
		)
	}

	err = runtime.BindQueryParameter(
		"form",
		true,
		true,
		"responseType",
		query,
		&params.ResponseType,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter responseType: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "redirectUrl" -------------

	err = runtime.BindQueryParameter("form", true, false, "redirectUrl", query, &params.RedirectUrl)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter redirectUrl: %w", err).Error(),
		)
	}

	return siw.Handler.Logout(c, params)
}

// ListOauth2 operation middleware
func (siw *ServerInterfaceWrapper) ListOauth2(c *fiber.Ctx) error {
	return siw.Handler.ListOauth2(c)
}

// OauthAuthorize operation middleware
func (siw *ServerInterfaceWrapper) OauthAuthorize(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "oauth2" -------------
	var oauth2 string

	err = runtime.BindStyledParameter("simple", false, "oauth2", c.Params("oauth2"), &oauth2)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter oauth2: %w", err).Error(),
		)
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params OauthAuthorizeParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "scopes" -------------

	err = runtime.BindQueryParameter("form", true, false, "scopes", query, &params.Scopes)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter scopes: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "cookie" -------------

	err = runtime.BindQueryParameter("form", true, false, "cookie", query, &params.Cookie)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter cookie: %w", err).Error(),
		)
	}

	// ------------- Required query parameter "responseType" -------------

	if paramValue := c.Query("responseType"); paramValue != "" {
	} else {
		err = fmt.Errorf("query argument responseType is required, but not found")
		return fiber.NewError(
			fiber.StatusBadRequest,
			err.Error(),
		)
	}

	err = runtime.BindQueryParameter(
		"form",
		true,
		true,
		"responseType",
		query,
		&params.ResponseType,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter responseType: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "redirectUrl" -------------

	err = runtime.BindQueryParameter("form", true, false, "redirectUrl", query, &params.RedirectUrl)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter redirectUrl: %w", err).Error(),
		)
	}

	return siw.Handler.OauthAuthorize(c, oauth2, params)
}

// Refresh operation middleware
func (siw *ServerInterfaceWrapper) Refresh(c *fiber.Ctx) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params RefreshParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "cookie" -------------

	err = runtime.BindQueryParameter("form", true, false, "cookie", query, &params.Cookie)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter cookie: %w", err).Error(),
		)
	}

	// ------------- Required query parameter "responseType" -------------

	if paramValue := c.Query("responseType"); paramValue != "" {
	} else {
		err = fmt.Errorf("query argument responseType is required, but not found")
		return fiber.NewError(
			fiber.StatusBadRequest,
			err.Error(),
		)
	}

	err = runtime.BindQueryParameter(
		"form",
		true,
		true,
		"responseType",
		query,
		&params.ResponseType,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter responseType: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "redirectUrl" -------------

	err = runtime.BindQueryParameter("form", true, false, "redirectUrl", query, &params.RedirectUrl)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter redirectUrl: %w", err).Error(),
		)
	}

	return siw.Handler.Refresh(c, params)
}

// Register operation middleware
func (siw *ServerInterfaceWrapper) Register(c *fiber.Ctx) error {
	var err error

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params RegisterParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "cookie" -------------

	err = runtime.BindQueryParameter("form", true, false, "cookie", query, &params.Cookie)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter cookie: %w", err).Error(),
		)
	}

	// ------------- Required query parameter "responseType" -------------

	if paramValue := c.Query("responseType"); paramValue != "" {
	} else {
		err = fmt.Errorf("query argument responseType is required, but not found")
		return fiber.NewError(
			fiber.StatusBadRequest,
			err.Error(),
		)
	}

	err = runtime.BindQueryParameter(
		"form",
		true,
		true,
		"responseType",
		query,
		&params.ResponseType,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter responseType: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "redirectUrl" -------------

	err = runtime.BindQueryParameter("form", true, false, "redirectUrl", query, &params.RedirectUrl)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter redirectUrl: %w", err).Error(),
		)
	}

	return siw.Handler.Register(c, params)
}

// GetToken operation middleware
func (siw *ServerInterfaceWrapper) GetToken(c *fiber.Ctx) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTokenParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "cookie" -------------

	err = runtime.BindQueryParameter("form", true, false, "cookie", query, &params.Cookie)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter cookie: %w", err).Error(),
		)
	}

	// ------------- Required query parameter "responseType" -------------

	if paramValue := c.Query("responseType"); paramValue != "" {
	} else {
		err = fmt.Errorf("query argument responseType is required, but not found")
		return fiber.NewError(
			fiber.StatusBadRequest,
			err.Error(),
		)
	}

	err = runtime.BindQueryParameter(
		"form",
		true,
		true,
		"responseType",
		query,
		&params.ResponseType,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter responseType: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "redirectUrl" -------------

	err = runtime.BindQueryParameter("form", true, false, "redirectUrl", query, &params.RedirectUrl)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter redirectUrl: %w", err).Error(),
		)
	}

	return siw.Handler.GetToken(c, params)
}

// ListDomains operation middleware
func (siw *ServerInterfaceWrapper) ListDomains(c *fiber.Ctx) error {
	var err error

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListDomainsParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "roles" -------------

	err = runtime.BindQueryParameter("form", true, false, "roles", query, &params.Roles)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter roles: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "groups" -------------

	err = runtime.BindQueryParameter("form", true, false, "groups", query, &params.Groups)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter groups: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "ordering" -------------

	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter ordering: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter offset: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter limit: %w", err).Error(),
		)
	}

	return siw.Handler.ListDomains(c, params)
}

// CreateDomain operation middleware
func (siw *ServerInterfaceWrapper) CreateDomain(c *fiber.Ctx) error {
	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.CreateDomain(c)
}

// SearchDomainGroups operation middleware
func (siw *ServerInterfaceWrapper) SearchDomainGroups(c *fiber.Ctx) error {
	var err error

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params SearchDomainGroupsParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Required query parameter "query" -------------

	if paramValue := c.Query("query"); paramValue != "" {
	} else {
		err = fmt.Errorf("query argument query is required, but not found")
		return fiber.NewError(
			fiber.StatusBadRequest,
			err.Error(),
		)
	}

	err = runtime.BindQueryParameter("form", true, true, "query", query, &params.Query)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter query: %w", err).Error(),
		)
	}

	return siw.Handler.SearchDomainGroups(c, params)
}

// DeleteDomain operation middleware
func (siw *ServerInterfaceWrapper) DeleteDomain(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.DeleteDomain(c, domain)
}

// GetDomain operation middleware
func (siw *ServerInterfaceWrapper) GetDomain(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.GetDomain(c, domain)
}

// UpdateDomain operation middleware
func (siw *ServerInterfaceWrapper) UpdateDomain(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.UpdateDomain(c, domain)
}

// SearchDomainCandidates operation middleware
func (siw *ServerInterfaceWrapper) SearchDomainCandidates(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params SearchDomainCandidatesParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Required query parameter "query" -------------

	if paramValue := c.Query("query"); paramValue != "" {
	} else {
		err = fmt.Errorf("query argument query is required, but not found")
		return fiber.NewError(
			fiber.StatusBadRequest,
			err.Error(),
		)
	}

	err = runtime.BindQueryParameter("form", true, true, "query", query, &params.Query)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter query: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "ordering" -------------

	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter ordering: %w", err).Error(),
		)
	}

	return siw.Handler.SearchDomainCandidates(c, domain, params)
}

// ListDomainInvitations operation middleware
func (siw *ServerInterfaceWrapper) ListDomainInvitations(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListDomainInvitationsParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "ordering" -------------

	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter ordering: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter offset: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter limit: %w", err).Error(),
		)
	}

	return siw.Handler.ListDomainInvitations(c, domain, params)
}

// CreateDomainInvitation operation middleware
func (siw *ServerInterfaceWrapper) CreateDomainInvitation(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.CreateDomainInvitation(c, domain)
}

// DeleteDomainInvitation operation middleware
func (siw *ServerInterfaceWrapper) DeleteDomainInvitation(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "invitation" -------------
	var invitation string

	err = runtime.BindStyledParameter(
		"simple",
		false,
		"invitation",
		c.Params("invitation"),
		&invitation,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter invitation: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.DeleteDomainInvitation(c, domain, invitation)
}

// GetDomainInvitation operation middleware
func (siw *ServerInterfaceWrapper) GetDomainInvitation(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "invitation" -------------
	var invitation string

	err = runtime.BindStyledParameter(
		"simple",
		false,
		"invitation",
		c.Params("invitation"),
		&invitation,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter invitation: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.GetDomainInvitation(c, domain, invitation)
}

// UpdateDomainInvitation operation middleware
func (siw *ServerInterfaceWrapper) UpdateDomainInvitation(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "invitation" -------------
	var invitation string

	err = runtime.BindStyledParameter(
		"simple",
		false,
		"invitation",
		c.Params("invitation"),
		&invitation,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter invitation: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.UpdateDomainInvitation(c, domain, invitation)
}

// JoinDomainByInvitation operation middleware
func (siw *ServerInterfaceWrapper) JoinDomainByInvitation(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params JoinDomainByInvitationParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Required query parameter "invitationCode" -------------

	if paramValue := c.Query("invitationCode"); paramValue != "" {
	} else {
		err = fmt.Errorf("query argument invitationCode is required, but not found")
		return fiber.NewError(
			fiber.StatusBadRequest,
			err.Error(),
		)
	}

	err = runtime.BindQueryParameter(
		"form",
		true,
		true,
		"invitationCode",
		query,
		&params.InvitationCode,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter invitationCode: %w", err).Error(),
		)
	}

	return siw.Handler.JoinDomainByInvitation(c, domain, params)
}

// ListProblemSets operation middleware
func (siw *ServerInterfaceWrapper) ListProblemSets(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListProblemSetsParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "ordering" -------------

	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter ordering: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter offset: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter limit: %w", err).Error(),
		)
	}

	return siw.Handler.ListProblemSets(c, domain, params)
}

// CreateProblemSet operation middleware
func (siw *ServerInterfaceWrapper) CreateProblemSet(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.CreateProblemSet(c, domain)
}

// DeleteProblemSet operation middleware
func (siw *ServerInterfaceWrapper) DeleteProblemSet(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problemSet" -------------
	var problemSet string

	err = runtime.BindStyledParameter(
		"simple",
		false,
		"problemSet",
		c.Params("problemSet"),
		&problemSet,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problemSet: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.DeleteProblemSet(c, domain, problemSet)
}

// GetProblemSet operation middleware
func (siw *ServerInterfaceWrapper) GetProblemSet(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problemSet" -------------
	var problemSet string

	err = runtime.BindStyledParameter(
		"simple",
		false,
		"problemSet",
		c.Params("problemSet"),
		&problemSet,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problemSet: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.GetProblemSet(c, domain, problemSet)
}

// UpdateProblemSet operation middleware
func (siw *ServerInterfaceWrapper) UpdateProblemSet(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problemSet" -------------
	var problemSet string

	err = runtime.BindStyledParameter(
		"simple",
		false,
		"problemSet",
		c.Params("problemSet"),
		&problemSet,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problemSet: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.UpdateProblemSet(c, domain, problemSet)
}

// ListProblemsInProblemSet operation middleware
func (siw *ServerInterfaceWrapper) ListProblemsInProblemSet(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problemSet" -------------
	var problemSet string

	err = runtime.BindStyledParameter(
		"simple",
		false,
		"problemSet",
		c.Params("problemSet"),
		&problemSet,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problemSet: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.ListProblemsInProblemSet(c, domain, problemSet)
}

// AddProblemInProblemSet operation middleware
func (siw *ServerInterfaceWrapper) AddProblemInProblemSet(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problemSet" -------------
	var problemSet string

	err = runtime.BindStyledParameter(
		"simple",
		false,
		"problemSet",
		c.Params("problemSet"),
		&problemSet,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problemSet: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.AddProblemInProblemSet(c, domain, problemSet)
}

// DeleteProblemInProblemSet operation middleware
func (siw *ServerInterfaceWrapper) DeleteProblemInProblemSet(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problemSet" -------------
	var problemSet string

	err = runtime.BindStyledParameter(
		"simple",
		false,
		"problemSet",
		c.Params("problemSet"),
		&problemSet,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problemSet: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problem" -------------
	var problem string

	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problem: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.DeleteProblemInProblemSet(c, domain, problemSet, problem)
}

// GetProblemInProblemSet operation middleware
func (siw *ServerInterfaceWrapper) GetProblemInProblemSet(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problemSet" -------------
	var problemSet string

	err = runtime.BindStyledParameter(
		"simple",
		false,
		"problemSet",
		c.Params("problemSet"),
		&problemSet,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problemSet: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problem" -------------
	var problem string

	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problem: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.GetProblemInProblemSet(c, domain, problemSet, problem)
}

// UpdateProblemInProblemSet operation middleware
func (siw *ServerInterfaceWrapper) UpdateProblemInProblemSet(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problemSet" -------------
	var problemSet string

	err = runtime.BindStyledParameter(
		"simple",
		false,
		"problemSet",
		c.Params("problemSet"),
		&problemSet,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problemSet: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problem" -------------
	var problem string

	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problem: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.UpdateProblemInProblemSet(c, domain, problemSet, problem)
}

// SubmitSolutionToProblemSet operation middleware
func (siw *ServerInterfaceWrapper) SubmitSolutionToProblemSet(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problemSet" -------------
	var problemSet string

	err = runtime.BindStyledParameter(
		"simple",
		false,
		"problemSet",
		c.Params("problemSet"),
		&problemSet,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problemSet: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problem" -------------
	var problem string

	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problem: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.SubmitSolutionToProblemSet(c, domain, problemSet, problem)
}

// ListProblems operation middleware
func (siw *ServerInterfaceWrapper) ListProblems(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListProblemsParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "ordering" -------------

	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter ordering: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter offset: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter limit: %w", err).Error(),
		)
	}

	return siw.Handler.ListProblems(c, domain, params)
}

// CreateProblem operation middleware
func (siw *ServerInterfaceWrapper) CreateProblem(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.CreateProblem(c, domain)
}

// CloneProblem operation middleware
func (siw *ServerInterfaceWrapper) CloneProblem(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.CloneProblem(c, domain)
}

// DeleteProblem operation middleware
func (siw *ServerInterfaceWrapper) DeleteProblem(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problem" -------------
	var problem string

	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problem: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.DeleteProblem(c, domain, problem)
}

// GetProblem operation middleware
func (siw *ServerInterfaceWrapper) GetProblem(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problem" -------------
	var problem string

	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problem: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.GetProblem(c, domain, problem)
}

// UpdateProblem operation middleware
func (siw *ServerInterfaceWrapper) UpdateProblem(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problem" -------------
	var problem string

	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problem: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.UpdateProblem(c, domain, problem)
}

// SubmitSolutionToProblem operation middleware
func (siw *ServerInterfaceWrapper) SubmitSolutionToProblem(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problem" -------------
	var problem string

	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problem: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.SubmitSolutionToProblem(c, domain, problem)
}

// ListProblemConfigCommits operation middleware
func (siw *ServerInterfaceWrapper) ListProblemConfigCommits(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problem" -------------
	var problem string

	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problem: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListProblemConfigCommitsParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "ordering" -------------

	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter ordering: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter offset: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter limit: %w", err).Error(),
		)
	}

	return siw.Handler.ListProblemConfigCommits(c, domain, problem, params)
}

// UpdateProblemConfigByArchive operation middleware
func (siw *ServerInterfaceWrapper) UpdateProblemConfigByArchive(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problem" -------------
	var problem string

	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problem: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params UpdateProblemConfigByArchiveParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "configJsonOnMissing" -------------

	err = runtime.BindQueryParameter(
		"form",
		true,
		false,
		"configJsonOnMissing",
		query,
		&params.ConfigJsonOnMissing,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter configJsonOnMissing: %w", err).Error(),
		)
	}

	return siw.Handler.UpdateProblemConfigByArchive(c, domain, problem, params)
}

// UpdateProblemConfigJson operation middleware
func (siw *ServerInterfaceWrapper) UpdateProblemConfigJson(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problem" -------------
	var problem string

	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problem: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.UpdateProblemConfigJson(c, domain, problem)
}

// DiffProblemConfigDefaultBranch operation middleware
func (siw *ServerInterfaceWrapper) DiffProblemConfigDefaultBranch(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problem" -------------
	var problem string

	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problem: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params DiffProblemConfigDefaultBranchParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "after" -------------

	err = runtime.BindQueryParameter("form", true, false, "after", query, &params.After)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter after: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "amount" -------------

	err = runtime.BindQueryParameter("form", true, false, "amount", query, &params.Amount)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter amount: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "delimiter" -------------

	err = runtime.BindQueryParameter("form", true, false, "delimiter", query, &params.Delimiter)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter delimiter: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "prefix" -------------

	err = runtime.BindQueryParameter("form", true, false, "prefix", query, &params.Prefix)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter prefix: %w", err).Error(),
		)
	}

	return siw.Handler.DiffProblemConfigDefaultBranch(c, domain, problem, params)
}

// ListLatestProblemConfigObjectsUnderAGivenPrefix operation middleware
func (siw *ServerInterfaceWrapper) ListLatestProblemConfigObjectsUnderAGivenPrefix(
	c *fiber.Ctx,
) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problem" -------------
	var problem string

	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problem: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListLatestProblemConfigObjectsUnderAGivenPrefixParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "after" -------------

	err = runtime.BindQueryParameter("form", true, false, "after", query, &params.After)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter after: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "amount" -------------

	err = runtime.BindQueryParameter("form", true, false, "amount", query, &params.Amount)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter amount: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "delimiter" -------------

	err = runtime.BindQueryParameter("form", true, false, "delimiter", query, &params.Delimiter)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter delimiter: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "prefix" -------------

	err = runtime.BindQueryParameter("form", true, false, "prefix", query, &params.Prefix)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter prefix: %w", err).Error(),
		)
	}

	return siw.Handler.ListLatestProblemConfigObjectsUnderAGivenPrefix(c, domain, problem, params)
}

// DownloadProblemConfigArchive operation middleware
func (siw *ServerInterfaceWrapper) DownloadProblemConfigArchive(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problem" -------------
	var problem string

	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problem: %w", err).Error(),
		)
	}

	// ------------- Path parameter "config" -------------
	var config string

	err = runtime.BindStyledParameter("simple", false, "config", c.Params("config"), &config)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter config: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params DownloadProblemConfigArchiveParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "archiveFormat" -------------

	err = runtime.BindQueryParameter(
		"form",
		true,
		false,
		"archiveFormat",
		query,
		&params.ArchiveFormat,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter archiveFormat: %w", err).Error(),
		)
	}

	return siw.Handler.DownloadProblemConfigArchive(c, domain, problem, config, params)
}

// GetProblemConfigJson operation middleware
func (siw *ServerInterfaceWrapper) GetProblemConfigJson(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "problem" -------------
	var problem string

	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problem: %w", err).Error(),
		)
	}

	// ------------- Path parameter "config" -------------
	var config string

	err = runtime.BindStyledParameter("simple", false, "config", c.Params("config"), &config)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter config: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.GetProblemConfigJson(c, domain, problem, config)
}

// ListRecordsInDomain operation middleware
func (siw *ServerInterfaceWrapper) ListRecordsInDomain(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListRecordsInDomainParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "problemSet" -------------

	err = runtime.BindQueryParameter("form", true, false, "problemSet", query, &params.ProblemSet)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problemSet: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "problem" -------------

	err = runtime.BindQueryParameter("form", true, false, "problem", query, &params.Problem)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter problem: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "submitterId" -------------

	err = runtime.BindQueryParameter("form", true, false, "submitterId", query, &params.SubmitterId)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter submitterId: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "ordering" -------------

	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter ordering: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter offset: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter limit: %w", err).Error(),
		)
	}

	return siw.Handler.ListRecordsInDomain(c, domain, params)
}

// GetRecord operation middleware
func (siw *ServerInterfaceWrapper) GetRecord(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "record" -------------
	var record openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "record", c.Params("record"), &record)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter record: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.GetRecord(c, domain, record)
}

// SubmitCaseByJudger operation middleware
func (siw *ServerInterfaceWrapper) SubmitCaseByJudger(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "record" -------------
	var record string

	err = runtime.BindStyledParameter("simple", false, "record", c.Params("record"), &record)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter record: %w", err).Error(),
		)
	}

	// ------------- Path parameter "index" -------------
	var index int

	err = runtime.BindStyledParameter("simple", false, "index", c.Params("index"), &index)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter index: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.SubmitCaseByJudger(c, domain, record, index)
}

// SubmitRecordByJudger operation middleware
func (siw *ServerInterfaceWrapper) SubmitRecordByJudger(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "record" -------------
	var record string

	err = runtime.BindStyledParameter("simple", false, "record", c.Params("record"), &record)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter record: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.SubmitRecordByJudger(c, domain, record)
}

// ClaimRecordByJudger operation middleware
func (siw *ServerInterfaceWrapper) ClaimRecordByJudger(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "record" -------------
	var record string

	err = runtime.BindStyledParameter("simple", false, "record", c.Params("record"), &record)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter record: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.ClaimRecordByJudger(c, domain, record)
}

// ListDomainRoles operation middleware
func (siw *ServerInterfaceWrapper) ListDomainRoles(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListDomainRolesParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "ordering" -------------

	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter ordering: %w", err).Error(),
		)
	}

	return siw.Handler.ListDomainRoles(c, domain, params)
}

// CreateDomainRole operation middleware
func (siw *ServerInterfaceWrapper) CreateDomainRole(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.CreateDomainRole(c, domain)
}

// DeleteDomainRole operation middleware
func (siw *ServerInterfaceWrapper) DeleteDomainRole(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "role" -------------
	var role string

	err = runtime.BindStyledParameter("simple", false, "role", c.Params("role"), &role)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter role: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.DeleteDomainRole(c, domain, role)
}

// GetDomainRole operation middleware
func (siw *ServerInterfaceWrapper) GetDomainRole(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "role" -------------
	var role string

	err = runtime.BindStyledParameter("simple", false, "role", c.Params("role"), &role)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter role: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.GetDomainRole(c, domain, role)
}

// UpdateDomainRole operation middleware
func (siw *ServerInterfaceWrapper) UpdateDomainRole(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "role" -------------
	var role string

	err = runtime.BindStyledParameter("simple", false, "role", c.Params("role"), &role)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter role: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.UpdateDomainRole(c, domain, role)
}

// TransferDomain operation middleware
func (siw *ServerInterfaceWrapper) TransferDomain(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.TransferDomain(c, domain)
}

// ListDomainUsers operation middleware
func (siw *ServerInterfaceWrapper) ListDomainUsers(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListDomainUsersParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "ordering" -------------

	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter ordering: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter offset: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter limit: %w", err).Error(),
		)
	}

	return siw.Handler.ListDomainUsers(c, domain, params)
}

// AddDomainUser operation middleware
func (siw *ServerInterfaceWrapper) AddDomainUser(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.AddDomainUser(c, domain)
}

// RemoveDomainUser operation middleware
func (siw *ServerInterfaceWrapper) RemoveDomainUser(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "user" -------------
	var user string

	err = runtime.BindStyledParameter("simple", false, "user", c.Params("user"), &user)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter user: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.RemoveDomainUser(c, domain, user)
}

// GetDomainUser operation middleware
func (siw *ServerInterfaceWrapper) GetDomainUser(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "user" -------------
	var user string

	err = runtime.BindStyledParameter("simple", false, "user", c.Params("user"), &user)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter user: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.GetDomainUser(c, domain, user)
}

// UpdateDomainUser operation middleware
func (siw *ServerInterfaceWrapper) UpdateDomainUser(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "user" -------------
	var user string

	err = runtime.BindStyledParameter("simple", false, "user", c.Params("user"), &user)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter user: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.UpdateDomainUser(c, domain, user)
}

// GetDomainUserPermission operation middleware
func (siw *ServerInterfaceWrapper) GetDomainUserPermission(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "domain" -------------
	var domain string

	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter domain: %w", err).Error(),
		)
	}

	// ------------- Path parameter "user" -------------
	var user string

	err = runtime.BindStyledParameter("simple", false, "user", c.Params("user"), &user)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter user: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.GetDomainUserPermission(c, domain, user)
}

// JwtDecoded operation middleware
func (siw *ServerInterfaceWrapper) JwtDecoded(c *fiber.Ctx) error {
	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.JwtDecoded(c)
}

// ListProblemGroups operation middleware
func (siw *ServerInterfaceWrapper) ListProblemGroups(c *fiber.Ctx) error {
	var err error

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListProblemGroupsParams

	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for query string: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "ordering" -------------

	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter ordering: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter offset: %w", err).Error(),
		)
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter limit: %w", err).Error(),
		)
	}

	return siw.Handler.ListProblemGroups(c, params)
}

// TestErrorReport operation middleware
func (siw *ServerInterfaceWrapper) TestErrorReport(c *fiber.Ctx) error {
	return siw.Handler.TestErrorReport(c)
}

// GetCurrentUser operation middleware
func (siw *ServerInterfaceWrapper) GetCurrentUser(c *fiber.Ctx) error {
	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.GetCurrentUser(c)
}

// UpdateCurrentUser operation middleware
func (siw *ServerInterfaceWrapper) UpdateCurrentUser(c *fiber.Ctx) error {
	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.UpdateCurrentUser(c)
}

// ChangePassword operation middleware
func (siw *ServerInterfaceWrapper) ChangePassword(c *fiber.Ctx) error {
	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.ChangePassword(c)
}

// GetUser operation middleware
func (siw *ServerInterfaceWrapper) GetUser(c *fiber.Ctx) error {
	var err error

	// ------------- Path parameter "uid" -------------
	var uid string

	err = runtime.BindStyledParameter("simple", false, "uid", c.Params("uid"), &uid)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Errorf("invalid format for parameter uid: %w", err).Error(),
		)
	}

	c.Context().SetUserValue(HTTPBearerScopes, []string{""})

	return siw.Handler.GetUser(c, uid)
}

// Version operation middleware
//
//	@Summary	Version
//	@Tags		miscellaneous
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	schemas.VersionResp
//	@Router		/version [get]
func (siw *ServerInterfaceWrapper) Version(c *fiber.Ctx) error {
	return siw.Handler.Version(c)
}

type AdminListDomainRolesRequestObject struct {
	Params AdminListDomainRolesParams
}

type AdminListJudgersRequestObject struct {
	Params AdminListJudgersParams
}

type AdminCreateJudgerRequestObject struct {
	Body *AdminCreateJudgerJSONRequestBody
}

type AdminListUsersRequestObject struct {
	Params AdminListUsersParams
}

type AdminGetUserRequestObject struct {
	Uid string `json:"uid"`
}

type AdminListUserDomainsRequestObject struct {
	Uid    string `json:"uid"`
	Params AdminListUserDomainsParams
}

type LoginRequestObject struct {
	Params LoginParams
	Body   *LoginFormdataRequestBody
}

type LogoutRequestObject struct {
	Params LogoutParams
}

type ListOauth2RequestObject struct{}

type OauthAuthorizeRequestObject struct {
	Oauth2 string `json:"oauth2"`
	Params OauthAuthorizeParams
}

type RefreshRequestObject struct {
	Params RefreshParams
}

type RegisterRequestObject struct {
	Params RegisterParams
	Body   *RegisterJSONRequestBody
}

type GetTokenRequestObject struct {
	Params GetTokenParams
}

type ListDomainsRequestObject struct {
	Params ListDomainsParams
}

type CreateDomainRequestObject struct {
	Body *CreateDomainJSONRequestBody
}

type SearchDomainGroupsRequestObject struct {
	Params SearchDomainGroupsParams
}

type DeleteDomainRequestObject struct {
	Domain string `json:"domain"`
}

type GetDomainRequestObject struct {
	Domain string `json:"domain"`
}

type UpdateDomainRequestObject struct {
	Domain string `json:"domain"`
	Body   *UpdateDomainJSONRequestBody
}

type SearchDomainCandidatesRequestObject struct {
	Domain string `json:"domain"`
	Params SearchDomainCandidatesParams
}

type ListDomainInvitationsRequestObject struct {
	Domain string `json:"domain"`
	Params ListDomainInvitationsParams
}

type CreateDomainInvitationRequestObject struct {
	Domain string `json:"domain"`
	Body   *CreateDomainInvitationJSONRequestBody
}

type DeleteDomainInvitationRequestObject struct {
	Domain     string `json:"domain"`
	Invitation string `json:"invitation"`
}

type GetDomainInvitationRequestObject struct {
	Domain     string `json:"domain"`
	Invitation string `json:"invitation"`
}

type UpdateDomainInvitationRequestObject struct {
	Domain     string `json:"domain"`
	Invitation string `json:"invitation"`
	Body       *UpdateDomainInvitationJSONRequestBody
}

type JoinDomainByInvitationRequestObject struct {
	Domain string `json:"domain"`
	Params JoinDomainByInvitationParams
}

type ListProblemSetsRequestObject struct {
	Domain string `json:"domain"`
	Params ListProblemSetsParams
}

type CreateProblemSetRequestObject struct {
	Domain string `json:"domain"`
	Body   *CreateProblemSetJSONRequestBody
}

type DeleteProblemSetRequestObject struct {
	Domain     string `json:"domain"`
	ProblemSet string `json:"problemSet"`
}

type GetProblemSetRequestObject struct {
	Domain     string `json:"domain"`
	ProblemSet string `json:"problemSet"`
}

type UpdateProblemSetRequestObject struct {
	Domain     string `json:"domain"`
	ProblemSet string `json:"problemSet"`
	Body       *UpdateProblemSetJSONRequestBody
}

type ListProblemsInProblemSetRequestObject struct {
	Domain     string `json:"domain"`
	ProblemSet string `json:"problemSet"`
}

type AddProblemInProblemSetRequestObject struct {
	Domain     string `json:"domain"`
	ProblemSet string `json:"problemSet"`
	Body       *AddProblemInProblemSetJSONRequestBody
}

type DeleteProblemInProblemSetRequestObject struct {
	Domain     string `json:"domain"`
	ProblemSet string `json:"problemSet"`
	Problem    string `json:"problem"`
}

type GetProblemInProblemSetRequestObject struct {
	Domain     string `json:"domain"`
	ProblemSet string `json:"problemSet"`
	Problem    string `json:"problem"`
}

type UpdateProblemInProblemSetRequestObject struct {
	Domain     string `json:"domain"`
	ProblemSet string `json:"problemSet"`
	Problem    string `json:"problem"`
	Body       *UpdateProblemInProblemSetJSONRequestBody
}

type SubmitSolutionToProblemSetRequestObject struct {
	Domain     string `json:"domain"`
	ProblemSet string `json:"problemSet"`
	Problem    string `json:"problem"`
	Body       *multipart.Reader
}

type ListProblemsRequestObject struct {
	Domain string `json:"domain"`
	Params ListProblemsParams
}

type CreateProblemRequestObject struct {
	Domain string `json:"domain"`
	Body   *CreateProblemJSONRequestBody
}

type CloneProblemRequestObject struct {
	Domain string `json:"domain"`
	Body   *CloneProblemJSONRequestBody
}

type DeleteProblemRequestObject struct {
	Domain  string `json:"domain"`
	Problem string `json:"problem"`
}

type GetProblemRequestObject struct {
	Domain  string `json:"domain"`
	Problem string `json:"problem"`
}

type UpdateProblemRequestObject struct {
	Domain  string `json:"domain"`
	Problem string `json:"problem"`
	Body    *UpdateProblemJSONRequestBody
}

type SubmitSolutionToProblemRequestObject struct {
	Domain  string `json:"domain"`
	Problem string `json:"problem"`
	Body    *multipart.Reader
}

type ListProblemConfigCommitsRequestObject struct {
	Domain  string `json:"domain"`
	Problem string `json:"problem"`
	Params  ListProblemConfigCommitsParams
}

type UpdateProblemConfigByArchiveRequestObject struct {
	Domain  string `json:"domain"`
	Problem string `json:"problem"`
	Params  UpdateProblemConfigByArchiveParams
	Body    *multipart.Reader
}

type UpdateProblemConfigJsonRequestObject struct {
	Domain  string `json:"domain"`
	Problem string `json:"problem"`
	Body    *UpdateProblemConfigJsonJSONRequestBody
}

type DiffProblemConfigDefaultBranchRequestObject struct {
	Domain  string `json:"domain"`
	Problem string `json:"problem"`
	Params  DiffProblemConfigDefaultBranchParams
}

type ListLatestProblemConfigObjectsUnderAGivenPrefixRequestObject struct {
	Domain  string `json:"domain"`
	Problem string `json:"problem"`
	Params  ListLatestProblemConfigObjectsUnderAGivenPrefixParams
}

type DownloadProblemConfigArchiveRequestObject struct {
	Domain  string `json:"domain"`
	Problem string `json:"problem"`
	Config  string `json:"config"`
	Params  DownloadProblemConfigArchiveParams
}

type GetProblemConfigJsonRequestObject struct {
	Domain  string `json:"domain"`
	Problem string `json:"problem"`
	Config  string `json:"config"`
}

type ListRecordsInDomainRequestObject struct {
	Domain string `json:"domain"`
	Params ListRecordsInDomainParams
}

type GetRecordRequestObject struct {
	Domain string             `json:"domain"`
	Record openapi_types.UUID `json:"record"`
}

type SubmitCaseByJudgerRequestObject struct {
	Domain string `json:"domain"`
	Record string `json:"record"`
	Index  int    `json:"index"`
	Body   *SubmitCaseByJudgerJSONRequestBody
}

type SubmitRecordByJudgerRequestObject struct {
	Domain string `json:"domain"`
	Record string `json:"record"`
	Body   *SubmitRecordByJudgerJSONRequestBody
}

type ClaimRecordByJudgerRequestObject struct {
	Domain string `json:"domain"`
	Record string `json:"record"`
	Body   *ClaimRecordByJudgerJSONRequestBody
}

type ListDomainRolesRequestObject struct {
	Domain string `json:"domain"`
	Params ListDomainRolesParams
}

type CreateDomainRoleRequestObject struct {
	Domain string `json:"domain"`
	Body   *CreateDomainRoleJSONRequestBody
}

type DeleteDomainRoleRequestObject struct {
	Domain string `json:"domain"`
	Role   string `json:"role"`
}

type GetDomainRoleRequestObject struct {
	Domain string `json:"domain"`
	Role   string `json:"role"`
}

type UpdateDomainRoleRequestObject struct {
	Domain string `json:"domain"`
	Role   string `json:"role"`
	Body   *UpdateDomainRoleJSONRequestBody
}

type TransferDomainRequestObject struct {
	Domain string `json:"domain"`
	Body   *TransferDomainJSONRequestBody
}

type ListDomainUsersRequestObject struct {
	Domain string `json:"domain"`
	Params ListDomainUsersParams
}

type AddDomainUserRequestObject struct {
	Domain string `json:"domain"`
	Body   *AddDomainUserJSONRequestBody
}

type RemoveDomainUserRequestObject struct {
	Domain string `json:"domain"`
	User   string `json:"user"`
}

type GetDomainUserRequestObject struct {
	Domain string `json:"domain"`
	User   string `json:"user"`
}

type UpdateDomainUserRequestObject struct {
	Domain string `json:"domain"`
	User   string `json:"user"`
	Body   *UpdateDomainUserJSONRequestBody
}

type GetDomainUserPermissionRequestObject struct {
	Domain string `json:"domain"`
	User   string `json:"user"`
}

type JwtDecodedRequestObject struct{}

type ListProblemGroupsRequestObject struct {
	Params ListProblemGroupsParams
}

type TestErrorReportRequestObject struct{}

type GetCurrentUserRequestObject struct{}

type UpdateCurrentUserRequestObject struct {
	Body *UpdateCurrentUserJSONRequestBody
}

type ChangePasswordRequestObject struct {
	Body *ChangePasswordJSONRequestBody
}

type GetUserRequestObject struct {
	Uid string `json:"uid"`
}

type VersionRequestObject struct{}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Admin List Domain Roles
	// (GET /admin/domain_roles)
	AdminListDomainRoles(
		ctx *fiber.Ctx,
		request AdminListDomainRolesRequestObject,
	) (any, error)
	// Admin List Judgers
	// (GET /admin/judgers)
	AdminListJudgers(ctx *fiber.Ctx, request AdminListJudgersRequestObject) (any, error)
	// Admin Create Judger
	// (POST /admin/judgers)
	AdminCreateJudger(ctx *fiber.Ctx, request AdminCreateJudgerRequestObject) (any, error)
	// Admin List Users
	// (GET /admin/users)
	AdminListUsers(ctx *fiber.Ctx, request AdminListUsersRequestObject) (any, error)
	// Admin Get User
	// (GET /admin/{uid})
	AdminGetUser(ctx *fiber.Ctx, request AdminGetUserRequestObject) (any, error)
	// Admin List User Domains
	// (GET /admin/{uid}/domains)
	AdminListUserDomains(
		ctx *fiber.Ctx,
		request AdminListUserDomainsRequestObject,
	) (any, error)
	// Login
	// (POST /auth/login)
	Login(ctx *fiber.Ctx, request LoginRequestObject) (any, error)
	// Logout
	// (POST /auth/logout)
	Logout(ctx *fiber.Ctx, request LogoutRequestObject) (any, error)
	// List Oauth2
	// (GET /auth/oauth2)
	ListOauth2(ctx *fiber.Ctx, request ListOauth2RequestObject) (any, error)
	// Oauth Authorize
	// (GET /auth/oauth2/{oauth2}/authorize)
	OauthAuthorize(ctx *fiber.Ctx, request OauthAuthorizeRequestObject) (any, error)
	// Refresh
	// (POST /auth/refresh)
	Refresh(ctx *fiber.Ctx, request RefreshRequestObject) (any, error)
	// Register
	// (POST /auth/register)
	Register(ctx *fiber.Ctx, request RegisterRequestObject) (any, error)
	// Get Token
	// (GET /auth/token)
	GetToken(ctx *fiber.Ctx, request GetTokenRequestObject) (any, error)
	// List Domains
	// (GET /domains)
	ListDomains(ctx *fiber.Ctx, request ListDomainsRequestObject) (any, error)
	// Create Domain
	// (POST /domains)
	CreateDomain(ctx *fiber.Ctx, request CreateDomainRequestObject) (any, error)
	// Search Domain Groups
	// (GET /domains/groups)
	SearchDomainGroups(ctx *fiber.Ctx, request SearchDomainGroupsRequestObject) (any, error)
	// Delete Domain
	// (DELETE /domains/{domain})
	DeleteDomain(ctx *fiber.Ctx, request DeleteDomainRequestObject) (any, error)
	// Get Domain
	// (GET /domains/{domain})
	GetDomain(ctx *fiber.Ctx, request GetDomainRequestObject) (any, error)
	// Update Domain
	// (PATCH /domains/{domain})
	UpdateDomain(ctx *fiber.Ctx, request UpdateDomainRequestObject) (any, error)
	// Search Domain Candidates
	// (GET /domains/{domain}/candidates)
	SearchDomainCandidates(
		ctx *fiber.Ctx,
		request SearchDomainCandidatesRequestObject,
	) (any, error)
	// List Domain Invitations
	// (GET /domains/{domain}/invitations)
	ListDomainInvitations(
		ctx *fiber.Ctx,
		request ListDomainInvitationsRequestObject,
	) (any, error)
	// Create Domain Invitation
	// (POST /domains/{domain}/invitations)
	CreateDomainInvitation(
		ctx *fiber.Ctx,
		request CreateDomainInvitationRequestObject,
	) (any, error)
	// Delete Domain Invitation
	// (DELETE /domains/{domain}/invitations/{invitation})
	DeleteDomainInvitation(
		ctx *fiber.Ctx,
		request DeleteDomainInvitationRequestObject,
	) (any, error)
	// Get Domain Invitation
	// (GET /domains/{domain}/invitations/{invitation})
	GetDomainInvitation(ctx *fiber.Ctx, request GetDomainInvitationRequestObject) (any, error)
	// Update Domain Invitation
	// (PATCH /domains/{domain}/invitations/{invitation})
	UpdateDomainInvitation(
		ctx *fiber.Ctx,
		request UpdateDomainInvitationRequestObject,
	) (any, error)
	// Join Domain By Invitation
	// (POST /domains/{domain}/join)
	JoinDomainByInvitation(
		ctx *fiber.Ctx,
		request JoinDomainByInvitationRequestObject,
	) (any, error)
	// List Problem Sets
	// (GET /domains/{domain}/problem_sets)
	ListProblemSets(ctx *fiber.Ctx, request ListProblemSetsRequestObject) (any, error)
	// Create Problem Set
	// (POST /domains/{domain}/problem_sets)
	CreateProblemSet(ctx *fiber.Ctx, request CreateProblemSetRequestObject) (any, error)
	// Delete Problem Set
	// (DELETE /domains/{domain}/problem_sets/{problemSet})
	DeleteProblemSet(ctx *fiber.Ctx, request DeleteProblemSetRequestObject) (any, error)
	// Get Problem Set
	// (GET /domains/{domain}/problem_sets/{problemSet})
	GetProblemSet(ctx *fiber.Ctx, request GetProblemSetRequestObject) (any, error)
	// Update Problem Set
	// (PATCH /domains/{domain}/problem_sets/{problemSet})
	UpdateProblemSet(ctx *fiber.Ctx, request UpdateProblemSetRequestObject) (any, error)
	// List Problems In Problem Set
	// (GET /domains/{domain}/problem_sets/{problemSet}/problems)
	ListProblemsInProblemSet(
		ctx *fiber.Ctx,
		request ListProblemsInProblemSetRequestObject,
	) (any, error)
	// Add Problem In Problem Set
	// (POST /domains/{domain}/problem_sets/{problemSet}/problems)
	AddProblemInProblemSet(
		ctx *fiber.Ctx,
		request AddProblemInProblemSetRequestObject,
	) (any, error)
	// Delete Problem In Problem Set
	// (DELETE /domains/{domain}/problem_sets/{problemSet}/problems/{problem})
	DeleteProblemInProblemSet(
		ctx *fiber.Ctx,
		request DeleteProblemInProblemSetRequestObject,
	) (any, error)
	// Get Problem In Problem Set
	// (GET /domains/{domain}/problem_sets/{problemSet}/problems/{problem})
	GetProblemInProblemSet(
		ctx *fiber.Ctx,
		request GetProblemInProblemSetRequestObject,
	) (any, error)
	// Update Problem In Problem Set
	// (PATCH /domains/{domain}/problem_sets/{problemSet}/problems/{problem})
	UpdateProblemInProblemSet(
		ctx *fiber.Ctx,
		request UpdateProblemInProblemSetRequestObject,
	) (any, error)
	// Submit Solution To Problem Set
	// (POST /domains/{domain}/problem_sets/{problemSet}/problems/{problem}/submit)
	SubmitSolutionToProblemSet(
		ctx *fiber.Ctx,
		request SubmitSolutionToProblemSetRequestObject,
	) (any, error)
	// List Problems
	// (GET /domains/{domain}/problems)
	ListProblems(ctx *fiber.Ctx, request ListProblemsRequestObject) (any, error)
	// Create Problem
	// (POST /domains/{domain}/problems)
	CreateProblem(ctx *fiber.Ctx, request CreateProblemRequestObject) (any, error)
	// Clone Problem
	// (POST /domains/{domain}/problems/clone)
	CloneProblem(ctx *fiber.Ctx, request CloneProblemRequestObject) (any, error)
	// Delete Problem
	// (DELETE /domains/{domain}/problems/{problem})
	DeleteProblem(ctx *fiber.Ctx, request DeleteProblemRequestObject) (any, error)
	// Get Problem
	// (GET /domains/{domain}/problems/{problem})
	GetProblem(ctx *fiber.Ctx, request GetProblemRequestObject) (any, error)
	// Update Problem
	// (PATCH /domains/{domain}/problems/{problem})
	UpdateProblem(ctx *fiber.Ctx, request UpdateProblemRequestObject) (any, error)
	// Submit Solution To Problem
	// (POST /domains/{domain}/problems/{problem})
	SubmitSolutionToProblem(
		ctx *fiber.Ctx,
		request SubmitSolutionToProblemRequestObject,
	) (any, error)
	// List Problem Config Commits
	// (GET /domains/{domain}/problems/{problem}/configs)
	ListProblemConfigCommits(
		ctx *fiber.Ctx,
		request ListProblemConfigCommitsRequestObject,
	) (any, error)
	// Update Problem Config By Archive
	// (POST /domains/{domain}/problems/{problem}/configs)
	UpdateProblemConfigByArchive(
		ctx *fiber.Ctx,
		request UpdateProblemConfigByArchiveRequestObject,
	) (any, error)
	// Update Problem Config Json
	// (POST /domains/{domain}/problems/{problem}/configs/json)
	UpdateProblemConfigJson(
		ctx *fiber.Ctx,
		request UpdateProblemConfigJsonRequestObject,
	) (any, error)
	// Diff Problem Config Default Branch
	// (GET /domains/{domain}/problems/{problem}/configs/latest/diff)
	DiffProblemConfigDefaultBranch(
		ctx *fiber.Ctx,
		request DiffProblemConfigDefaultBranchRequestObject,
	) (any, error)
	// List Latest Problem Config Objects Under A Given Prefix
	// (GET /domains/{domain}/problems/{problem}/configs/latest/ls)
	ListLatestProblemConfigObjectsUnderAGivenPrefix(
		ctx *fiber.Ctx,
		request ListLatestProblemConfigObjectsUnderAGivenPrefixRequestObject,
	) (any, error)
	// Download Problem Config Archive
	// (GET /domains/{domain}/problems/{problem}/configs/{config})
	DownloadProblemConfigArchive(
		ctx *fiber.Ctx,
		request DownloadProblemConfigArchiveRequestObject,
	) (any, error)
	// Get Problem Config Json
	// (GET /domains/{domain}/problems/{problem}/configs/{config}/json)
	GetProblemConfigJson(
		ctx *fiber.Ctx,
		request GetProblemConfigJsonRequestObject,
	) (any, error)
	// List Records In Domain
	// (GET /domains/{domain}/records)
	ListRecordsInDomain(ctx *fiber.Ctx, request ListRecordsInDomainRequestObject) (any, error)
	// Get Record
	// (GET /domains/{domain}/records/{record})
	GetRecord(ctx *fiber.Ctx, request GetRecordRequestObject) (any, error)
	// Submit Case By Judger
	// (PUT /domains/{domain}/records/{record}/cases/{index}/judge)
	SubmitCaseByJudger(ctx *fiber.Ctx, request SubmitCaseByJudgerRequestObject) (any, error)
	// Submit Record By Judger
	// (PUT /domains/{domain}/records/{record}/judge)
	SubmitRecordByJudger(
		ctx *fiber.Ctx,
		request SubmitRecordByJudgerRequestObject,
	) (any, error)
	// Claim Record By Judger
	// (POST /domains/{domain}/records/{record}/judge/claim)
	ClaimRecordByJudger(ctx *fiber.Ctx, request ClaimRecordByJudgerRequestObject) (any, error)
	// List Domain Roles
	// (GET /domains/{domain}/roles)
	ListDomainRoles(ctx *fiber.Ctx, request ListDomainRolesRequestObject) (any, error)
	// Create Domain Role
	// (POST /domains/{domain}/roles)
	CreateDomainRole(ctx *fiber.Ctx, request CreateDomainRoleRequestObject) (any, error)
	// Delete Domain Role
	// (DELETE /domains/{domain}/roles/{role})
	DeleteDomainRole(ctx *fiber.Ctx, request DeleteDomainRoleRequestObject) (any, error)
	// Get Domain Role
	// (GET /domains/{domain}/roles/{role})
	GetDomainRole(ctx *fiber.Ctx, request GetDomainRoleRequestObject) (any, error)
	// Update Domain Role
	// (PATCH /domains/{domain}/roles/{role})
	UpdateDomainRole(ctx *fiber.Ctx, request UpdateDomainRoleRequestObject) (any, error)
	// Transfer Domain
	// (POST /domains/{domain}/transfer)
	TransferDomain(ctx *fiber.Ctx, request TransferDomainRequestObject) (any, error)
	// List Domain Users
	// (GET /domains/{domain}/users)
	ListDomainUsers(ctx *fiber.Ctx, request ListDomainUsersRequestObject) (any, error)
	// Add Domain User
	// (POST /domains/{domain}/users)
	AddDomainUser(ctx *fiber.Ctx, request AddDomainUserRequestObject) (any, error)
	// Remove Domain User
	// (DELETE /domains/{domain}/users/{user})
	RemoveDomainUser(ctx *fiber.Ctx, request RemoveDomainUserRequestObject) (any, error)
	// Get Domain User
	// (GET /domains/{domain}/users/{user})
	GetDomainUser(ctx *fiber.Ctx, request GetDomainUserRequestObject) (any, error)
	// Update Domain User
	// (PATCH /domains/{domain}/users/{user})
	UpdateDomainUser(ctx *fiber.Ctx, request UpdateDomainUserRequestObject) (any, error)
	// Get Domain User Permission
	// (GET /domains/{domain}/users/{user}/permission)
	GetDomainUserPermission(
		ctx *fiber.Ctx,
		request GetDomainUserPermissionRequestObject,
	) (any, error)
	// Jwt Decoded
	// (GET /jwt_decoded)
	JwtDecoded(ctx *fiber.Ctx, request JwtDecodedRequestObject) (any, error)
	// List Problem Groups
	// (GET /problem_groups)
	ListProblemGroups(ctx *fiber.Ctx, request ListProblemGroupsRequestObject) (any, error)
	// Test Error Report
	// (GET /test/report)
	TestErrorReport(ctx *fiber.Ctx, request TestErrorReportRequestObject) (any, error)
	// Get Current User
	// (GET /users/me)
	GetCurrentUser(ctx *fiber.Ctx, request GetCurrentUserRequestObject) (any, error)
	// Update Current User
	// (PATCH /users/me)
	UpdateCurrentUser(ctx *fiber.Ctx, request UpdateCurrentUserRequestObject) (any, error)
	// Change Password
	// (PATCH /users/me/password)
	ChangePassword(ctx *fiber.Ctx, request ChangePasswordRequestObject) (any, error)
	// Get User
	// (GET /users/{uid})
	GetUser(ctx *fiber.Ctx, request GetUserRequestObject) (any, error)
	// Version
	// (GET /version)
	Version(ctx *fiber.Ctx, request VersionRequestObject) (any, error)
}

type StrictHandlerFunc func(ctx *fiber.Ctx, args any) (any, error)

type StrictMiddlewareFunc func(f StrictHandlerFunc, operationID string) StrictHandlerFunc

type StrictRespHandlerFunc func(ctx *fiber.Ctx, response any, err error) error

func NewStrictHandler(
	ssi StrictServerInterface,
	middlewares []StrictMiddlewareFunc,
	responseHandler StrictRespHandlerFunc,
) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, responseHandler: responseHandler}
}

type strictHandler struct {
	ssi             StrictServerInterface
	middlewares     []StrictMiddlewareFunc
	responseHandler StrictRespHandlerFunc
}

// AdminListDomainRoles operation middleware
func (sh *strictHandler) AdminListDomainRoles(
	ctx *fiber.Ctx,
	params AdminListDomainRolesParams,
) error {
	var request AdminListDomainRolesRequestObject

	request.Params = params

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.AdminListDomainRoles(
			ctx,
			request.(AdminListDomainRolesRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AdminListDomainRoles")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// AdminListJudgers operation middleware
func (sh *strictHandler) AdminListJudgers(ctx *fiber.Ctx, params AdminListJudgersParams) error {
	var request AdminListJudgersRequestObject

	request.Params = params

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.AdminListJudgers(ctx, request.(AdminListJudgersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AdminListJudgers")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// AdminCreateJudger operation middleware
func (sh *strictHandler) AdminCreateJudger(ctx *fiber.Ctx) error {
	var request AdminCreateJudgerRequestObject

	var body AdminCreateJudgerJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.AdminCreateJudger(ctx, request.(AdminCreateJudgerRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AdminCreateJudger")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// AdminListUsers operation middleware
func (sh *strictHandler) AdminListUsers(ctx *fiber.Ctx, params AdminListUsersParams) error {
	var request AdminListUsersRequestObject

	request.Params = params

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.AdminListUsers(ctx, request.(AdminListUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AdminListUsers")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// AdminGetUser operation middleware
func (sh *strictHandler) AdminGetUser(ctx *fiber.Ctx, uid string) error {
	var request AdminGetUserRequestObject

	request.Uid = uid

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.AdminGetUser(ctx, request.(AdminGetUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AdminGetUser")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// AdminListUserDomains operation middleware
func (sh *strictHandler) AdminListUserDomains(
	ctx *fiber.Ctx,
	uid string,
	params AdminListUserDomainsParams,
) error {
	var request AdminListUserDomainsRequestObject

	request.Uid = uid
	request.Params = params

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.AdminListUserDomains(
			ctx,
			request.(AdminListUserDomainsRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AdminListUserDomains")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// Login operation middleware
func (sh *strictHandler) Login(ctx *fiber.Ctx, params LoginParams) error {
	var request LoginRequestObject

	request.Params = params

	var body LoginFormdataRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.Login(ctx, request.(LoginRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Login")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// Logout operation middleware
func (sh *strictHandler) Logout(ctx *fiber.Ctx, params LogoutParams) error {
	var request LogoutRequestObject

	request.Params = params

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.Logout(ctx, request.(LogoutRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Logout")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// ListOauth2 operation middleware
func (sh *strictHandler) ListOauth2(ctx *fiber.Ctx) error {
	var request ListOauth2RequestObject

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.ListOauth2(ctx, request.(ListOauth2RequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListOauth2")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// OauthAuthorize operation middleware
func (sh *strictHandler) OauthAuthorize(
	ctx *fiber.Ctx,
	oauth2 string,
	params OauthAuthorizeParams,
) error {
	var request OauthAuthorizeRequestObject

	request.Oauth2 = oauth2
	request.Params = params

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.OauthAuthorize(ctx, request.(OauthAuthorizeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "OauthAuthorize")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// Refresh operation middleware
func (sh *strictHandler) Refresh(ctx *fiber.Ctx, params RefreshParams) error {
	var request RefreshRequestObject

	request.Params = params

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.Refresh(ctx, request.(RefreshRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Refresh")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// Register operation middleware
func (sh *strictHandler) Register(ctx *fiber.Ctx, params RegisterParams) error {
	var request RegisterRequestObject

	request.Params = params

	var body RegisterJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.Register(ctx, request.(RegisterRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Register")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// GetToken operation middleware
func (sh *strictHandler) GetToken(ctx *fiber.Ctx, params GetTokenParams) error {
	var request GetTokenRequestObject

	request.Params = params

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.GetToken(ctx, request.(GetTokenRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetToken")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// ListDomains operation middleware
func (sh *strictHandler) ListDomains(ctx *fiber.Ctx, params ListDomainsParams) error {
	var request ListDomainsRequestObject

	request.Params = params

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.ListDomains(ctx, request.(ListDomainsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListDomains")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// CreateDomain operation middleware
func (sh *strictHandler) CreateDomain(ctx *fiber.Ctx) error {
	var request CreateDomainRequestObject

	var body CreateDomainJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.CreateDomain(ctx, request.(CreateDomainRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateDomain")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// SearchDomainGroups operation middleware
func (sh *strictHandler) SearchDomainGroups(ctx *fiber.Ctx, params SearchDomainGroupsParams) error {
	var request SearchDomainGroupsRequestObject

	request.Params = params

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.SearchDomainGroups(
			ctx,
			request.(SearchDomainGroupsRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "SearchDomainGroups")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// DeleteDomain operation middleware
func (sh *strictHandler) DeleteDomain(ctx *fiber.Ctx, domain string) error {
	var request DeleteDomainRequestObject

	request.Domain = domain

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.DeleteDomain(ctx, request.(DeleteDomainRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteDomain")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// GetDomain operation middleware
func (sh *strictHandler) GetDomain(ctx *fiber.Ctx, domain string) error {
	var request GetDomainRequestObject

	request.Domain = domain

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.GetDomain(ctx, request.(GetDomainRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetDomain")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// UpdateDomain operation middleware
func (sh *strictHandler) UpdateDomain(ctx *fiber.Ctx, domain string) error {
	var request UpdateDomainRequestObject

	request.Domain = domain

	var body UpdateDomainJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.UpdateDomain(ctx, request.(UpdateDomainRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateDomain")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// SearchDomainCandidates operation middleware
func (sh *strictHandler) SearchDomainCandidates(
	ctx *fiber.Ctx,
	domain string,
	params SearchDomainCandidatesParams,
) error {
	var request SearchDomainCandidatesRequestObject

	request.Domain = domain
	request.Params = params

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.SearchDomainCandidates(
			ctx,
			request.(SearchDomainCandidatesRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "SearchDomainCandidates")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// ListDomainInvitations operation middleware
func (sh *strictHandler) ListDomainInvitations(
	ctx *fiber.Ctx,
	domain string,
	params ListDomainInvitationsParams,
) error {
	var request ListDomainInvitationsRequestObject

	request.Domain = domain
	request.Params = params

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.ListDomainInvitations(
			ctx,
			request.(ListDomainInvitationsRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListDomainInvitations")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// CreateDomainInvitation operation middleware
func (sh *strictHandler) CreateDomainInvitation(ctx *fiber.Ctx, domain string) error {
	var request CreateDomainInvitationRequestObject

	request.Domain = domain

	var body CreateDomainInvitationJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.CreateDomainInvitation(
			ctx,
			request.(CreateDomainInvitationRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateDomainInvitation")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// DeleteDomainInvitation operation middleware
func (sh *strictHandler) DeleteDomainInvitation(
	ctx *fiber.Ctx,
	domain string,
	invitation string,
) error {
	var request DeleteDomainInvitationRequestObject

	request.Domain = domain
	request.Invitation = invitation

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.DeleteDomainInvitation(
			ctx,
			request.(DeleteDomainInvitationRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteDomainInvitation")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// GetDomainInvitation operation middleware
func (sh *strictHandler) GetDomainInvitation(
	ctx *fiber.Ctx,
	domain string,
	invitation string,
) error {
	var request GetDomainInvitationRequestObject

	request.Domain = domain
	request.Invitation = invitation

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.GetDomainInvitation(
			ctx,
			request.(GetDomainInvitationRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetDomainInvitation")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// UpdateDomainInvitation operation middleware
func (sh *strictHandler) UpdateDomainInvitation(
	ctx *fiber.Ctx,
	domain string,
	invitation string,
) error {
	var request UpdateDomainInvitationRequestObject

	request.Domain = domain
	request.Invitation = invitation

	var body UpdateDomainInvitationJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.UpdateDomainInvitation(
			ctx,
			request.(UpdateDomainInvitationRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateDomainInvitation")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// JoinDomainByInvitation operation middleware
func (sh *strictHandler) JoinDomainByInvitation(
	ctx *fiber.Ctx,
	domain string,
	params JoinDomainByInvitationParams,
) error {
	var request JoinDomainByInvitationRequestObject

	request.Domain = domain
	request.Params = params

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.JoinDomainByInvitation(
			ctx,
			request.(JoinDomainByInvitationRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "JoinDomainByInvitation")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// ListProblemSets operation middleware
func (sh *strictHandler) ListProblemSets(
	ctx *fiber.Ctx,
	domain string,
	params ListProblemSetsParams,
) error {
	var request ListProblemSetsRequestObject

	request.Domain = domain
	request.Params = params

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.ListProblemSets(ctx, request.(ListProblemSetsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListProblemSets")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// CreateProblemSet operation middleware
func (sh *strictHandler) CreateProblemSet(ctx *fiber.Ctx, domain string) error {
	var request CreateProblemSetRequestObject

	request.Domain = domain

	var body CreateProblemSetJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.CreateProblemSet(ctx, request.(CreateProblemSetRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateProblemSet")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// DeleteProblemSet operation middleware
func (sh *strictHandler) DeleteProblemSet(ctx *fiber.Ctx, domain string, problemSet string) error {
	var request DeleteProblemSetRequestObject

	request.Domain = domain
	request.ProblemSet = problemSet

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.DeleteProblemSet(ctx, request.(DeleteProblemSetRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteProblemSet")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// GetProblemSet operation middleware
func (sh *strictHandler) GetProblemSet(ctx *fiber.Ctx, domain string, problemSet string) error {
	var request GetProblemSetRequestObject

	request.Domain = domain
	request.ProblemSet = problemSet

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.GetProblemSet(ctx, request.(GetProblemSetRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetProblemSet")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// UpdateProblemSet operation middleware
func (sh *strictHandler) UpdateProblemSet(ctx *fiber.Ctx, domain string, problemSet string) error {
	var request UpdateProblemSetRequestObject

	request.Domain = domain
	request.ProblemSet = problemSet

	var body UpdateProblemSetJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.UpdateProblemSet(ctx, request.(UpdateProblemSetRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateProblemSet")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// ListProblemsInProblemSet operation middleware
func (sh *strictHandler) ListProblemsInProblemSet(
	ctx *fiber.Ctx,
	domain string,
	problemSet string,
) error {
	var request ListProblemsInProblemSetRequestObject

	request.Domain = domain
	request.ProblemSet = problemSet

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.ListProblemsInProblemSet(
			ctx,
			request.(ListProblemsInProblemSetRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListProblemsInProblemSet")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// AddProblemInProblemSet operation middleware
func (sh *strictHandler) AddProblemInProblemSet(
	ctx *fiber.Ctx,
	domain string,
	problemSet string,
) error {
	var request AddProblemInProblemSetRequestObject

	request.Domain = domain
	request.ProblemSet = problemSet

	var body AddProblemInProblemSetJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.AddProblemInProblemSet(
			ctx,
			request.(AddProblemInProblemSetRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AddProblemInProblemSet")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// DeleteProblemInProblemSet operation middleware
func (sh *strictHandler) DeleteProblemInProblemSet(
	ctx *fiber.Ctx,
	domain string,
	problemSet string,
	problem string,
) error {
	var request DeleteProblemInProblemSetRequestObject

	request.Domain = domain
	request.ProblemSet = problemSet
	request.Problem = problem

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.DeleteProblemInProblemSet(
			ctx,
			request.(DeleteProblemInProblemSetRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteProblemInProblemSet")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// GetProblemInProblemSet operation middleware
func (sh *strictHandler) GetProblemInProblemSet(
	ctx *fiber.Ctx,
	domain string,
	problemSet string,
	problem string,
) error {
	var request GetProblemInProblemSetRequestObject

	request.Domain = domain
	request.ProblemSet = problemSet
	request.Problem = problem

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.GetProblemInProblemSet(
			ctx,
			request.(GetProblemInProblemSetRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetProblemInProblemSet")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// UpdateProblemInProblemSet operation middleware
func (sh *strictHandler) UpdateProblemInProblemSet(
	ctx *fiber.Ctx,
	domain string,
	problemSet string,
	problem string,
) error {
	var request UpdateProblemInProblemSetRequestObject

	request.Domain = domain
	request.ProblemSet = problemSet
	request.Problem = problem

	var body UpdateProblemInProblemSetJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.UpdateProblemInProblemSet(
			ctx,
			request.(UpdateProblemInProblemSetRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateProblemInProblemSet")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// SubmitSolutionToProblemSet operation middleware
func (sh *strictHandler) SubmitSolutionToProblemSet(
	ctx *fiber.Ctx,
	domain string,
	problemSet string,
	problem string,
) error {
	var request SubmitSolutionToProblemSetRequestObject

	request.Domain = domain
	request.ProblemSet = problemSet
	request.Problem = problem

	request.Body = multipart.NewReader(
		bytes.NewReader(ctx.Request().Body()),
		string(ctx.Request().Header.MultipartFormBoundary()),
	)

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.SubmitSolutionToProblemSet(
			ctx,
			request.(SubmitSolutionToProblemSetRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "SubmitSolutionToProblemSet")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// ListProblems operation middleware
func (sh *strictHandler) ListProblems(
	ctx *fiber.Ctx,
	domain string,
	params ListProblemsParams,
) error {
	var request ListProblemsRequestObject

	request.Domain = domain
	request.Params = params

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.ListProblems(ctx, request.(ListProblemsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListProblems")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// CreateProblem operation middleware
func (sh *strictHandler) CreateProblem(ctx *fiber.Ctx, domain string) error {
	var request CreateProblemRequestObject

	request.Domain = domain

	var body CreateProblemJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.CreateProblem(ctx, request.(CreateProblemRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateProblem")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// CloneProblem operation middleware
func (sh *strictHandler) CloneProblem(ctx *fiber.Ctx, domain string) error {
	var request CloneProblemRequestObject

	request.Domain = domain

	var body CloneProblemJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.CloneProblem(ctx, request.(CloneProblemRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CloneProblem")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// DeleteProblem operation middleware
func (sh *strictHandler) DeleteProblem(ctx *fiber.Ctx, domain string, problem string) error {
	var request DeleteProblemRequestObject

	request.Domain = domain
	request.Problem = problem

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.DeleteProblem(ctx, request.(DeleteProblemRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteProblem")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// GetProblem operation middleware
func (sh *strictHandler) GetProblem(ctx *fiber.Ctx, domain string, problem string) error {
	var request GetProblemRequestObject

	request.Domain = domain
	request.Problem = problem

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.GetProblem(ctx, request.(GetProblemRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetProblem")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// UpdateProblem operation middleware
func (sh *strictHandler) UpdateProblem(ctx *fiber.Ctx, domain string, problem string) error {
	var request UpdateProblemRequestObject

	request.Domain = domain
	request.Problem = problem

	var body UpdateProblemJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.UpdateProblem(ctx, request.(UpdateProblemRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateProblem")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// SubmitSolutionToProblem operation middleware
func (sh *strictHandler) SubmitSolutionToProblem(
	ctx *fiber.Ctx,
	domain string,
	problem string,
) error {
	var request SubmitSolutionToProblemRequestObject

	request.Domain = domain
	request.Problem = problem

	request.Body = multipart.NewReader(
		bytes.NewReader(ctx.Request().Body()),
		string(ctx.Request().Header.MultipartFormBoundary()),
	)

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.SubmitSolutionToProblem(
			ctx,
			request.(SubmitSolutionToProblemRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "SubmitSolutionToProblem")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// ListProblemConfigCommits operation middleware
func (sh *strictHandler) ListProblemConfigCommits(
	ctx *fiber.Ctx,
	domain string,
	problem string,
	params ListProblemConfigCommitsParams,
) error {
	var request ListProblemConfigCommitsRequestObject

	request.Domain = domain
	request.Problem = problem
	request.Params = params

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.ListProblemConfigCommits(
			ctx,
			request.(ListProblemConfigCommitsRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListProblemConfigCommits")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// UpdateProblemConfigByArchive operation middleware
func (sh *strictHandler) UpdateProblemConfigByArchive(
	ctx *fiber.Ctx,
	domain string,
	problem string,
	params UpdateProblemConfigByArchiveParams,
) error {
	var request UpdateProblemConfigByArchiveRequestObject

	request.Domain = domain
	request.Problem = problem
	request.Params = params

	request.Body = multipart.NewReader(
		bytes.NewReader(ctx.Request().Body()),
		string(ctx.Request().Header.MultipartFormBoundary()),
	)

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.UpdateProblemConfigByArchive(
			ctx,
			request.(UpdateProblemConfigByArchiveRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateProblemConfigByArchive")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// UpdateProblemConfigJson operation middleware
func (sh *strictHandler) UpdateProblemConfigJson(
	ctx *fiber.Ctx,
	domain string,
	problem string,
) error {
	var request UpdateProblemConfigJsonRequestObject

	request.Domain = domain
	request.Problem = problem

	var body UpdateProblemConfigJsonJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.UpdateProblemConfigJson(
			ctx,
			request.(UpdateProblemConfigJsonRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateProblemConfigJson")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// DiffProblemConfigDefaultBranch operation middleware
func (sh *strictHandler) DiffProblemConfigDefaultBranch(
	ctx *fiber.Ctx,
	domain string,
	problem string,
	params DiffProblemConfigDefaultBranchParams,
) error {
	var request DiffProblemConfigDefaultBranchRequestObject

	request.Domain = domain
	request.Problem = problem
	request.Params = params

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.DiffProblemConfigDefaultBranch(
			ctx,
			request.(DiffProblemConfigDefaultBranchRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DiffProblemConfigDefaultBranch")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// ListLatestProblemConfigObjectsUnderAGivenPrefix operation middleware
func (sh *strictHandler) ListLatestProblemConfigObjectsUnderAGivenPrefix(
	ctx *fiber.Ctx,
	domain string,
	problem string,
	params ListLatestProblemConfigObjectsUnderAGivenPrefixParams,
) error {
	var request ListLatestProblemConfigObjectsUnderAGivenPrefixRequestObject

	request.Domain = domain
	request.Problem = problem
	request.Params = params

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.ListLatestProblemConfigObjectsUnderAGivenPrefix(
			ctx,
			request.(ListLatestProblemConfigObjectsUnderAGivenPrefixRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListLatestProblemConfigObjectsUnderAGivenPrefix")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// DownloadProblemConfigArchive operation middleware
func (sh *strictHandler) DownloadProblemConfigArchive(
	ctx *fiber.Ctx,
	domain string,
	problem string,
	config string,
	params DownloadProblemConfigArchiveParams,
) error {
	var request DownloadProblemConfigArchiveRequestObject

	request.Domain = domain
	request.Problem = problem
	request.Config = config
	request.Params = params

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.DownloadProblemConfigArchive(
			ctx,
			request.(DownloadProblemConfigArchiveRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DownloadProblemConfigArchive")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// GetProblemConfigJson operation middleware
func (sh *strictHandler) GetProblemConfigJson(
	ctx *fiber.Ctx,
	domain string,
	problem string,
	config string,
) error {
	var request GetProblemConfigJsonRequestObject

	request.Domain = domain
	request.Problem = problem
	request.Config = config

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.GetProblemConfigJson(
			ctx,
			request.(GetProblemConfigJsonRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetProblemConfigJson")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// ListRecordsInDomain operation middleware
func (sh *strictHandler) ListRecordsInDomain(
	ctx *fiber.Ctx,
	domain string,
	params ListRecordsInDomainParams,
) error {
	var request ListRecordsInDomainRequestObject

	request.Domain = domain
	request.Params = params

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.ListRecordsInDomain(
			ctx,
			request.(ListRecordsInDomainRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListRecordsInDomain")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// GetRecord operation middleware
func (sh *strictHandler) GetRecord(ctx *fiber.Ctx, domain string, record openapi_types.UUID) error {
	var request GetRecordRequestObject

	request.Domain = domain
	request.Record = record

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.GetRecord(ctx, request.(GetRecordRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetRecord")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// SubmitCaseByJudger operation middleware
func (sh *strictHandler) SubmitCaseByJudger(
	ctx *fiber.Ctx,
	domain string,
	record string,
	index int,
) error {
	var request SubmitCaseByJudgerRequestObject

	request.Domain = domain
	request.Record = record
	request.Index = index

	var body SubmitCaseByJudgerJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.SubmitCaseByJudger(
			ctx,
			request.(SubmitCaseByJudgerRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "SubmitCaseByJudger")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// SubmitRecordByJudger operation middleware
func (sh *strictHandler) SubmitRecordByJudger(ctx *fiber.Ctx, domain string, record string) error {
	var request SubmitRecordByJudgerRequestObject

	request.Domain = domain
	request.Record = record

	var body SubmitRecordByJudgerJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.SubmitRecordByJudger(
			ctx,
			request.(SubmitRecordByJudgerRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "SubmitRecordByJudger")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// ClaimRecordByJudger operation middleware
func (sh *strictHandler) ClaimRecordByJudger(ctx *fiber.Ctx, domain string, record string) error {
	var request ClaimRecordByJudgerRequestObject

	request.Domain = domain
	request.Record = record

	var body ClaimRecordByJudgerJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.ClaimRecordByJudger(
			ctx,
			request.(ClaimRecordByJudgerRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ClaimRecordByJudger")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// ListDomainRoles operation middleware
func (sh *strictHandler) ListDomainRoles(
	ctx *fiber.Ctx,
	domain string,
	params ListDomainRolesParams,
) error {
	var request ListDomainRolesRequestObject

	request.Domain = domain
	request.Params = params

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.ListDomainRoles(ctx, request.(ListDomainRolesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListDomainRoles")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// CreateDomainRole operation middleware
func (sh *strictHandler) CreateDomainRole(ctx *fiber.Ctx, domain string) error {
	var request CreateDomainRoleRequestObject

	request.Domain = domain

	var body CreateDomainRoleJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.CreateDomainRole(ctx, request.(CreateDomainRoleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateDomainRole")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// DeleteDomainRole operation middleware
func (sh *strictHandler) DeleteDomainRole(ctx *fiber.Ctx, domain string, role string) error {
	var request DeleteDomainRoleRequestObject

	request.Domain = domain
	request.Role = role

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.DeleteDomainRole(ctx, request.(DeleteDomainRoleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteDomainRole")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// GetDomainRole operation middleware
func (sh *strictHandler) GetDomainRole(ctx *fiber.Ctx, domain string, role string) error {
	var request GetDomainRoleRequestObject

	request.Domain = domain
	request.Role = role

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.GetDomainRole(ctx, request.(GetDomainRoleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetDomainRole")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// UpdateDomainRole operation middleware
func (sh *strictHandler) UpdateDomainRole(ctx *fiber.Ctx, domain string, role string) error {
	var request UpdateDomainRoleRequestObject

	request.Domain = domain
	request.Role = role

	var body UpdateDomainRoleJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.UpdateDomainRole(ctx, request.(UpdateDomainRoleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateDomainRole")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// TransferDomain operation middleware
func (sh *strictHandler) TransferDomain(ctx *fiber.Ctx, domain string) error {
	var request TransferDomainRequestObject

	request.Domain = domain

	var body TransferDomainJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.TransferDomain(ctx, request.(TransferDomainRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "TransferDomain")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// ListDomainUsers operation middleware
func (sh *strictHandler) ListDomainUsers(
	ctx *fiber.Ctx,
	domain string,
	params ListDomainUsersParams,
) error {
	var request ListDomainUsersRequestObject

	request.Domain = domain
	request.Params = params

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.ListDomainUsers(ctx, request.(ListDomainUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListDomainUsers")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// AddDomainUser operation middleware
func (sh *strictHandler) AddDomainUser(ctx *fiber.Ctx, domain string) error {
	var request AddDomainUserRequestObject

	request.Domain = domain

	var body AddDomainUserJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.AddDomainUser(ctx, request.(AddDomainUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AddDomainUser")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// RemoveDomainUser operation middleware
func (sh *strictHandler) RemoveDomainUser(ctx *fiber.Ctx, domain string, user string) error {
	var request RemoveDomainUserRequestObject

	request.Domain = domain
	request.User = user

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.RemoveDomainUser(ctx, request.(RemoveDomainUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "RemoveDomainUser")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// GetDomainUser operation middleware
func (sh *strictHandler) GetDomainUser(ctx *fiber.Ctx, domain string, user string) error {
	var request GetDomainUserRequestObject

	request.Domain = domain
	request.User = user

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.GetDomainUser(ctx, request.(GetDomainUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetDomainUser")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// UpdateDomainUser operation middleware
func (sh *strictHandler) UpdateDomainUser(ctx *fiber.Ctx, domain string, user string) error {
	var request UpdateDomainUserRequestObject

	request.Domain = domain
	request.User = user

	var body UpdateDomainUserJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.UpdateDomainUser(ctx, request.(UpdateDomainUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateDomainUser")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// GetDomainUserPermission operation middleware
func (sh *strictHandler) GetDomainUserPermission(ctx *fiber.Ctx, domain string, user string) error {
	var request GetDomainUserPermissionRequestObject

	request.Domain = domain
	request.User = user

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.GetDomainUserPermission(
			ctx,
			request.(GetDomainUserPermissionRequestObject),
		)
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetDomainUserPermission")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// JwtDecoded operation middleware
func (sh *strictHandler) JwtDecoded(ctx *fiber.Ctx) error {
	var request JwtDecodedRequestObject

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.JwtDecoded(ctx, request.(JwtDecodedRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "JwtDecoded")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// ListProblemGroups operation middleware
func (sh *strictHandler) ListProblemGroups(ctx *fiber.Ctx, params ListProblemGroupsParams) error {
	var request ListProblemGroupsRequestObject

	request.Params = params

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.ListProblemGroups(ctx, request.(ListProblemGroupsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListProblemGroups")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// TestErrorReport operation middleware
func (sh *strictHandler) TestErrorReport(ctx *fiber.Ctx) error {
	var request TestErrorReportRequestObject

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.TestErrorReport(ctx, request.(TestErrorReportRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "TestErrorReport")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// GetCurrentUser operation middleware
func (sh *strictHandler) GetCurrentUser(ctx *fiber.Ctx) error {
	var request GetCurrentUserRequestObject

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.GetCurrentUser(ctx, request.(GetCurrentUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetCurrentUser")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// UpdateCurrentUser operation middleware
func (sh *strictHandler) UpdateCurrentUser(ctx *fiber.Ctx) error {
	var request UpdateCurrentUserRequestObject

	var body UpdateCurrentUserJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.UpdateCurrentUser(ctx, request.(UpdateCurrentUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateCurrentUser")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// ChangePassword operation middleware
func (sh *strictHandler) ChangePassword(ctx *fiber.Ctx) error {
	var request ChangePasswordRequestObject

	var body ChangePasswordJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.ChangePassword(ctx, request.(ChangePasswordRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ChangePassword")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// GetUser operation middleware
func (sh *strictHandler) GetUser(ctx *fiber.Ctx, uid string) error {
	var request GetUserRequestObject

	request.Uid = uid

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.GetUser(ctx, request.(GetUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUser")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}

// Version operation middleware
func (sh *strictHandler) Version(ctx *fiber.Ctx) error {
	var request VersionRequestObject

	handler := func(ctx *fiber.Ctx, request any) (any, error) {
		return sh.ssi.Version(ctx, request.(VersionRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Version")
	}

	response, err := handler(ctx, request)
	return sh.responseHandler(ctx, response, err)
}
