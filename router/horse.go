// Package router provides primitives to interact with the openapi HTTP API.
package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/types"
)

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterStrictHandlers(router fiber.Router, ssi types.StrictServerInterface, middlewares []types.StrictMiddlewareFunc) {
	wrapper := types.ServerInterfaceWrapper{
		Handler: types.NewStrictHandler(ssi, middlewares),
	}

	router.Get("/admin/domain_roles", wrapper.V1AdminListDomainRoles)

	router.Get("/admin/judgers", wrapper.V1AdminListJudgers)

	router.Post("/admin/judgers", wrapper.V1AdminCreateJudger)

	router.Get("/admin/users", wrapper.V1AdminListUsers)

	router.Get("/admin/:uid", wrapper.V1AdminGetUser)

	router.Get("/admin/:uid/domains", wrapper.V1AdminListUserDomains)

	router.Post("/auth/login", wrapper.V1Login)

	router.Post("/auth/logout", wrapper.V1Logout)

	router.Get("/auth/oauth2", wrapper.V1ListOauth2)

	router.Get("/auth/oauth2/:oauth2/authorize", wrapper.V1OauthAuthorize)

	router.Post("/auth/refresh", wrapper.V1Refresh)

	router.Post("/auth/register", wrapper.V1Register)

	router.Get("/auth/token", wrapper.V1GetToken)

	router.Get("/domains", wrapper.V1ListDomains)

	router.Post("/domains", wrapper.V1CreateDomain)

	router.Get("/domains/groups", wrapper.V1SearchDomainGroups)

	router.Delete("/domains/:domain", wrapper.V1DeleteDomain)

	router.Get("/domains/:domain", wrapper.V1GetDomain)

	router.Patch("/domains/:domain", wrapper.V1UpdateDomain)

	router.Get("/domains/:domain/candidates", wrapper.V1SearchDomainCandidates)

	router.Get("/domains/:domain/invitations", wrapper.V1ListDomainInvitations)

	router.Post("/domains/:domain/invitations", wrapper.V1CreateDomainInvitation)

	router.Delete("/domains/:domain/invitations/:invitation", wrapper.V1DeleteDomainInvitation)

	router.Get("/domains/:domain/invitations/:invitation", wrapper.V1GetDomainInvitation)

	router.Patch("/domains/:domain/invitations/:invitation", wrapper.V1UpdateDomainInvitation)

	router.Post("/domains/:domain/join", wrapper.V1JoinDomainByInvitation)

	router.Get("/domains/:domain/problem_sets", wrapper.V1ListProblemSets)

	router.Post("/domains/:domain/problem_sets", wrapper.V1CreateProblemSet)

	router.Delete("/domains/:domain/problem_sets/:problemSet", wrapper.V1DeleteProblemSet)

	router.Get("/domains/:domain/problem_sets/:problemSet", wrapper.V1GetProblemSet)

	router.Patch("/domains/:domain/problem_sets/:problemSet", wrapper.V1UpdateProblemSet)

	router.Get("/domains/:domain/problem_sets/:problemSet/problems", wrapper.V1ListProblemsInProblemSet)

	router.Post("/domains/:domain/problem_sets/:problemSet/problems", wrapper.V1AddProblemInProblemSet)

	router.Delete("/domains/:domain/problem_sets/:problemSet/problems/:problem", wrapper.V1DeleteProblemInProblemSet)

	router.Get("/domains/:domain/problem_sets/:problemSet/problems/:problem", wrapper.V1GetProblemInProblemSet)

	router.Patch("/domains/:domain/problem_sets/:problemSet/problems/:problem", wrapper.V1UpdateProblemInProblemSet)

	router.Post("/domains/:domain/problem_sets/:problemSet/problems/:problem/submit", wrapper.V1SubmitSolutionToProblemSet)

	router.Get("/domains/:domain/problems", wrapper.V1ListProblems)

	router.Post("/domains/:domain/problems", wrapper.V1CreateProblem)

	router.Post("/domains/:domain/problems/clone", wrapper.V1CloneProblem)

	router.Delete("/domains/:domain/problems/:problem", wrapper.V1DeleteProblem)

	router.Get("/domains/:domain/problems/:problem", wrapper.V1GetProblem)

	router.Patch("/domains/:domain/problems/:problem", wrapper.V1UpdateProblem)

	router.Post("/domains/:domain/problems/:problem", wrapper.V1SubmitSolutionToProblem)

	router.Get("/domains/:domain/problems/:problem/configs", wrapper.V1ListProblemConfigCommits)

	router.Post("/domains/:domain/problems/:problem/configs", wrapper.V1UpdateProblemConfigByArchive)

	router.Post("/domains/:domain/problems/:problem/configs/json", wrapper.V1UpdateProblemConfigJson)

	router.Get("/domains/:domain/problems/:problem/configs/latest/diff", wrapper.V1DiffProblemConfigDefaultBranch)

	router.Get("/domains/:domain/problems/:problem/configs/latest/ls", wrapper.V1ListLatestProblemConfigObjectsUnderAGivenPrefix)

	router.Get("/domains/:domain/problems/:problem/configs/:config", wrapper.V1DownloadProblemConfigArchive)

	router.Get("/domains/:domain/problems/:problem/configs/:config/json", wrapper.V1GetProblemConfigJson)

	router.Get("/domains/:domain/records", wrapper.V1ListRecordsInDomain)

	router.Get("/domains/:domain/records/:record", wrapper.V1GetRecord)

	router.Put("/domains/:domain/records/:record/cases/:index/judge", wrapper.V1SubmitCaseByJudger)

	router.Put("/domains/:domain/records/:record/judge", wrapper.V1SubmitRecordByJudger)

	router.Post("/domains/:domain/records/:record/judge/claim", wrapper.V1ClaimRecordByJudger)

	router.Get("/domains/:domain/roles", wrapper.V1ListDomainRoles)

	router.Post("/domains/:domain/roles", wrapper.V1CreateDomainRole)

	router.Delete("/domains/:domain/roles/:role", wrapper.V1DeleteDomainRole)

	router.Get("/domains/:domain/roles/:role", wrapper.V1GetDomainRole)

	router.Patch("/domains/:domain/roles/:role", wrapper.V1UpdateDomainRole)

	router.Post("/domains/:domain/transfer", wrapper.V1TransferDomain)

	router.Get("/domains/:domain/users", wrapper.V1ListDomainUsers)

	router.Post("/domains/:domain/users", wrapper.V1AddDomainUser)

	router.Delete("/domains/:domain/users/:user", wrapper.V1RemoveDomainUser)

	router.Get("/domains/:domain/users/:user", wrapper.V1GetDomainUser)

	router.Patch("/domains/:domain/users/:user", wrapper.V1UpdateDomainUser)

	router.Get("/domains/:domain/users/:user/permission", wrapper.V1GetDomainUserPermission)

	router.Get("/jwt_decoded", wrapper.V1JwtDecoded)

	router.Get("/problem_groups", wrapper.V1ListProblemGroups)

	router.Get("/test/report", wrapper.V1TestErrorReport)

	router.Get("/users/me", wrapper.V1GetCurrentUser)

	router.Patch("/users/me", wrapper.V1UpdateCurrentUser)

	router.Patch("/users/me/password", wrapper.V1ChangePassword)

	router.Get("/users/:uid", wrapper.V1GetUser)

	router.Get("/version", wrapper.V1Version)

}
