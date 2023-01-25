package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/models"
	"github.com/joint-online-judge/go-horse/app/querys"
	"github.com/joint-online-judge/go-horse/app/schemas"
	"github.com/joint-online-judge/go-horse/pkg/middlewares"
	log "github.com/sirupsen/logrus"
)

// Login
// (POST /auth/login)
func (s *ApiV1) Login(
	c *fiber.Ctx,
	request schemas.LoginRequestObject,
) (any, error) {
	userModel := models.User{Username: *request.Body.Username}
	user, err := querys.GetObj[models.User, schemas.User](&userModel)
	if err != nil {
		return nil, schemas.NewBizError(schemas.UserNotFoundError)
	}
	if !user.VerifyPassword(*request.Body.Password) {
		return nil, schemas.NewBizError(
			schemas.UsernamePasswordError,
			"incorrect password",
		)
	}
	log.Infof("user login: %+v", user)
	userModel.ID = user.Id.String()
	userModel.LoginAt = time.Now()
	userModel.LoginIP = c.Context().RemoteAddr().String()
	err = querys.SaveObj(&userModel)
	if err != nil {
		return nil, err
	}
	accessToken, err := middlewares.NewAccessToken(user, "user", "", true)
	// TODO: refresh token
	return schemas.AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: accessToken,
		TokenType:    "bearer",
	}, err
}

// Logout
// (POST /auth/logout)
func (s *ApiV1) Logout(
	c *fiber.Ctx,
	request schemas.LogoutRequestObject,
) (any, error) {
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
func (s *ApiV1) Refresh(
	c *fiber.Ctx,
	request schemas.RefreshRequestObject,
) (any, error) {
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
