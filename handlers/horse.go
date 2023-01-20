package handlers

// import (
// 	"fmt"
// 	"net/url"

// 	"github.com/deepmap/oapi-codegen/pkg/runtime"
// 	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/joint-online-judge/go-horse/model"
// )

// // ServerInterface represents all server handlers.
// type ServerInterface interface {
// 	// Admin List Domain Roles
// 	// (GET /admin/domain_roles)
// 	V1AdminListDomainRoles(c *fiber.Ctx, params model.V1AdminListDomainRolesParams) error
// 	// Admin List Judgers
// 	// (GET /admin/judgers)
// 	V1AdminListJudgers(c *fiber.Ctx, params model.V1AdminListJudgersParams) error
// 	// Admin Create Judger
// 	// (POST /admin/judgers)
// 	V1AdminCreateJudger(c *fiber.Ctx) error
// 	// Admin List Users
// 	// (GET /admin/users)
// 	V1AdminListUsers(c *fiber.Ctx, params model.V1AdminListUsersParams) error
// 	// Admin Get User
// 	// (GET /admin/{uid})
// 	V1AdminGetUser(c *fiber.Ctx, uid string) error
// 	// Admin List User Domains
// 	// (GET /admin/{uid}/domains)
// 	V1AdminListUserDomains(c *fiber.Ctx, uid string, params model.V1AdminListUserDomainsParams) error
// 	// Login
// 	// (POST /auth/login)
// 	V1Login(c *fiber.Ctx, params model.V1LoginParams) error
// 	// Logout
// 	// (POST /auth/logout)
// 	V1Logout(c *fiber.Ctx, params model.V1LogoutParams) error
// 	// List Oauth2
// 	// (GET /auth/oauth2)
// 	V1ListOauth2(c *fiber.Ctx) error
// 	// Oauth Authorize
// 	// (GET /auth/oauth2/{oauth2}/authorize)
// 	V1OauthAuthorize(c *fiber.Ctx, oauth2 string, params model.V1OauthAuthorizeParams) error
// 	// Refresh
// 	// (POST /auth/refresh)
// 	V1Refresh(c *fiber.Ctx, params model.V1RefreshParams) error
// 	// Register
// 	// (POST /auth/register)
// 	V1Register(c *fiber.Ctx, params model.V1RegisterParams) error
// 	// Get Token
// 	// (GET /auth/token)
// 	V1GetToken(c *fiber.Ctx, params model.V1GetTokenParams) error
// 	// List Domains
// 	// (GET /domains)
// 	V1ListDomains(c *fiber.Ctx, params model.V1ListDomainsParams) error
// 	// Create Domain
// 	// (POST /domains)
// 	V1CreateDomain(c *fiber.Ctx) error
// 	// Search Domain Groups
// 	// (GET /domains/groups)
// 	V1SearchDomainGroups(c *fiber.Ctx, params model.V1SearchDomainGroupsParams) error
// 	// Delete Domain
// 	// (DELETE /domains/{domain})
// 	V1DeleteDomain(c *fiber.Ctx, domain string) error
// 	// Get Domain
// 	// (GET /domains/{domain})
// 	V1GetDomain(c *fiber.Ctx, domain string) error
// 	// Update Domain
// 	// (PATCH /domains/{domain})
// 	V1UpdateDomain(c *fiber.Ctx, domain string) error
// 	// Search Domain Candidates
// 	// (GET /domains/{domain}/candidates)
// 	V1SearchDomainCandidates(c *fiber.Ctx, domain string, params model.V1SearchDomainCandidatesParams) error
// 	// List Domain Invitations
// 	// (GET /domains/{domain}/invitations)
// 	V1ListDomainInvitations(c *fiber.Ctx, domain string, params model.V1ListDomainInvitationsParams) error
// 	// Create Domain Invitation
// 	// (POST /domains/{domain}/invitations)
// 	V1CreateDomainInvitation(c *fiber.Ctx, domain string) error
// 	// Delete Domain Invitation
// 	// (DELETE /domains/{domain}/invitations/{invitation})
// 	V1DeleteDomainInvitation(c *fiber.Ctx, domain string, invitation string) error
// 	// Get Domain Invitation
// 	// (GET /domains/{domain}/invitations/{invitation})
// 	V1GetDomainInvitation(c *fiber.Ctx, domain string, invitation string) error
// 	// Update Domain Invitation
// 	// (PATCH /domains/{domain}/invitations/{invitation})
// 	V1UpdateDomainInvitation(c *fiber.Ctx, domain string, invitation string) error
// 	// Join Domain By Invitation
// 	// (POST /domains/{domain}/join)
// 	V1JoinDomainByInvitation(c *fiber.Ctx, domain string, params model.V1JoinDomainByInvitationParams) error
// 	// List Problem Sets
// 	// (GET /domains/{domain}/problem_sets)
// 	V1ListProblemSets(c *fiber.Ctx, domain string, params model.V1ListProblemSetsParams) error
// 	// Create Problem Set
// 	// (POST /domains/{domain}/problem_sets)
// 	V1CreateProblemSet(c *fiber.Ctx, domain string) error
// 	// Delete Problem Set
// 	// (DELETE /domains/{domain}/problem_sets/{problemSet})
// 	V1DeleteProblemSet(c *fiber.Ctx, domain string, problemSet string) error
// 	// Get Problem Set
// 	// (GET /domains/{domain}/problem_sets/{problemSet})
// 	V1GetProblemSet(c *fiber.Ctx, domain string, problemSet string) error
// 	// Update Problem Set
// 	// (PATCH /domains/{domain}/problem_sets/{problemSet})
// 	V1UpdateProblemSet(c *fiber.Ctx, domain string, problemSet string) error
// 	// List Problems In Problem Set
// 	// (GET /domains/{domain}/problem_sets/{problemSet}/problems)
// 	V1ListProblemsInProblemSet(c *fiber.Ctx, domain string, problemSet string) error
// 	// Add Problem In Problem Set
// 	// (POST /domains/{domain}/problem_sets/{problemSet}/problems)
// 	V1AddProblemInProblemSet(c *fiber.Ctx, domain string, problemSet string) error
// 	// Delete Problem In Problem Set
// 	// (DELETE /domains/{domain}/problem_sets/{problemSet}/problems/{problem})
// 	V1DeleteProblemInProblemSet(c *fiber.Ctx, domain string, problemSet string, problem string) error
// 	// Get Problem In Problem Set
// 	// (GET /domains/{domain}/problem_sets/{problemSet}/problems/{problem})
// 	V1GetProblemInProblemSet(c *fiber.Ctx, domain string, problemSet string, problem string) error
// 	// Update Problem In Problem Set
// 	// (PATCH /domains/{domain}/problem_sets/{problemSet}/problems/{problem})
// 	V1UpdateProblemInProblemSet(c *fiber.Ctx, domain string, problemSet string, problem string) error
// 	// Submit Solution To Problem Set
// 	// (POST /domains/{domain}/problem_sets/{problemSet}/problems/{problem}/submit)
// 	V1SubmitSolutionToProblemSet(c *fiber.Ctx, domain string, problemSet string, problem string) error
// 	// List Problems
// 	// (GET /domains/{domain}/problems)
// 	V1ListProblems(c *fiber.Ctx, domain string, params model.V1ListProblemsParams) error
// 	// Create Problem
// 	// (POST /domains/{domain}/problems)
// 	V1CreateProblem(c *fiber.Ctx, domain string) error
// 	// Clone Problem
// 	// (POST /domains/{domain}/problems/clone)
// 	V1CloneProblem(c *fiber.Ctx, domain string) error
// 	// Delete Problem
// 	// (DELETE /domains/{domain}/problems/{problem})
// 	V1DeleteProblem(c *fiber.Ctx, domain string, problem string) error
// 	// Get Problem
// 	// (GET /domains/{domain}/problems/{problem})
// 	V1GetProblem(c *fiber.Ctx, domain string, problem string) error
// 	// Update Problem
// 	// (PATCH /domains/{domain}/problems/{problem})
// 	V1UpdateProblem(c *fiber.Ctx, domain string, problem string) error
// 	// Submit Solution To Problem
// 	// (POST /domains/{domain}/problems/{problem})
// 	V1SubmitSolutionToProblem(c *fiber.Ctx, domain string, problem string) error
// 	// List Problem Config Commits
// 	// (GET /domains/{domain}/problems/{problem}/configs)
// 	V1ListProblemConfigCommits(c *fiber.Ctx, domain string, problem string, params model.V1ListProblemConfigCommitsParams) error
// 	// Update Problem Config By Archive
// 	// (POST /domains/{domain}/problems/{problem}/configs)
// 	V1UpdateProblemConfigByArchive(c *fiber.Ctx, domain string, problem string, params model.V1UpdateProblemConfigByArchiveParams) error
// 	// Update Problem Config Json
// 	// (POST /domains/{domain}/problems/{problem}/configs/json)
// 	V1UpdateProblemConfigJson(c *fiber.Ctx, domain string, problem string) error
// 	// Diff Problem Config Default Branch
// 	// (GET /domains/{domain}/problems/{problem}/configs/latest/diff)
// 	V1DiffProblemConfigDefaultBranch(c *fiber.Ctx, domain string, problem string, params model.V1DiffProblemConfigDefaultBranchParams) error
// 	// List Latest Problem Config Objects Under A Given Prefix
// 	// (GET /domains/{domain}/problems/{problem}/configs/latest/ls)
// 	V1ListLatestProblemConfigObjectsUnderAGivenPrefix(c *fiber.Ctx, domain string, problem string, params model.V1ListLatestProblemConfigObjectsUnderAGivenPrefixParams) error
// 	// Download Problem Config Archive
// 	// (GET /domains/{domain}/problems/{problem}/configs/{config})
// 	V1DownloadProblemConfigArchive(c *fiber.Ctx, domain string, problem string, config string, params model.V1DownloadProblemConfigArchiveParams) error
// 	// Get Problem Config Json
// 	// (GET /domains/{domain}/problems/{problem}/configs/{config}/json)
// 	V1GetProblemConfigJson(c *fiber.Ctx, domain string, problem string, config string) error
// 	// List Records In Domain
// 	// (GET /domains/{domain}/records)
// 	V1ListRecordsInDomain(c *fiber.Ctx, domain string, params model.V1ListRecordsInDomainParams) error
// 	// Get Record
// 	// (GET /domains/{domain}/records/{record})
// 	V1GetRecord(c *fiber.Ctx, domain string, record openapi_types.UUID) error
// 	// Submit Case By Judger
// 	// (PUT /domains/{domain}/records/{record}/cases/{index}/judge)
// 	V1SubmitCaseByJudger(c *fiber.Ctx, domain string, record string, index int) error
// 	// Submit Record By Judger
// 	// (PUT /domains/{domain}/records/{record}/judge)
// 	V1SubmitRecordByJudger(c *fiber.Ctx, domain string, record string) error
// 	// Claim Record By Judger
// 	// (POST /domains/{domain}/records/{record}/judge/claim)
// 	V1ClaimRecordByJudger(c *fiber.Ctx, domain string, record string) error
// 	// List Domain Roles
// 	// (GET /domains/{domain}/roles)
// 	V1ListDomainRoles(c *fiber.Ctx, domain string, params model.V1ListDomainRolesParams) error
// 	// Create Domain Role
// 	// (POST /domains/{domain}/roles)
// 	V1CreateDomainRole(c *fiber.Ctx, domain string) error
// 	// Delete Domain Role
// 	// (DELETE /domains/{domain}/roles/{role})
// 	V1DeleteDomainRole(c *fiber.Ctx, domain string, role string) error
// 	// Get Domain Role
// 	// (GET /domains/{domain}/roles/{role})
// 	V1GetDomainRole(c *fiber.Ctx, domain string, role string) error
// 	// Update Domain Role
// 	// (PATCH /domains/{domain}/roles/{role})
// 	V1UpdateDomainRole(c *fiber.Ctx, domain string, role string) error
// 	// Transfer Domain
// 	// (POST /domains/{domain}/transfer)
// 	V1TransferDomain(c *fiber.Ctx, domain string) error
// 	// List Domain Users
// 	// (GET /domains/{domain}/users)
// 	V1ListDomainUsers(c *fiber.Ctx, domain string, params model.V1ListDomainUsersParams) error
// 	// Add Domain User
// 	// (POST /domains/{domain}/users)
// 	V1AddDomainUser(c *fiber.Ctx, domain string) error
// 	// Remove Domain User
// 	// (DELETE /domains/{domain}/users/{user})
// 	V1RemoveDomainUser(c *fiber.Ctx, domain string, user string) error
// 	// Get Domain User
// 	// (GET /domains/{domain}/users/{user})
// 	V1GetDomainUser(c *fiber.Ctx, domain string, user string) error
// 	// Update Domain User
// 	// (PATCH /domains/{domain}/users/{user})
// 	V1UpdateDomainUser(c *fiber.Ctx, domain string, user string) error
// 	// Get Domain User Permission
// 	// (GET /domains/{domain}/users/{user}/permission)
// 	V1GetDomainUserPermission(c *fiber.Ctx, domain string, user string) error
// 	// Jwt Decoded
// 	// (GET /jwt_decoded)
// 	V1JwtDecoded(c *fiber.Ctx) error
// 	// List Problem Groups
// 	// (GET /problem_groups)
// 	V1ListProblemGroups(c *fiber.Ctx, params model.V1ListProblemGroupsParams) error
// 	// Test Error Report
// 	// (GET /test/report)
// 	V1TestErrorReport(c *fiber.Ctx) error
// 	// Get Current User
// 	// (GET /users/me)
// 	V1GetCurrentUser(c *fiber.Ctx) error
// 	// Update Current User
// 	// (PATCH /users/me)
// 	V1UpdateCurrentUser(c *fiber.Ctx) error
// 	// Change Password
// 	// (PATCH /users/me/password)
// 	V1ChangePassword(c *fiber.Ctx) error
// 	// Get User
// 	// (GET /users/{uid})
// 	V1GetUser(c *fiber.Ctx, uid string) error
// 	// Version
// 	// (GET /version)
// 	V1Version(c *fiber.Ctx) error
// }

// // ServerInterfaceWrapper converts contexts to parameters.
// type ServerInterfaceWrapper struct {
// 	Handler ServerInterface
// }

// type MiddlewareFunc fiber.Handler

// // V1AdminListDomainRoles operation middleware
// func (siw *ServerInterfaceWrapper) V1AdminListDomainRoles(c *fiber.Ctx) error {

// 	var err error

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1AdminListDomainRolesParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "ordering" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter ordering: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "offset" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter offset: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "limit" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter limit: %w", err).Error())
// 	}

// 	return siw.Handler.V1AdminListDomainRoles(c, params)
// }

// // V1AdminListJudgers operation middleware
// func (siw *ServerInterfaceWrapper) V1AdminListJudgers(c *fiber.Ctx) error {

// 	var err error

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1AdminListJudgersParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "ordering" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter ordering: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "offset" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter offset: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "limit" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter limit: %w", err).Error())
// 	}

// 	return siw.Handler.V1AdminListJudgers(c, params)
// }

// // V1AdminCreateJudger operation middleware
// func (siw *ServerInterfaceWrapper) V1AdminCreateJudger(c *fiber.Ctx) error {

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1AdminCreateJudger(c)
// }

// // V1AdminListUsers operation middleware
// func (siw *ServerInterfaceWrapper) V1AdminListUsers(c *fiber.Ctx) error {

// 	var err error

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1AdminListUsersParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "ordering" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter ordering: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "offset" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter offset: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "limit" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter limit: %w", err).Error())
// 	}

// 	return siw.Handler.V1AdminListUsers(c, params)
// }

// // V1AdminGetUser operation middleware
// func (siw *ServerInterfaceWrapper) V1AdminGetUser(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "uid" -------------
// 	var uid string

// 	err = runtime.BindStyledParameter("simple", false, "uid", c.Params("uid"), &uid)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter uid: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1AdminGetUser(c, uid)
// }

// // V1AdminListUserDomains operation middleware
// func (siw *ServerInterfaceWrapper) V1AdminListUserDomains(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "uid" -------------
// 	var uid string

// 	err = runtime.BindStyledParameter("simple", false, "uid", c.Params("uid"), &uid)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter uid: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1AdminListUserDomainsParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "role" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "role", query, &params.Role)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter role: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "groups" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "groups", query, &params.Groups)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter groups: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "ordering" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter ordering: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "offset" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter offset: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "limit" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter limit: %w", err).Error())
// 	}

// 	return siw.Handler.V1AdminListUserDomains(c, uid, params)
// }

// // V1Login operation middleware
// func (siw *ServerInterfaceWrapper) V1Login(c *fiber.Ctx) error {

// 	var err error

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1LoginParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "cookie" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "cookie", query, &params.Cookie)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter cookie: %w", err).Error())
// 	}

// 	// ------------- Required query parameter "responseType" -------------

// 	if paramValue := c.Query("responseType"); paramValue != "" {

// 	} else {
// 		err = fmt.Errorf("Query argument responseType is required, but not found")
// 		c.Status(fiber.StatusBadRequest).JSON(err)
// 		return err
// 	}

// 	err = runtime.BindQueryParameter("form", true, true, "responseType", query, &params.ResponseType)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter responseType: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "redirectUrl" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "redirectUrl", query, &params.RedirectUrl)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter redirectUrl: %w", err).Error())
// 	}

// 	return siw.Handler.V1Login(c, params)
// }

// // V1Logout operation middleware
// func (siw *ServerInterfaceWrapper) V1Logout(c *fiber.Ctx) error {

// 	var err error

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1LogoutParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "cookie" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "cookie", query, &params.Cookie)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter cookie: %w", err).Error())
// 	}

// 	// ------------- Required query parameter "responseType" -------------

// 	if paramValue := c.Query("responseType"); paramValue != "" {

// 	} else {
// 		err = fmt.Errorf("Query argument responseType is required, but not found")
// 		c.Status(fiber.StatusBadRequest).JSON(err)
// 		return err
// 	}

// 	err = runtime.BindQueryParameter("form", true, true, "responseType", query, &params.ResponseType)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter responseType: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "redirectUrl" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "redirectUrl", query, &params.RedirectUrl)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter redirectUrl: %w", err).Error())
// 	}

// 	return siw.Handler.V1Logout(c, params)
// }

// // V1ListOauth2 operation middleware
// func (siw *ServerInterfaceWrapper) V1ListOauth2(c *fiber.Ctx) error {

// 	return siw.Handler.V1ListOauth2(c)
// }

// // V1OauthAuthorize operation middleware
// func (siw *ServerInterfaceWrapper) V1OauthAuthorize(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "oauth2" -------------
// 	var oauth2 string

// 	err = runtime.BindStyledParameter("simple", false, "oauth2", c.Params("oauth2"), &oauth2)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter oauth2: %w", err).Error())
// 	}

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1OauthAuthorizeParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "scopes" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "scopes", query, &params.Scopes)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter scopes: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "cookie" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "cookie", query, &params.Cookie)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter cookie: %w", err).Error())
// 	}

// 	// ------------- Required query parameter "responseType" -------------

// 	if paramValue := c.Query("responseType"); paramValue != "" {

// 	} else {
// 		err = fmt.Errorf("Query argument responseType is required, but not found")
// 		c.Status(fiber.StatusBadRequest).JSON(err)
// 		return err
// 	}

// 	err = runtime.BindQueryParameter("form", true, true, "responseType", query, &params.ResponseType)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter responseType: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "redirectUrl" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "redirectUrl", query, &params.RedirectUrl)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter redirectUrl: %w", err).Error())
// 	}

// 	return siw.Handler.V1OauthAuthorize(c, oauth2, params)
// }

// // V1Refresh operation middleware
// func (siw *ServerInterfaceWrapper) V1Refresh(c *fiber.Ctx) error {

// 	var err error

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1RefreshParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "cookie" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "cookie", query, &params.Cookie)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter cookie: %w", err).Error())
// 	}

// 	// ------------- Required query parameter "responseType" -------------

// 	if paramValue := c.Query("responseType"); paramValue != "" {

// 	} else {
// 		err = fmt.Errorf("Query argument responseType is required, but not found")
// 		c.Status(fiber.StatusBadRequest).JSON(err)
// 		return err
// 	}

// 	err = runtime.BindQueryParameter("form", true, true, "responseType", query, &params.ResponseType)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter responseType: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "redirectUrl" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "redirectUrl", query, &params.RedirectUrl)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter redirectUrl: %w", err).Error())
// 	}

// 	return siw.Handler.V1Refresh(c, params)
// }

// // V1Register operation middleware
// func (siw *ServerInterfaceWrapper) V1Register(c *fiber.Ctx) error {

// 	var err error

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1RegisterParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "cookie" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "cookie", query, &params.Cookie)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter cookie: %w", err).Error())
// 	}

// 	// ------------- Required query parameter "responseType" -------------

// 	if paramValue := c.Query("responseType"); paramValue != "" {

// 	} else {
// 		err = fmt.Errorf("Query argument responseType is required, but not found")
// 		c.Status(fiber.StatusBadRequest).JSON(err)
// 		return err
// 	}

// 	err = runtime.BindQueryParameter("form", true, true, "responseType", query, &params.ResponseType)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter responseType: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "redirectUrl" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "redirectUrl", query, &params.RedirectUrl)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter redirectUrl: %w", err).Error())
// 	}

// 	return siw.Handler.V1Register(c, params)
// }

// // V1GetToken operation middleware
// func (siw *ServerInterfaceWrapper) V1GetToken(c *fiber.Ctx) error {

// 	var err error

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1GetTokenParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "cookie" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "cookie", query, &params.Cookie)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter cookie: %w", err).Error())
// 	}

// 	// ------------- Required query parameter "responseType" -------------

// 	if paramValue := c.Query("responseType"); paramValue != "" {

// 	} else {
// 		err = fmt.Errorf("Query argument responseType is required, but not found")
// 		c.Status(fiber.StatusBadRequest).JSON(err)
// 		return err
// 	}

// 	err = runtime.BindQueryParameter("form", true, true, "responseType", query, &params.ResponseType)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter responseType: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "redirectUrl" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "redirectUrl", query, &params.RedirectUrl)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter redirectUrl: %w", err).Error())
// 	}

// 	return siw.Handler.V1GetToken(c, params)
// }

// // V1ListDomains operation middleware
// func (siw *ServerInterfaceWrapper) V1ListDomains(c *fiber.Ctx) error {

// 	var err error

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1ListDomainsParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "roles" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "roles", query, &params.Roles)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter roles: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "groups" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "groups", query, &params.Groups)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter groups: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "ordering" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter ordering: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "offset" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter offset: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "limit" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter limit: %w", err).Error())
// 	}

// 	return siw.Handler.V1ListDomains(c, params)
// }

// // V1CreateDomain operation middleware
// func (siw *ServerInterfaceWrapper) V1CreateDomain(c *fiber.Ctx) error {

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1CreateDomain(c)
// }

// // V1SearchDomainGroups operation middleware
// func (siw *ServerInterfaceWrapper) V1SearchDomainGroups(c *fiber.Ctx) error {

// 	var err error

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1SearchDomainGroupsParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Required query parameter "query" -------------

// 	if paramValue := c.Query("query"); paramValue != "" {

// 	} else {
// 		err = fmt.Errorf("Query argument query is required, but not found")
// 		c.Status(fiber.StatusBadRequest).JSON(err)
// 		return err
// 	}

// 	err = runtime.BindQueryParameter("form", true, true, "query", query, &params.Query)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter query: %w", err).Error())
// 	}

// 	return siw.Handler.V1SearchDomainGroups(c, params)
// }

// // V1DeleteDomain operation middleware
// func (siw *ServerInterfaceWrapper) V1DeleteDomain(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1DeleteDomain(c, domain)
// }

// // V1GetDomain operation middleware
// func (siw *ServerInterfaceWrapper) V1GetDomain(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1GetDomain(c, domain)
// }

// // V1UpdateDomain operation middleware
// func (siw *ServerInterfaceWrapper) V1UpdateDomain(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1UpdateDomain(c, domain)
// }

// // V1SearchDomainCandidates operation middleware
// func (siw *ServerInterfaceWrapper) V1SearchDomainCandidates(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1SearchDomainCandidatesParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Required query parameter "query" -------------

// 	if paramValue := c.Query("query"); paramValue != "" {

// 	} else {
// 		err = fmt.Errorf("Query argument query is required, but not found")
// 		c.Status(fiber.StatusBadRequest).JSON(err)
// 		return err
// 	}

// 	err = runtime.BindQueryParameter("form", true, true, "query", query, &params.Query)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter query: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "ordering" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter ordering: %w", err).Error())
// 	}

// 	return siw.Handler.V1SearchDomainCandidates(c, domain, params)
// }

// // V1ListDomainInvitations operation middleware
// func (siw *ServerInterfaceWrapper) V1ListDomainInvitations(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1ListDomainInvitationsParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "ordering" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter ordering: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "offset" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter offset: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "limit" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter limit: %w", err).Error())
// 	}

// 	return siw.Handler.V1ListDomainInvitations(c, domain, params)
// }

// // V1CreateDomainInvitation operation middleware
// func (siw *ServerInterfaceWrapper) V1CreateDomainInvitation(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1CreateDomainInvitation(c, domain)
// }

// // V1DeleteDomainInvitation operation middleware
// func (siw *ServerInterfaceWrapper) V1DeleteDomainInvitation(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "invitation" -------------
// 	var invitation string

// 	err = runtime.BindStyledParameter("simple", false, "invitation", c.Params("invitation"), &invitation)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter invitation: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1DeleteDomainInvitation(c, domain, invitation)
// }

// // V1GetDomainInvitation operation middleware
// func (siw *ServerInterfaceWrapper) V1GetDomainInvitation(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "invitation" -------------
// 	var invitation string

// 	err = runtime.BindStyledParameter("simple", false, "invitation", c.Params("invitation"), &invitation)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter invitation: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1GetDomainInvitation(c, domain, invitation)
// }

// // V1UpdateDomainInvitation operation middleware
// func (siw *ServerInterfaceWrapper) V1UpdateDomainInvitation(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "invitation" -------------
// 	var invitation string

// 	err = runtime.BindStyledParameter("simple", false, "invitation", c.Params("invitation"), &invitation)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter invitation: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1UpdateDomainInvitation(c, domain, invitation)
// }

// // V1JoinDomainByInvitation operation middleware
// func (siw *ServerInterfaceWrapper) V1JoinDomainByInvitation(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1JoinDomainByInvitationParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Required query parameter "invitationCode" -------------

// 	if paramValue := c.Query("invitationCode"); paramValue != "" {

// 	} else {
// 		err = fmt.Errorf("Query argument invitationCode is required, but not found")
// 		c.Status(fiber.StatusBadRequest).JSON(err)
// 		return err
// 	}

// 	err = runtime.BindQueryParameter("form", true, true, "invitationCode", query, &params.InvitationCode)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter invitationCode: %w", err).Error())
// 	}

// 	return siw.Handler.V1JoinDomainByInvitation(c, domain, params)
// }

// // V1ListProblemSets operation middleware
// func (siw *ServerInterfaceWrapper) V1ListProblemSets(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1ListProblemSetsParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "ordering" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter ordering: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "offset" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter offset: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "limit" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter limit: %w", err).Error())
// 	}

// 	return siw.Handler.V1ListProblemSets(c, domain, params)
// }

// // V1CreateProblemSet operation middleware
// func (siw *ServerInterfaceWrapper) V1CreateProblemSet(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1CreateProblemSet(c, domain)
// }

// // V1DeleteProblemSet operation middleware
// func (siw *ServerInterfaceWrapper) V1DeleteProblemSet(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problemSet" -------------
// 	var problemSet string

// 	err = runtime.BindStyledParameter("simple", false, "problemSet", c.Params("problemSet"), &problemSet)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problemSet: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1DeleteProblemSet(c, domain, problemSet)
// }

// // V1GetProblemSet operation middleware
// func (siw *ServerInterfaceWrapper) V1GetProblemSet(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problemSet" -------------
// 	var problemSet string

// 	err = runtime.BindStyledParameter("simple", false, "problemSet", c.Params("problemSet"), &problemSet)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problemSet: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1GetProblemSet(c, domain, problemSet)
// }

// // V1UpdateProblemSet operation middleware
// func (siw *ServerInterfaceWrapper) V1UpdateProblemSet(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problemSet" -------------
// 	var problemSet string

// 	err = runtime.BindStyledParameter("simple", false, "problemSet", c.Params("problemSet"), &problemSet)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problemSet: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1UpdateProblemSet(c, domain, problemSet)
// }

// // V1ListProblemsInProblemSet operation middleware
// func (siw *ServerInterfaceWrapper) V1ListProblemsInProblemSet(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problemSet" -------------
// 	var problemSet string

// 	err = runtime.BindStyledParameter("simple", false, "problemSet", c.Params("problemSet"), &problemSet)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problemSet: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1ListProblemsInProblemSet(c, domain, problemSet)
// }

// // V1AddProblemInProblemSet operation middleware
// func (siw *ServerInterfaceWrapper) V1AddProblemInProblemSet(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problemSet" -------------
// 	var problemSet string

// 	err = runtime.BindStyledParameter("simple", false, "problemSet", c.Params("problemSet"), &problemSet)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problemSet: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1AddProblemInProblemSet(c, domain, problemSet)
// }

// // V1DeleteProblemInProblemSet operation middleware
// func (siw *ServerInterfaceWrapper) V1DeleteProblemInProblemSet(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problemSet" -------------
// 	var problemSet string

// 	err = runtime.BindStyledParameter("simple", false, "problemSet", c.Params("problemSet"), &problemSet)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problemSet: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problem" -------------
// 	var problem string

// 	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problem: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1DeleteProblemInProblemSet(c, domain, problemSet, problem)
// }

// // V1GetProblemInProblemSet operation middleware
// func (siw *ServerInterfaceWrapper) V1GetProblemInProblemSet(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problemSet" -------------
// 	var problemSet string

// 	err = runtime.BindStyledParameter("simple", false, "problemSet", c.Params("problemSet"), &problemSet)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problemSet: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problem" -------------
// 	var problem string

// 	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problem: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1GetProblemInProblemSet(c, domain, problemSet, problem)
// }

// // V1UpdateProblemInProblemSet operation middleware
// func (siw *ServerInterfaceWrapper) V1UpdateProblemInProblemSet(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problemSet" -------------
// 	var problemSet string

// 	err = runtime.BindStyledParameter("simple", false, "problemSet", c.Params("problemSet"), &problemSet)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problemSet: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problem" -------------
// 	var problem string

// 	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problem: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1UpdateProblemInProblemSet(c, domain, problemSet, problem)
// }

// // V1SubmitSolutionToProblemSet operation middleware
// func (siw *ServerInterfaceWrapper) V1SubmitSolutionToProblemSet(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problemSet" -------------
// 	var problemSet string

// 	err = runtime.BindStyledParameter("simple", false, "problemSet", c.Params("problemSet"), &problemSet)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problemSet: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problem" -------------
// 	var problem string

// 	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problem: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1SubmitSolutionToProblemSet(c, domain, problemSet, problem)
// }

// // V1ListProblems operation middleware
// func (siw *ServerInterfaceWrapper) V1ListProblems(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1ListProblemsParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "ordering" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter ordering: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "offset" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter offset: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "limit" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter limit: %w", err).Error())
// 	}

// 	return siw.Handler.V1ListProblems(c, domain, params)
// }

// // V1CreateProblem operation middleware
// func (siw *ServerInterfaceWrapper) V1CreateProblem(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1CreateProblem(c, domain)
// }

// // V1CloneProblem operation middleware
// func (siw *ServerInterfaceWrapper) V1CloneProblem(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1CloneProblem(c, domain)
// }

// // V1DeleteProblem operation middleware
// func (siw *ServerInterfaceWrapper) V1DeleteProblem(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problem" -------------
// 	var problem string

// 	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problem: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1DeleteProblem(c, domain, problem)
// }

// // V1GetProblem operation middleware
// func (siw *ServerInterfaceWrapper) V1GetProblem(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problem" -------------
// 	var problem string

// 	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problem: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1GetProblem(c, domain, problem)
// }

// // V1UpdateProblem operation middleware
// func (siw *ServerInterfaceWrapper) V1UpdateProblem(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problem" -------------
// 	var problem string

// 	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problem: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1UpdateProblem(c, domain, problem)
// }

// // V1SubmitSolutionToProblem operation middleware
// func (siw *ServerInterfaceWrapper) V1SubmitSolutionToProblem(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problem" -------------
// 	var problem string

// 	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problem: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1SubmitSolutionToProblem(c, domain, problem)
// }

// // V1ListProblemConfigCommits operation middleware
// func (siw *ServerInterfaceWrapper) V1ListProblemConfigCommits(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problem" -------------
// 	var problem string

// 	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problem: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1ListProblemConfigCommitsParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "ordering" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter ordering: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "offset" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter offset: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "limit" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter limit: %w", err).Error())
// 	}

// 	return siw.Handler.V1ListProblemConfigCommits(c, domain, problem, params)
// }

// // V1UpdateProblemConfigByArchive operation middleware
// func (siw *ServerInterfaceWrapper) V1UpdateProblemConfigByArchive(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problem" -------------
// 	var problem string

// 	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problem: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1UpdateProblemConfigByArchiveParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "configJsonOnMissing" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "configJsonOnMissing", query, &params.ConfigJsonOnMissing)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter configJsonOnMissing: %w", err).Error())
// 	}

// 	return siw.Handler.V1UpdateProblemConfigByArchive(c, domain, problem, params)
// }

// // V1UpdateProblemConfigJson operation middleware
// func (siw *ServerInterfaceWrapper) V1UpdateProblemConfigJson(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problem" -------------
// 	var problem string

// 	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problem: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1UpdateProblemConfigJson(c, domain, problem)
// }

// // V1DiffProblemConfigDefaultBranch operation middleware
// func (siw *ServerInterfaceWrapper) V1DiffProblemConfigDefaultBranch(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problem" -------------
// 	var problem string

// 	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problem: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1DiffProblemConfigDefaultBranchParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "after" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "after", query, &params.After)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter after: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "amount" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "amount", query, &params.Amount)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter amount: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "delimiter" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "delimiter", query, &params.Delimiter)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter delimiter: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "prefix" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "prefix", query, &params.Prefix)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter prefix: %w", err).Error())
// 	}

// 	return siw.Handler.V1DiffProblemConfigDefaultBranch(c, domain, problem, params)
// }

// // V1ListLatestProblemConfigObjectsUnderAGivenPrefix operation middleware
// func (siw *ServerInterfaceWrapper) V1ListLatestProblemConfigObjectsUnderAGivenPrefix(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problem" -------------
// 	var problem string

// 	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problem: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1ListLatestProblemConfigObjectsUnderAGivenPrefixParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "after" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "after", query, &params.After)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter after: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "amount" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "amount", query, &params.Amount)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter amount: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "delimiter" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "delimiter", query, &params.Delimiter)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter delimiter: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "prefix" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "prefix", query, &params.Prefix)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter prefix: %w", err).Error())
// 	}

// 	return siw.Handler.V1ListLatestProblemConfigObjectsUnderAGivenPrefix(c, domain, problem, params)
// }

// // V1DownloadProblemConfigArchive operation middleware
// func (siw *ServerInterfaceWrapper) V1DownloadProblemConfigArchive(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problem" -------------
// 	var problem string

// 	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problem: %w", err).Error())
// 	}

// 	// ------------- Path parameter "config" -------------
// 	var config string

// 	err = runtime.BindStyledParameter("simple", false, "config", c.Params("config"), &config)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter config: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1DownloadProblemConfigArchiveParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "archiveFormat" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "archiveFormat", query, &params.ArchiveFormat)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter archiveFormat: %w", err).Error())
// 	}

// 	return siw.Handler.V1DownloadProblemConfigArchive(c, domain, problem, config, params)
// }

// // V1GetProblemConfigJson operation middleware
// func (siw *ServerInterfaceWrapper) V1GetProblemConfigJson(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "problem" -------------
// 	var problem string

// 	err = runtime.BindStyledParameter("simple", false, "problem", c.Params("problem"), &problem)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problem: %w", err).Error())
// 	}

// 	// ------------- Path parameter "config" -------------
// 	var config string

// 	err = runtime.BindStyledParameter("simple", false, "config", c.Params("config"), &config)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter config: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1GetProblemConfigJson(c, domain, problem, config)
// }

// // V1ListRecordsInDomain operation middleware
// func (siw *ServerInterfaceWrapper) V1ListRecordsInDomain(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1ListRecordsInDomainParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "problemSet" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "problemSet", query, &params.ProblemSet)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problemSet: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "problem" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "problem", query, &params.Problem)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter problem: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "submitterId" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "submitterId", query, &params.SubmitterId)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter submitterId: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "ordering" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter ordering: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "offset" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter offset: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "limit" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter limit: %w", err).Error())
// 	}

// 	return siw.Handler.V1ListRecordsInDomain(c, domain, params)
// }

// // V1GetRecord operation middleware
// func (siw *ServerInterfaceWrapper) V1GetRecord(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "record" -------------
// 	var record openapi_types.UUID

// 	err = runtime.BindStyledParameter("simple", false, "record", c.Params("record"), &record)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter record: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1GetRecord(c, domain, record)
// }

// // V1SubmitCaseByJudger operation middleware
// func (siw *ServerInterfaceWrapper) V1SubmitCaseByJudger(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "record" -------------
// 	var record string

// 	err = runtime.BindStyledParameter("simple", false, "record", c.Params("record"), &record)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter record: %w", err).Error())
// 	}

// 	// ------------- Path parameter "index" -------------
// 	var index int

// 	err = runtime.BindStyledParameter("simple", false, "index", c.Params("index"), &index)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter index: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1SubmitCaseByJudger(c, domain, record, index)
// }

// // V1SubmitRecordByJudger operation middleware
// func (siw *ServerInterfaceWrapper) V1SubmitRecordByJudger(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "record" -------------
// 	var record string

// 	err = runtime.BindStyledParameter("simple", false, "record", c.Params("record"), &record)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter record: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1SubmitRecordByJudger(c, domain, record)
// }

// // V1ClaimRecordByJudger operation middleware
// func (siw *ServerInterfaceWrapper) V1ClaimRecordByJudger(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "record" -------------
// 	var record string

// 	err = runtime.BindStyledParameter("simple", false, "record", c.Params("record"), &record)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter record: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1ClaimRecordByJudger(c, domain, record)
// }

// // V1ListDomainRoles operation middleware
// func (siw *ServerInterfaceWrapper) V1ListDomainRoles(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1ListDomainRolesParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "ordering" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter ordering: %w", err).Error())
// 	}

// 	return siw.Handler.V1ListDomainRoles(c, domain, params)
// }

// // V1CreateDomainRole operation middleware
// func (siw *ServerInterfaceWrapper) V1CreateDomainRole(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1CreateDomainRole(c, domain)
// }

// // V1DeleteDomainRole operation middleware
// func (siw *ServerInterfaceWrapper) V1DeleteDomainRole(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "role" -------------
// 	var role string

// 	err = runtime.BindStyledParameter("simple", false, "role", c.Params("role"), &role)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter role: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1DeleteDomainRole(c, domain, role)
// }

// // V1GetDomainRole operation middleware
// func (siw *ServerInterfaceWrapper) V1GetDomainRole(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "role" -------------
// 	var role string

// 	err = runtime.BindStyledParameter("simple", false, "role", c.Params("role"), &role)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter role: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1GetDomainRole(c, domain, role)
// }

// // V1UpdateDomainRole operation middleware
// func (siw *ServerInterfaceWrapper) V1UpdateDomainRole(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "role" -------------
// 	var role string

// 	err = runtime.BindStyledParameter("simple", false, "role", c.Params("role"), &role)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter role: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1UpdateDomainRole(c, domain, role)
// }

// // V1TransferDomain operation middleware
// func (siw *ServerInterfaceWrapper) V1TransferDomain(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1TransferDomain(c, domain)
// }

// // V1ListDomainUsers operation middleware
// func (siw *ServerInterfaceWrapper) V1ListDomainUsers(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1ListDomainUsersParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "ordering" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter ordering: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "offset" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter offset: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "limit" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter limit: %w", err).Error())
// 	}

// 	return siw.Handler.V1ListDomainUsers(c, domain, params)
// }

// // V1AddDomainUser operation middleware
// func (siw *ServerInterfaceWrapper) V1AddDomainUser(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1AddDomainUser(c, domain)
// }

// // V1RemoveDomainUser operation middleware
// func (siw *ServerInterfaceWrapper) V1RemoveDomainUser(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "user" -------------
// 	var user string

// 	err = runtime.BindStyledParameter("simple", false, "user", c.Params("user"), &user)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter user: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1RemoveDomainUser(c, domain, user)
// }

// // V1GetDomainUser operation middleware
// func (siw *ServerInterfaceWrapper) V1GetDomainUser(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "user" -------------
// 	var user string

// 	err = runtime.BindStyledParameter("simple", false, "user", c.Params("user"), &user)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter user: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1GetDomainUser(c, domain, user)
// }

// // V1UpdateDomainUser operation middleware
// func (siw *ServerInterfaceWrapper) V1UpdateDomainUser(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "user" -------------
// 	var user string

// 	err = runtime.BindStyledParameter("simple", false, "user", c.Params("user"), &user)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter user: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1UpdateDomainUser(c, domain, user)
// }

// // V1GetDomainUserPermission operation middleware
// func (siw *ServerInterfaceWrapper) V1GetDomainUserPermission(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "domain" -------------
// 	var domain string

// 	err = runtime.BindStyledParameter("simple", false, "domain", c.Params("domain"), &domain)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter domain: %w", err).Error())
// 	}

// 	// ------------- Path parameter "user" -------------
// 	var user string

// 	err = runtime.BindStyledParameter("simple", false, "user", c.Params("user"), &user)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter user: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1GetDomainUserPermission(c, domain, user)
// }

// // V1JwtDecoded operation middleware
// func (siw *ServerInterfaceWrapper) V1JwtDecoded(c *fiber.Ctx) error {

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1JwtDecoded(c)
// }

// // V1ListProblemGroups operation middleware
// func (siw *ServerInterfaceWrapper) V1ListProblemGroups(c *fiber.Ctx) error {

// 	var err error

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	// Parameter object where we will unmarshal all parameters from the context
// 	var params model.V1ListProblemGroupsParams

// 	query, err := url.ParseQuery(string(c.Request().URI().QueryString()))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "ordering" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "ordering", query, &params.Ordering)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter ordering: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "offset" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "offset", query, &params.Offset)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter offset: %w", err).Error())
// 	}

// 	// ------------- Optional query parameter "limit" -------------

// 	err = runtime.BindQueryParameter("form", true, false, "limit", query, &params.Limit)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter limit: %w", err).Error())
// 	}

// 	return siw.Handler.V1ListProblemGroups(c, params)
// }

// // V1TestErrorReport operation middleware
// func (siw *ServerInterfaceWrapper) V1TestErrorReport(c *fiber.Ctx) error {

// 	return siw.Handler.V1TestErrorReport(c)
// }

// // V1GetCurrentUser operation middleware
// func (siw *ServerInterfaceWrapper) V1GetCurrentUser(c *fiber.Ctx) error {

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1GetCurrentUser(c)
// }

// // V1UpdateCurrentUser operation middleware
// func (siw *ServerInterfaceWrapper) V1UpdateCurrentUser(c *fiber.Ctx) error {

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1UpdateCurrentUser(c)
// }

// // V1ChangePassword operation middleware
// func (siw *ServerInterfaceWrapper) V1ChangePassword(c *fiber.Ctx) error {

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1ChangePassword(c)
// }

// // V1GetUser operation middleware
// func (siw *ServerInterfaceWrapper) V1GetUser(c *fiber.Ctx) error {

// 	var err error

// 	// ------------- Path parameter "uid" -------------
// 	var uid string

// 	err = runtime.BindStyledParameter("simple", false, "uid", c.Params("uid"), &uid)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter uid: %w", err).Error())
// 	}

// 	c.Context().SetUserValue(model.HTTPBearerScopes, []string{""})

// 	return siw.Handler.V1GetUser(c, uid)
// }

// // V1Version operation middleware
// func (siw *ServerInterfaceWrapper) V1Version(c *fiber.Ctx) error {

// 	return siw.Handler.V1Version(c)
// }
