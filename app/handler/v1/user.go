package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/app/service"
	"github.com/joint-online-judge/go-horse/pkg/convert"
)

// Get Current User
// (GET /users/me)
func (s *Api) GetCurrentUser(
	c *fiber.Ctx,
	request schema.GetCurrentUserRequestObject,
) (any, error) {
	return convert.WithErr[schema.UserDetail](
		service.User(c).GetCurrentUser(),
	)
}

// Update Current User
// (PATCH /users/me)
func (s *Api) UpdateCurrentUser(
	c *fiber.Ctx,
	request schema.UpdateCurrentUserRequestObject,
) (any, error) {
	return convert.WithErr[schema.UserDetail](
		service.User(c).UpdateCurrentUser(*request.Body),
	)
}

// Change Password
// (PATCH /users/me/password)
func (s *Api) ChangePassword(
	c *fiber.Ctx,
	request schema.ChangePasswordRequestObject,
) (any, error) {
	return convert.WithErr[schema.UserDetail](
		service.User(c).ChangePassword(*request.Body),
	)
}

// Get User
// (GET /users/{uid})
func (s *Api) GetUser(
	c *fiber.Ctx,
	request schema.GetUserRequestObject,
) (any, error) {
	return convert.WithErr[schema.UserPreview](
		service.User(c).GetUser(request.Uid),
	)
}
