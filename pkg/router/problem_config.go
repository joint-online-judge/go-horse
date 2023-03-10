package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/pkg/middleware"
)

func RegisterProblemConfig(
	router fiber.Router,
	wrapper schema.ServerInterfaceWrapper,
) {
	problem_configs := router.Group(
		"/domains/:domain/problems/:problem/configs",
		middleware.Domain,
	)
	problem_config := problem_configs.Group("/:config")
	latest_problem_config := problem_configs.Group("/latest")
	problem_configs.Get("", wrapper.ListProblemConfigCommits)
	problem_configs.Post("", wrapper.UpdateProblemConfigByArchive)
	problem_configs.Post("/json", wrapper.UpdateProblemConfigJson)
	problem_config.Get("", wrapper.DownloadProblemConfigArchive)
	problem_config.Get("/json", wrapper.GetProblemConfigJson)
	latest_problem_config.Get("/diff", wrapper.DiffProblemConfigDefaultBranch)
	latest_problem_config.Get(
		"/ls",
		wrapper.ListLatestProblemConfigObjectsUnderAGivenPrefix,
	)
}
