package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/db"
	"github.com/joint-online-judge/go-horse/middlewares"
	"github.com/joint-online-judge/go-horse/schemas"
)

// Login
// (POST /auth/login)
func (s *ApiV1) Login(c *fiber.Ctx, request schemas.LoginRequestObject) (any, error) {
	user, _ := db.GetUser()
	token, err := middlewares.NewAccessToken(user, "user", "", true)
	return fiber.Map{"token": token}, err
}

// Logout
// (POST /auth/logout)
func (s *ApiV1) Logout(c *fiber.Ctx, request schemas.LogoutRequestObject) (any, error) {
	return nil, nil
}

// List Oauth2
// (GET /auth/oauth2)
func (s *ApiV1) ListOauth2(
	c *fiber.Ctx,
	request schemas.ListOauth2RequestObject,
) (any, error) {
	return nil, nil
}

// Oauth Authorize
// (GET /auth/oauth2/{oauth2}/authorize)
func (s *ApiV1) OauthAuthorize(
	c *fiber.Ctx,
	request schemas.OauthAuthorizeRequestObject,
) (any, error) {
	return nil, nil
}

// Refresh
// (POST /auth/refresh)
func (s *ApiV1) Refresh(c *fiber.Ctx, request schemas.RefreshRequestObject) (any, error) {
	return nil, nil
}

// Register
// (POST /auth/register)
func (s *ApiV1) Register(
	c *fiber.Ctx,
	request schemas.RegisterRequestObject,
) (any, error) {
	return nil, nil
}

// Get Token
// (GET /auth/token)
func (s *ApiV1) GetToken(
	c *fiber.Ctx,
	request schemas.GetTokenRequestObject,
) (any, error) {
	return nil, nil
}
