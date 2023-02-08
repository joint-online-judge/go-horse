package main

import (
	"fmt"

	"github.com/joint-online-judge/go-horse/app/handlers"
	"github.com/joint-online-judge/go-horse/pkg/configs"
	"github.com/joint-online-judge/go-horse/pkg/routers"
	"github.com/joint-online-judge/go-horse/platform"
	"github.com/sirupsen/logrus"

	_ "github.com/joint-online-judge/go-horse/docs" // load API Docs files (Swagger)

	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/pkg/json"
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
		ErrorHandler: handlers.Error,
		Prefork:      !configs.Conf.Debug,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})
	routers.Register(app)
	logrus.Fatal(
		app.Listen(fmt.Sprintf("%s:%d", configs.Conf.Host, configs.Conf.Port)),
	)
}
