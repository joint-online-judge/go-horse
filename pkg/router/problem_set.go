package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/pkg/middleware"
)

func RegisterProblemSet(
	router fiber.Router,
	wrapper schema.ServerInterfaceWrapper,
) {
	problem_sets := router.Group(
		"/domains/:domain/problem_sets",
		middleware.Domain,
	)
	problem_set := problem_sets.Group("/:problemSet", middleware.ProblemSet)
	problems := problem_set.Group("/problems")
	problem := problems.Group("/:problem", middleware.ProblemInProblemSet)
	problem_sets.Get("", wrapper.ListProblemSets)
	problem_sets.Post("", wrapper.CreateProblemSet)
	problem_set.Delete("", wrapper.DeleteProblemSet)
	problem_set.Get("", wrapper.GetProblemSet)
	problem_set.Patch("", wrapper.UpdateProblemSet)
	problems.Get("", wrapper.ListProblemsInProblemSet)
	problems.Post("", wrapper.AddProblemInProblemSet)
	problem.Get("", wrapper.GetProblemInProblemSet)
	problem.Delete("", wrapper.DeleteProblemInProblemSet)
	problem.Patch("", wrapper.UpdateProblemInProblemSet)
	problem.Post("/submit", wrapper.SubmitSolutionToProblemInProblemSet)
}
