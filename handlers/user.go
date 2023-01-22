package handlers

import (
	"context"

	"github.com/joint-online-judge/go-horse/schemas"
)

// Get Current User
// (GET /users/me)
func (s *ApiV1) GetCurrentUser(
	ctx context.Context,
	request schemas.GetCurrentUserRequestObject,
) (any, error) {
	return nil, nil
}

// Update Current User
// (PATCH /users/me)
func (s *ApiV1) UpdateCurrentUser(
	ctx context.Context,
	request schemas.UpdateCurrentUserRequestObject,
) (any, error) {
	return nil, nil
}

// Change Password
// (PATCH /users/me/password)
func (s *ApiV1) ChangePassword(
	ctx context.Context,
	request schemas.ChangePasswordRequestObject,
) (any, error) {
	return nil, nil
}

// Get User
// (GET /users/{uid})
func (s *ApiV1) GetUser(ctx context.Context, request schemas.GetUserRequestObject) (any, error) {
	return nil, nil
}
