package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/handler"
	v1 "github.com/joint-online-judge/go-horse/app/handler/v1"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/pkg/middleware"
)

func Register(router *fiber.App) {
	middleware.Register(router)
	wrapper := schema.ServerInterfaceWrapper{
		Handler: schema.NewStrictHandler(
			&v1.Api{},
			[]schema.StrictMiddlewareFunc{handler.ValidateRequest},
			handler.ResponseHandler,
		),
	}
	api := router.Group("/api") // /api
	v1 := api.Group("/v1")      // /api/v1
	RegisterHorseStatic(v1)
	RegisterSwagger(v1)
	RegisterMisc(v1, wrapper)
	RegisterAuth(v1, wrapper)
	v1.Use(middleware.JWT())
	RegisterDomain(v1, wrapper)
	RegisterProblemSet(v1, wrapper)
	RegisterProblem(v1, wrapper)
	RegisterProblemConfig(v1, wrapper)
	RegisterProblemGroup(v1, wrapper)
	RegisterRecord(v1, wrapper)
	RegisterUser(v1, wrapper)
	RegisterAdmin(v1, wrapper)
	RegisterJudge(v1, wrapper)
}
