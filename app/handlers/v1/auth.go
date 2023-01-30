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
	if err = querys.SaveObj(&userModel); err != nil {
		return nil, err
	}
	return schemas.NewAuthTokens(user, "")
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
	userCreate := request.Body
	if *userCreate.Password == "" {
		return nil, schemas.NewBizError(
			schemas.UserRegisterError,
			"password not provided",
		)
	} else if *userCreate.Username == "" {
		return nil, schemas.NewBizError(
			schemas.UserRegisterError,
			"username not provided",
		)
	} else if *userCreate.Email == "" {
		return nil, schemas.NewBizError(
			schemas.UserRegisterError,
			"email not provided",
		)
	}
	hashedPassword, err := schemas.HashPassword(*userCreate.Password)
	if err != nil {
		return nil, err
	}
	ip := c.Context().RemoteAddr().String()
	userModel := models.User{
		Username:       *userCreate.Username,
		Email:          *userCreate.Email,
		StudentID:      "",
		RealName:       "",
		IsActive:       false,
		HashedPassword: hashedPassword,
		RegisterIP:     ip,
		LoginIP:        ip,
	}
	user, err := querys.CreateObj[schemas.User](userModel)
	if err != nil {
		return nil, err
	}
	return schemas.NewAuthTokens(user, "")
}

// Get Token
// (GET /auth/token)
func (s *Api) GetToken(
	c *fiber.Ctx,
	request schemas.GetTokenRequestObject,
) (any, error) {
	user := schemas.JWTUser(c)
	return schemas.NewAuthTokens(*user, "")
}
