package main

import (
	"log"
	"os"

	"github.com/joint-online-judge/go-horse/database"
	"github.com/joint-online-judge/go-horse/handlers"
	"github.com/joint-online-judge/go-horse/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	godotenv.Load()
	app := fiber.New(fiber.Config{
		ErrorHandler: handlers.ErrorHandler,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // comma string format
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	database.ConnectDB()

	router.Initalize(app)
	log.Fatal(app.Listen(":" + getenv("PORT", "3000")))
}

/*
ENV Variables:
will auto set to 3000 if not set
PORT=3000
this should be a connection string or url
DATABASE_URL="host=localhost port=5432 user=postgres password= dbname= sslmode=disable"
**
Docker Command for Postgres database:
docker run --name database -d -p 5432:5432 -e POSTGRES_PASSWORD=password postgres:alpine

DB_URL Variable for docker database
DATABASE_URL="host=localhost port=5432 user=postgres password=password dbname=postgres sslmode=disable"
**
Docker build base image in first stage for development
docker build --target build -t base .
**
run dev container
docker run -p 3000:3000 --mount type=bind,source=/root/go/src/fiber,target=/go/src/app --name fiber -td base
**
rebuild and run package
docker exec -it web go run main.go
**
stop and remove container
docker stop fiber; docker rm fiber
*/
