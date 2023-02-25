package v1

import (
	"time"

	"github.com/joint-online-judge/go-horse/app/service"

	"github.com/gofiber/fiber/v2"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/query"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/pkg/convert"
	"github.com/sirupsen/logrus"
)

// Login
// (POST /auth/login)
func (s *Api) Login(
	c *fiber.Ctx,
	request schema.LoginRequestObject,
) (any, error) {
	userModel, err := query.GetObj[model.User](
		&model.User{Username: *request.Body.Username},
	)
	if err != nil {
		return nil, schema.NewBizError(schema.UserNotFoundError)
	}
	if !schema.VerifyPassword(
		*request.Body.Password,
		userModel.HashedPassword,
	) {
		return nil, schema.NewBizError(
			schema.UsernamePasswordError,
			"incorrect password",
		)
	}
	logrus.Debugf("user login: %+v", userModel)
	userModel.LoginAt = time.Now()
	userModel.LoginIP = c.IP()
	if err = query.SaveObj(&userModel); err != nil {
		return nil, err
	}
	user, err := convert.To[schema.User](&userModel)
	if err != nil {
		return nil, err
	}
	return service.Auth(c).NewAuthTokens(user, "", true)
}

// Logout
// (POST /auth/logout)
func (s *Api) Logout(
	c *fiber.Ctx,
	request schema.LogoutRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// List Oauth2
// (GET /auth/oauth2)
func (s *Api) ListOauth2(
	c *fiber.Ctx,
	request schema.ListOauth2RequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Oauth Authorize
// (GET /auth/oauth2/{oauth2}/authorize)
func (s *Api) OauthAuthorize(
	c *fiber.Ctx,
	request schema.OauthAuthorizeRequestObject,
) (any, error) {
	return nil, schema.NewBizError(schema.APINotImplementedError)
}

// Refresh
// (POST /auth/refresh)
func (s *Api) Refresh(
	c *fiber.Ctx,
	request schema.RefreshRequestObject,
) (any, error) {
	user := service.Auth(c).JWTUser()
	return service.Auth(c).NewAuthTokens(*user, "", false)
}

// Register
// (POST /auth/register)
func (s *Api) Register(
	c *fiber.Ctx,
	request schema.RegisterRequestObject,
) (any, error) {
	userCreate := request.Body
	if *userCreate.Password == "" {
		return nil, schema.NewBizError(
			schema.UserRegisterError,
			"password not provided",
		)
	} else if *userCreate.Username == "" {
		return nil, schema.NewBizError(
			schema.UserRegisterError,
			"username not provided",
		)
	} else if *userCreate.Email == "" {
		return nil, schema.NewBizError(
			schema.UserRegisterError,
			"email not provided",
		)
	}
	user, err := service.Auth(c).RegisterNewUser(userCreate)
	if err != nil {
		return nil, err
	}
	logrus.Infof("user: %T + 1, %v", user, user)
	return service.Auth(c).CreateAuthTokens(user, "", true)
}

// Get Token
// (GET /auth/token)
func (s *Api) GetToken(
	c *fiber.Ctx,
	request schema.GetTokenRequestObject,
) (any, error) {
	user := service.Auth(c).JWTUser()
	return service.Auth(c).NewAuthTokens(*user, "", true)
}
