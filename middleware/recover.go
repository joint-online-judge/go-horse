package middleware

import (
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	log "github.com/sirupsen/logrus"
)

var recoverHandler = recover.New(recover.Config{
	Next:             nil,
	EnableStackTrace: true,
	StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
		// TODO: set the formatter elsewhere
		log.SetFormatter(&log.TextFormatter{
			DisableQuote: true,
		})
		log.Errorf("panic: %v\n%s\n", e, debug.Stack())
	},
})

func Recover(c *fiber.Ctx) error {
	return recoverHandler(c)
}
