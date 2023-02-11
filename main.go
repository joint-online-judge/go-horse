package main

import (
	"fmt"
	"runtime/debug"

	_ "github.com/joint-online-judge/go-horse/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/handler"
	_ "github.com/joint-online-judge/go-horse/docs" // load API Docs files (Swagger)
	"github.com/joint-online-judge/go-horse/pkg/config"
	"github.com/joint-online-judge/go-horse/pkg/json"
	"github.com/joint-online-judge/go-horse/pkg/router"
	"github.com/joint-online-judge/go-horse/platform"
	"github.com/sirupsen/logrus"
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
	platform.Bootstrap()
	app := fiber.New(fiber.Config{
		ErrorHandler: handler.Error,
		Prefork:      !config.Conf.Debug,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})
	router.Register(app)
	err := app.Listen(fmt.Sprintf("%s:%d", config.Conf.Host, config.Conf.Port))
	logrus.Fatalf("app.Listen error: %v\n%s\n", err, debug.Stack())
}
