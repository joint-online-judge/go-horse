package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/schemas"
)

func RegisterDomain(router fiber.Router, wrapper schemas.ServerInterfaceWrapper) {
	domains := router.Group("/domains")
	domain := domains.Group("/:domain")
	domain_users := domain.Group("/users")
	domain_user := domain_users.Group("/:user")
	domains.Get("", wrapper.ListDomains)
	domains.Post("", wrapper.CreateDomain)
	domains.Get("/groups", wrapper.SearchDomainGroups)
	domain.Delete("", wrapper.DeleteDomain)
	domain.Get("", wrapper.GetDomain)
	domain.Patch("", wrapper.UpdateDomain)
	domain.Get("/candidates", wrapper.SearchDomainCandidates)
	domain.Get("/invitations", wrapper.ListDomainInvitations)
	domain.Post("/invitations", wrapper.CreateDomainInvitation)
	domain.Delete("/invitations/:invitation", wrapper.DeleteDomainInvitation)
	domain.Get("/invitations/:invitation", wrapper.GetDomainInvitation)
	domain.Patch("/invitations/:invitation", wrapper.UpdateDomainInvitation)
	domain.Post("/join", wrapper.JoinDomainByInvitation)
	domain.Get("/roles", wrapper.ListDomainRoles)
	domain.Post("/roles", wrapper.CreateDomainRole)
	domain.Delete("/roles/:role", wrapper.DeleteDomainRole)
	domain.Get("/roles/:role", wrapper.GetDomainRole)
	domain.Patch("/roles/:role", wrapper.UpdateDomainRole)
	domain.Post("/transfer", wrapper.TransferDomain)
	domain_users.Get("", wrapper.ListDomainUsers)
	domain_users.Post("", wrapper.AddDomainUser)
	domain_user.Delete("/users/:user", wrapper.RemoveDomainUser)
	domain_user.Get("/users/:user", wrapper.GetDomainUser)
	domain_user.Patch("/users/:user", wrapper.UpdateDomainUser)
	domain_user.Get("/users/:user/permission", wrapper.GetDomainUserPermission)
}
