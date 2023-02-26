package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/service"
	"github.com/joint-online-judge/go-horse/pkg/config"
	fibercasbin "github.com/joint-online-judge/go-horse/pkg/fibercasbin"
	"github.com/joint-online-judge/go-horse/platform/casbin"
)

var Perm *fibercasbin.Middleware

func InitCasbinMiddleware() {
	Perm = fibercasbin.New(fibercasbin.Config{
		Enforcer:      casbin.Enforcer,
		ModelFilePath: config.Conf.CasbinModelFilePath,
		PolicyAdapter: casbin.Adapter,
		Lookup: func(c *fiber.Ctx) string {
			return service.Auth(c).JWTUser().Username
		},
		LookupDom: func(c *fiber.Ctx) string {
			return c.Locals("domainId").(string)
		},
	})
}
