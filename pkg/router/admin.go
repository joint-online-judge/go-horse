package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schema"
)

func RegisterAdmin(
	router fiber.Router,
	wrapper schema.ServerInterfaceWrapper,
) {
	admin := router.Group("/admin")
	admin.Get("/domain_roles", wrapper.AdminListDomainRoles)
	admin.Get("/judgers", wrapper.AdminListJudgers)
	admin.Post("/judgers", wrapper.AdminCreateJudger)
	admin.Get("/users", wrapper.AdminListUsers)
	admin.Get("/:uid", wrapper.AdminGetUser)
	admin.Get("/:uid/domains", wrapper.AdminListUserDomains)
}
