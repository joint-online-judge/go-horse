package router

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/joint-online-judge/go-horse/docs/horse"
)

func RegisterHorseStatic(router fiber.Router) {
	static := router.Group("/horse")
	static.Use("", filesystem.New(filesystem.Config{
		Root: http.FS(horse.Static),
	}))
}
