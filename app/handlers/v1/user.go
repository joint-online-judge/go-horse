package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schemas"
)

// Get Current User
// (GET /users/me)
func (s *Api) GetCurrentUser(
	c *fiber.Ctx,
	request schemas.GetCurrentUserRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Update Current User
// (PATCH /users/me)
func (s *Api) UpdateCurrentUser(
	c *fiber.Ctx,
	request schemas.UpdateCurrentUserRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Change Password
// (PATCH /users/me/password)
func (s *Api) ChangePassword(
	c *fiber.Ctx,
	request schemas.ChangePasswordRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Get User
// (GET /users/{uid})
func (s *Api) GetUser(
	c *fiber.Ctx,
	request schemas.GetUserRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}
