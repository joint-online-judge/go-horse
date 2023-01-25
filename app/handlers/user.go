package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/schemas"
)

// Get Current User
// (GET /users/me)
func (s *ApiV1) GetCurrentUser(
	c *fiber.Ctx,
	request schemas.GetCurrentUserRequestObject,
) (any, error) {
	return nil, nil
}

// Update Current User
// (PATCH /users/me)
func (s *ApiV1) UpdateCurrentUser(
	c *fiber.Ctx,
	request schemas.UpdateCurrentUserRequestObject,
) (any, error) {
	return nil, nil
}

// Change Password
// (PATCH /users/me/password)
func (s *ApiV1) ChangePassword(
	c *fiber.Ctx,
	request schemas.ChangePasswordRequestObject,
) (any, error) {
	return nil, nil
}

// Get User
// (GET /users/{uid})
func (s *ApiV1) GetUser(
	c *fiber.Ctx,
	request schemas.GetUserRequestObject,
) (any, error) {
	return nil, nil
}
