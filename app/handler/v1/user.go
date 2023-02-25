package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/app/service"
)

// Get Current User
// (GET /users/me)
func (s *Api) GetCurrentUser(
	c *fiber.Ctx,
	request schema.GetCurrentUserRequestObject,
) (any, error) {
	return service.User(c).GetCurrentUser()
}

// Update Current User
// (PATCH /users/me)
func (s *Api) UpdateCurrentUser(
	c *fiber.Ctx,
	request schema.UpdateCurrentUserRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Change Password
// (PATCH /users/me/password)
func (s *Api) ChangePassword(
	c *fiber.Ctx,
	request schema.ChangePasswordRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Get User
// (GET /users/{uid})
func (s *Api) GetUser(
	c *fiber.Ctx,
	request schema.GetUserRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}
