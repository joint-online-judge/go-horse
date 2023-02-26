package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/pkg/middleware"
)

func RegisterProblem(
	router fiber.Router,
	wrapper schema.ServerInterfaceWrapper,
) {
	problems := router.Group("/domains/:domain/problems", middleware.Domain)
	problem := problems.Group("/:problem")
	problems.Get("", wrapper.ListProblems)
	problems.Post("", wrapper.CreateProblem)
	problems.Post("/clone", wrapper.CloneProblem)
	problem.Get("", wrapper.GetProblem)
	problem.Post("", wrapper.SubmitSolutionToProblem)
	problem.Delete("", wrapper.DeleteProblem)
	problem.Patch("", wrapper.UpdateProblem)
}
