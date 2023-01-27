package middlewares

import (
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joint-online-judge/go-horse/pkg/logger"

	"github.com/rollbar/rollbar-go"
)

// func getCallers(skip int) (pc []uintptr) {
// 	pc = make([]uintptr, 1000)
// 	i := runtime.Callers(skip+1, pc)
// 	return pc[0:i]
// }

var recoverHandler = recover.New(recover.Config{
	Next:             nil,
	EnableStackTrace: true,
	StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
		logger.Errorf("panic: %v\n%s\n", e, debug.Stack())
		rollbar.Critical(
			errors.New(fmt.Sprint(e)),
			// getCallers(3),
			map[string]interface{}{
				"endpoint": c.Request().URI().String(),
			},
		)
	},
})

func Recover(c *fiber.Ctx) error {
	return recoverHandler(c)
}
