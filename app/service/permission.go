package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/platform/auth"
)

var SiteContext = casbin.NewEnforceContext("2")

type permissionImpl struct {
	c *fiber.Ctx
	e *casbin.Enforcer
}

func Permission(c *fiber.Ctx) *permissionImpl {
	return &permissionImpl{
		c: c,
		e: auth.Enforcer,
	}
}

func (s *permissionImpl) UpdateUserSiteRole(username string, role schema.DefaultRole) error {
	//s.e.AddNamedGroupingPolicy("g2", username, role)
	return nil
}

func (s *permissionImpl) CreateDomainRole(domain string, role string, permission schema.DomainPermission) error {
	//ok, _ := s.e.AddPoliciesEx()
	return nil
}

// InitDomainRoles Pop in all default roles and permissions
func (s *permissionImpl) InitDomainRoles(domain string) error {
	//s.e.AddPoliciesEx([][]string{
	//	{"admin", domain, "problem", "view"},
	//	{"admin", domain, "problem", "edit"},
	//	{"admin", domain, "problem", "submit"},
	//	{"user", domain, "general", "view"},
	//})
	//s.e.AddNamedPoliciesEx("p2", [][]string{
	//	{"root", "domain", "create"},
	//	{"root", "domain", "delete"},
	//	{"guest", "domain", "join"},
	//})
	//
	//s.e.AddGroupingPoliciesEx([][]string{
	//	{"nichujie", "admin", domain},
	//	{"boyanzh", "user", domain},
	//})
	//s.e.AddNamedGroupingPolicies("g2", [][]string{
	//	{"nichujie", "root"},
	//	{"boyanzh", "guest"},
	//})
	//ok, _ := s.e.Enforce("nichujie", domain, "problem", "view")
	//
	//ctx := casbin.NewEnforceContext("2")
	//ok, _ = s.e.Enforce(ctx, "nichujie", "domain", "create")

	return nil
}

func (s *permissionImpl) UpdateDomainPermission(domain string) error {
	//auth.Enforcer.AddNamedDomainMatchingFunc()
	return nil
}
