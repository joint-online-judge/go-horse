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
	userModel, err := service.User(c).GetCurrentUser()
	if err != nil {
		return nil, schema.NewBizError(schema.UserNotFoundError)
	}
	return convert.To[schema.UserDetail](userModel)
}

// Update Current User
// (PATCH /users/me)
func (s *Api) UpdateCurrentUser(
	c *fiber.Ctx,
	request schema.UpdateCurrentUserRequestObject,
) (any, error) {
	user, err := service.User(c).UpdateCurrentUser(*request.Body)
	if err != nil {
		return nil, err
	}
	return convert.To[schema.UserDetail](user)
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
	userModel, err := service.User(c).GetUser(request.Uid)
	if err != nil {
		return nil, schema.NewBizError(schema.UserNotFoundError)
	}
	return convert.To[schema.UserPreview](userModel)
}
