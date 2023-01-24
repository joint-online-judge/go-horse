package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schemas"
)

func RegisterProblem(router fiber.Router, wrapper schemas.ServerInterfaceWrapper) {
	problems := router.Group("/domains/:domain/problems")
	problem := problems.Group("/:problem")
	problems.Get("", wrapper.ListProblems)
	problems.Post("", wrapper.CreateProblem)
	problems.Post("/clone", wrapper.CloneProblem)
	problem.Get("", wrapper.GetProblem)
	problem.Post("", wrapper.SubmitSolutionToProblem)
	problem.Delete("", wrapper.DeleteProblem)
	problem.Patch("", wrapper.UpdateProblem)
}
