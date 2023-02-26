package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/service"
	casbin "github.com/joint-online-judge/go-horse/pkg/fibercasbin"
	"github.com/joint-online-judge/go-horse/platform/auth"
)

var Perm *casbin.Middleware

func InitCasbinMiddleware() {
	Perm = casbin.New(casbin.Config{
		Enforcer:      auth.Enforcer,
		PolicyAdapter: auth.Adapter,
		Lookup: func(c *fiber.Ctx) string {
			return service.Auth(c).JWTUser().Username
		},
		LookupDomain: func(c *fiber.Ctx) string {
			return c.Locals("domainId").(string)
		},
	})
}
