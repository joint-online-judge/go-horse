// Package router provides primitives to interact with the openapi HTTP API.
package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/handlers"
)

type MiddlewareFunc fiber.Handler

// FiberServerOptions provides options for the Fiber server.
type FiberServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router fiber.Router, si handlers.ServerInterface) {
	RegisterHandlersWithOptions(router, si, FiberServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router fiber.Router, si handlers.ServerInterface, options FiberServerOptions) {
	wrapper := handlers.ServerInterfaceWrapper{
		Handler: si,
	}

	for _, m := range options.Middlewares {
		router.Use(m)
	}

	router.Get(options.BaseURL+"/admin/domain_roles", wrapper.V1AdminListDomainRoles)

	router.Get(options.BaseURL+"/admin/judgers", wrapper.V1AdminListJudgers)

	router.Post(options.BaseURL+"/admin/judgers", wrapper.V1AdminCreateJudger)

	router.Get(options.BaseURL+"/admin/users", wrapper.V1AdminListUsers)

	router.Get(options.BaseURL+"/admin/:uid", wrapper.V1AdminGetUser)

	router.Get(options.BaseURL+"/admin/:uid/domains", wrapper.V1AdminListUserDomains)

	router.Post(options.BaseURL+"/auth/login", wrapper.V1Login)

	router.Post(options.BaseURL+"/auth/logout", wrapper.V1Logout)

	router.Get(options.BaseURL+"/auth/oauth2", wrapper.V1ListOauth2)

	router.Get(options.BaseURL+"/auth/oauth2/:oauth2/authorize", wrapper.V1OauthAuthorize)

	router.Post(options.BaseURL+"/auth/refresh", wrapper.V1Refresh)

	router.Post(options.BaseURL+"/auth/register", wrapper.V1Register)

	router.Get(options.BaseURL+"/auth/token", wrapper.V1GetToken)

	router.Get(options.BaseURL+"/domains", wrapper.V1ListDomains)

	router.Post(options.BaseURL+"/domains", wrapper.V1CreateDomain)

	router.Get(options.BaseURL+"/domains/groups", wrapper.V1SearchDomainGroups)

	router.Delete(options.BaseURL+"/domains/:domain", wrapper.V1DeleteDomain)

	router.Get(options.BaseURL+"/domains/:domain", wrapper.V1GetDomain)

	router.Patch(options.BaseURL+"/domains/:domain", wrapper.V1UpdateDomain)

	router.Get(options.BaseURL+"/domains/:domain/candidates", wrapper.V1SearchDomainCandidates)

	router.Get(options.BaseURL+"/domains/:domain/invitations", wrapper.V1ListDomainInvitations)

	router.Post(options.BaseURL+"/domains/:domain/invitations", wrapper.V1CreateDomainInvitation)

	router.Delete(options.BaseURL+"/domains/:domain/invitations/:invitation", wrapper.V1DeleteDomainInvitation)

	router.Get(options.BaseURL+"/domains/:domain/invitations/:invitation", wrapper.V1GetDomainInvitation)

	router.Patch(options.BaseURL+"/domains/:domain/invitations/:invitation", wrapper.V1UpdateDomainInvitation)

	router.Post(options.BaseURL+"/domains/:domain/join", wrapper.V1JoinDomainByInvitation)

	router.Get(options.BaseURL+"/domains/:domain/problem_sets", wrapper.V1ListProblemSets)

	router.Post(options.BaseURL+"/domains/:domain/problem_sets", wrapper.V1CreateProblemSet)

	router.Delete(options.BaseURL+"/domains/:domain/problem_sets/:problemSet", wrapper.V1DeleteProblemSet)

	router.Get(options.BaseURL+"/domains/:domain/problem_sets/:problemSet", wrapper.V1GetProblemSet)

	router.Patch(options.BaseURL+"/domains/:domain/problem_sets/:problemSet", wrapper.V1UpdateProblemSet)

	router.Get(options.BaseURL+"/domains/:domain/problem_sets/:problemSet/problems", wrapper.V1ListProblemsInProblemSet)

	router.Post(options.BaseURL+"/domains/:domain/problem_sets/:problemSet/problems", wrapper.V1AddProblemInProblemSet)

	router.Delete(options.BaseURL+"/domains/:domain/problem_sets/:problemSet/problems/:problem", wrapper.V1DeleteProblemInProblemSet)

	router.Get(options.BaseURL+"/domains/:domain/problem_sets/:problemSet/problems/:problem", wrapper.V1GetProblemInProblemSet)

	router.Patch(options.BaseURL+"/domains/:domain/problem_sets/:problemSet/problems/:problem", wrapper.V1UpdateProblemInProblemSet)

	router.Post(options.BaseURL+"/domains/:domain/problem_sets/:problemSet/problems/:problem/submit", wrapper.V1SubmitSolutionToProblemSet)

	router.Get(options.BaseURL+"/domains/:domain/problems", wrapper.V1ListProblems)

	router.Post(options.BaseURL+"/domains/:domain/problems", wrapper.V1CreateProblem)

	router.Post(options.BaseURL+"/domains/:domain/problems/clone", wrapper.V1CloneProblem)

	router.Delete(options.BaseURL+"/domains/:domain/problems/:problem", wrapper.V1DeleteProblem)

	router.Get(options.BaseURL+"/domains/:domain/problems/:problem", wrapper.V1GetProblem)

	router.Patch(options.BaseURL+"/domains/:domain/problems/:problem", wrapper.V1UpdateProblem)

	router.Post(options.BaseURL+"/domains/:domain/problems/:problem", wrapper.V1SubmitSolutionToProblem)

	router.Get(options.BaseURL+"/domains/:domain/problems/:problem/configs", wrapper.V1ListProblemConfigCommits)

	router.Post(options.BaseURL+"/domains/:domain/problems/:problem/configs", wrapper.V1UpdateProblemConfigByArchive)

	router.Post(options.BaseURL+"/domains/:domain/problems/:problem/configs/json", wrapper.V1UpdateProblemConfigJson)

	router.Get(options.BaseURL+"/domains/:domain/problems/:problem/configs/latest/diff", wrapper.V1DiffProblemConfigDefaultBranch)

	router.Get(options.BaseURL+"/domains/:domain/problems/:problem/configs/latest/ls", wrapper.V1ListLatestProblemConfigObjectsUnderAGivenPrefix)

	router.Get(options.BaseURL+"/domains/:domain/problems/:problem/configs/:config", wrapper.V1DownloadProblemConfigArchive)

	router.Get(options.BaseURL+"/domains/:domain/problems/:problem/configs/:config/json", wrapper.V1GetProblemConfigJson)

	router.Get(options.BaseURL+"/domains/:domain/records", wrapper.V1ListRecordsInDomain)

	router.Get(options.BaseURL+"/domains/:domain/records/:record", wrapper.V1GetRecord)

	router.Put(options.BaseURL+"/domains/:domain/records/:record/cases/:index/judge", wrapper.V1SubmitCaseByJudger)

	router.Put(options.BaseURL+"/domains/:domain/records/:record/judge", wrapper.V1SubmitRecordByJudger)

	router.Post(options.BaseURL+"/domains/:domain/records/:record/judge/claim", wrapper.V1ClaimRecordByJudger)

	router.Get(options.BaseURL+"/domains/:domain/roles", wrapper.V1ListDomainRoles)

	router.Post(options.BaseURL+"/domains/:domain/roles", wrapper.V1CreateDomainRole)

	router.Delete(options.BaseURL+"/domains/:domain/roles/:role", wrapper.V1DeleteDomainRole)

	router.Get(options.BaseURL+"/domains/:domain/roles/:role", wrapper.V1GetDomainRole)

	router.Patch(options.BaseURL+"/domains/:domain/roles/:role", wrapper.V1UpdateDomainRole)

	router.Post(options.BaseURL+"/domains/:domain/transfer", wrapper.V1TransferDomain)

	router.Get(options.BaseURL+"/domains/:domain/users", wrapper.V1ListDomainUsers)

	router.Post(options.BaseURL+"/domains/:domain/users", wrapper.V1AddDomainUser)

	router.Delete(options.BaseURL+"/domains/:domain/users/:user", wrapper.V1RemoveDomainUser)

	router.Get(options.BaseURL+"/domains/:domain/users/:user", wrapper.V1GetDomainUser)

	router.Patch(options.BaseURL+"/domains/:domain/users/:user", wrapper.V1UpdateDomainUser)

	router.Get(options.BaseURL+"/domains/:domain/users/:user/permission", wrapper.V1GetDomainUserPermission)

	router.Get(options.BaseURL+"/jwt_decoded", wrapper.V1JwtDecoded)

	router.Get(options.BaseURL+"/problem_groups", wrapper.V1ListProblemGroups)

	router.Get(options.BaseURL+"/test/report", wrapper.V1TestErrorReport)

	router.Get(options.BaseURL+"/users/me", wrapper.V1GetCurrentUser)

	router.Patch(options.BaseURL+"/users/me", wrapper.V1UpdateCurrentUser)

	router.Patch(options.BaseURL+"/users/me/password", wrapper.V1ChangePassword)

	router.Get(options.BaseURL+"/users/:uid", wrapper.V1GetUser)

	router.Get(options.BaseURL+"/version", wrapper.V1Version)

}
