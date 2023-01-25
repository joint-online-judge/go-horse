package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schemas"
)

func RegisterProblemGroup(
	router fiber.Router,
	wrapper schemas.ServerInterfaceWrapper,
) {
	problem_groups := router.Group("/problem_groups")
	problem_groups.Get("", wrapper.ListProblemGroups)
}
