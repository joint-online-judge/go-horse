package main

import (
	"fmt"

	"github.com/joint-online-judge/go-horse/pkg/configs"
	"github.com/joint-online-judge/go-horse/pkg/middlewares"
	"github.com/joint-online-judge/go-horse/pkg/routers"
	"github.com/joint-online-judge/go-horse/platform/db"

	_ "github.com/joint-online-judge/go-horse/docs" // load API Docs files (Swagger)

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// Swagger information
//
//	@title			API
//	@version		1.0
//	@description	This is an auto-generated API Docs.
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	API Support
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@BasePath		/api/v1
func main() {
	configs.Initalize()
	db.ConnectDB()
	app := fiber.New(fiber.Config{
		// ErrorHandler: utils.Panic,
		Prefork: !configs.Conf.Debug,
	})
	middlewares.Initalize(app)
	routers.Initalize(app)
	log.Fatal(
		app.Listen(fmt.Sprintf("%s:%d", configs.Conf.Host, configs.Conf.Port)),
	)
}
