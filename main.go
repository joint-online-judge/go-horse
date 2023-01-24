package main

import (
	"fmt"

	"github.com/joint-online-judge/go-horse/config"
	"github.com/joint-online-judge/go-horse/db"
	"github.com/joint-online-judge/go-horse/middlewares"
	"github.com/joint-online-judge/go-horse/routers"
	"github.com/rollbar/rollbar-go"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.Initalize()
	db.ConnectDB()
	rollbar.SetToken(config.Config.RollbarAccessToken)
	app := fiber.New(fiber.Config{
		// ErrorHandler: utils.Panic,
		Prefork: !config.Config.Debug,
	})
	middlewares.Initalize(app)
	routers.Initalize(app)
	log.Fatal(app.Listen(fmt.Sprintf("%s:%d", config.Config.Host, config.Config.Port)))
}
