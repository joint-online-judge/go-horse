package main

import (
	"fmt"

	"github.com/joint-online-judge/go-horse/pkg/configs"
	"github.com/joint-online-judge/go-horse/pkg/middlewares"
	"github.com/joint-online-judge/go-horse/pkg/routers"
	"github.com/joint-online-judge/go-horse/platform/db"
	"github.com/rollbar/rollbar-go"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func main() {
	configs.Initalize()
	db.ConnectDB()
	rollbar.SetToken(configs.Conf.RollbarAccessToken)
	app := fiber.New(fiber.Config{
		// ErrorHandler: utils.Panic,
		Prefork: !configs.Conf.Debug,
	})
	middlewares.Initalize(app)
	routers.Initalize(app)
	log.Fatal(app.Listen(fmt.Sprintf("%s:%d", configs.Conf.Host, configs.Conf.Port)))
}
