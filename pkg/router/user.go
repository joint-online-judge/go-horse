package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schema"
)

func RegisterUser(router fiber.Router, wrapper schema.ServerInterfaceWrapper) {
	users := router.Group("/users")
	me := users.Group("/me")
	user := users.Group("/:uid")
	me.Get("", wrapper.GetCurrentUser)
	me.Patch("", wrapper.UpdateCurrentUser)
	me.Patch("/password", wrapper.ChangePassword)
	user.Get("", wrapper.GetUser)
}
