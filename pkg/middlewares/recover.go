package middlewares

import (
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rollbar/rollbar-go"
	"github.com/sirupsen/logrus"
)

// func getCallers(skip int) (pc []uintptr) {
// 	pc = make([]uintptr, 1000)
// 	i := runtime.Callers(skip+1, pc)
// 	return pc[0:i]
// }

var recoverHandler = recover.New(recover.Config{
	Next:             nil,
	EnableStackTrace: true,
	StackTraceHandler: func(ctx *fiber.Ctx, err interface{}) {
		logrus.Errorf("recover middleware panic: %v\n%s\n", err, debug.Stack())
		ctx.Locals("rollbar_reported", 1)
		rollbar.Critical(
			errors.New(fmt.Sprint(err)),
			// getCallers(3),
			map[string]interface{}{
				"endpoint": ctx.Request().URI().String(),
			},
		)
	},
})

func Recover(c *fiber.Ctx) error {
	return recoverHandler(c)
}
