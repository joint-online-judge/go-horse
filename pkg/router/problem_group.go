package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schema"
)

func RegisterProblemGroup(
	router fiber.Router,
	wrapper schema.ServerInterfaceWrapper,
) {
	problem_groups := router.Group("/problem_groups")
	problem_groups.Get("", wrapper.ListProblemGroups)
}
