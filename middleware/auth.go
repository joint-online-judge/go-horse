package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/handlers"
	"github.com/joint-online-judge/go-horse/model"
)

func Authenticated(c *fiber.Ctx) error {
	json := new(model.Session)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid Session Format",
		})
	}
	user, err := handlers.GetUser(json.Sessionid)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "404: not found",
		})
	}
	c.Locals("user", user)
	return c.Next()
}
