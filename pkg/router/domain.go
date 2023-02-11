package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schema"
)

func RegisterDomain(
	router fiber.Router,
	wrapper schema.ServerInterfaceWrapper,
) {
	domains := router.Group("/domains")
	domain := domains.Group("/:domain")
	domain_invitations := domain.Group("/invitations")
	domain_invitation := domain_invitations.Group("/:invitation")
	domain_roles := domain.Group("/roles")
	domain_role := domain_roles.Group(":/role")
	domain_users := domain.Group("/users")
	domain_user := domain_users.Group("/:user")
	domains.Get("", wrapper.ListDomains)
	domains.Post("", wrapper.CreateDomain)
	domains.Get("/groups", wrapper.SearchDomainGroups)
	domain.Delete("", wrapper.DeleteDomain)
	domain.Get("", wrapper.GetDomain)
	domain.Patch("", wrapper.UpdateDomain)
	domain.Get("/candidates", wrapper.SearchDomainCandidates)
	domain_invitations.Get("", wrapper.ListDomainInvitations)
	domain_invitations.Post("", wrapper.CreateDomainInvitation)
	domain_invitation.Delete("", wrapper.DeleteDomainInvitation)
	domain_invitation.Get("", wrapper.GetDomainInvitation)
	domain_invitation.Patch("", wrapper.UpdateDomainInvitation)
	domain.Post("/join", wrapper.JoinDomainByInvitation)
	domain_roles.Get("", wrapper.ListDomainRoles)
	domain_roles.Post("", wrapper.CreateDomainRole)
	domain_role.Delete("", wrapper.DeleteDomainRole)
	domain_role.Get("", wrapper.GetDomainRole)
	domain_role.Patch("", wrapper.UpdateDomainRole)
	domain.Post("/transfer", wrapper.TransferDomain)
	domain_users.Get("", wrapper.ListDomainUsers)
	domain_users.Post("", wrapper.AddDomainUser)
	domain_user.Delete("", wrapper.RemoveDomainUser)
	domain_user.Get("", wrapper.GetDomainUser)
	domain_user.Patch("", wrapper.UpdateDomainUser)
	domain_user.Get("/permission", wrapper.GetDomainUserPermission)
}
