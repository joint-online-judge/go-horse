package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/envvar"
	"github.com/gofiber/fiber/v2/middleware/expvar"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/redirect/v2"

	"github.com/joint-online-judge/go-horse/pkg/configs"
)

func Initalize(router *fiber.App) {
	if configs.Conf.Debug {
		router.Use(redirect.New(redirect.Config{
			Rules: map[string]string{
				"/":       "/api/v1",
				"/api/v1": "/api/v1/docs",
			},
			StatusCode: 301,
		}))
		router.Use(pprof.New())                     // /debug/pprof
		router.Use(expvar.New())                    // /debug/vars
		router.Use("/debug/envvars", envvar.New())  // /debug/envvars
		router.Get("/debug/metrics", monitor.New()) // /debug/metrics
	}
	router.Use(logger.New(logger.Config{
		// Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n",
		Format:     "${time} [${ip}:${port}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:03.000",
	}))
	router.Use(cors.New(cors.Config{
		AllowOrigins: "*", // comma string format
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	router.Use(Recover)
	if !configs.Conf.Debug {
		router.Use(Security)
		router.Use(csrf.New())
	}
	router.Use(Json)
}
