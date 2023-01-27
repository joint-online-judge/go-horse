package v1

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/models"
	"github.com/joint-online-judge/go-horse/app/querys"
	"github.com/joint-online-judge/go-horse/app/schemas"
	"github.com/sirupsen/logrus"
)

// Login
// (POST /auth/login)
func (s *Api) Login(
	c *fiber.Ctx,
	request schemas.LoginRequestObject,
) (any, error) {
	userModel := models.User{Username: *request.Body.Username}
	user, err := querys.GetObj[models.User, schemas.User](&userModel)
	if err != nil {
		return nil, schemas.NewBizError(schemas.UserNotFoundError)
	}
	if !schemas.VerifyPassword(
		*request.Body.Password,
		userModel.HashedPassword,
	) {
		return nil, schemas.NewBizError(
			schemas.UsernamePasswordError,
			"incorrect password",
		)
	}
	logrus.Infof("user login: %+v", user)
	userModel.ID = user.Id
	userModel.LoginAt = time.Now()
	userModel.LoginIP = c.Context().RemoteAddr().String()
	err = querys.SaveObj(&userModel)
	if err != nil {
		return nil, err
	}
	accessToken, err := schemas.NewAccessToken(user, "", "user", true)
	if err != nil {
		return nil, err
	}
	refreshToken, err := schemas.NewRefreshToken(user, "")
	if err != nil {
		return nil, err
	}
	return schemas.AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "bearer",
	}, err
}

// Logout
// (POST /auth/logout)
func (s *Api) Logout(
	c *fiber.Ctx,
	request schemas.LogoutRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// List Oauth2
// (GET /auth/oauth2)
func (s *Api) ListOauth2(
	c *fiber.Ctx,
	request schemas.ListOauth2RequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Oauth Authorize
// (GET /auth/oauth2/{oauth2}/authorize)
func (s *Api) OauthAuthorize(
	c *fiber.Ctx,
	request schemas.OauthAuthorizeRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Refresh
// (POST /auth/refresh)
func (s *Api) Refresh(
	c *fiber.Ctx,
	request schemas.RefreshRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Register
// (POST /auth/register)
func (s *Api) Register(
	c *fiber.Ctx,
	request schemas.RegisterRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}

// Get Token
// (GET /auth/token)
func (s *Api) GetToken(
	c *fiber.Ctx,
	request schemas.GetTokenRequestObject,
) (any, error) {
	return nil, schemas.NewBizError(schemas.APINotImplementedError)
}