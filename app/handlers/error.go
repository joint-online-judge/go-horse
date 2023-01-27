package handlers

import (
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schemas"
	"github.com/joint-online-judge/go-horse/pkg/configs"
	"github.com/rollbar/rollbar-go"
	"github.com/sirupsen/logrus"
)

func Error(ctx *fiber.Ctx, err error) error {
	stack := debug.Stack()
	if n, ok := ctx.Locals("rollbar_reported").(int); !ok || n == 0 {
		logrus.Errorf("error handler panic: %v\n%s\n", err, stack)
		rollbar.Critical(
			errors.New(fmt.Sprint(err)),
			map[string]interface{}{
				"endpoint": ctx.Request().URI().String(),
			},
		)
	}
	errorMsg := ""
	if configs.Conf.Debug {
		errorMsg = string(stack)
	}
	return ctx.Status(fiber.StatusOK).
		JSON(schemas.NewEmptyResp(schemas.InternalServerError, errorMsg))
}
