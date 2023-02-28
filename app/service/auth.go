package service

import (
	"time"

	"github.com/joint-online-judge/go-horse/pkg/convert"
	"github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/pkg/config"
)

type authImpl struct {
	c *fiber.Ctx
}

func Auth(c *fiber.Ctx) *authImpl {
	return &authImpl{
		c: c,
	}
}

func (s *authImpl) JWT() *schema.JWTClaims {
	return s.c.Locals("jwt").(*schema.JWTClaims)
}

func (s *authImpl) JWTUser() schema.User {
	claims := s.c.Locals("jwt").(*schema.JWTClaims)
	user := schema.User{
		Gravatar: &claims.Gravatar,
		Id:       uuid.MustParse(claims.ID),
		IsActive: &claims.IsActive,
		Role:     &claims.Role,
		Username: claims.Username,
	}
	return user
}

func (s *authImpl) NewAccessToken(
	user schema.User,
	oauth_name, category string,
	fresh bool,
) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, schema.JWTClaims{
		JWTCommon: schema.JWTCommon{
			Type:      "access",
			OauthName: oauth_name,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(
					time.Now().
						Add(time.Duration(config.Conf.JwtExpireSeconds) * time.Second),
				),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				Issuer:    uuid.New().String(),
				ID:        user.Id.String(),
				Subject:   user.Id.String(),
				// Audience:  []string{"somebody_else"},
			},
		},
		Fresh:    fresh,
		Csrf:     "test", // FIXME: do we need it as we can use csrf middleware?
		Category: category,
		Username: user.Username,
		Gravatar: *user.Gravatar,
		Role:     *user.Role,
		IsActive: *user.IsActive,
	})
	return token.SignedString([]byte(config.Conf.JwtSecret))
}

func (s *authImpl) NewRefreshToken(
	user schema.User,
	oauth_name string,
) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, schema.JWTCommon{
		Type:      "refresh",
		OauthName: oauth_name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().
					Add(time.Duration(config.Conf.JwtExpireSeconds) * time.Second),
			),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    uuid.New().String(),
			ID:        user.Id.String(),
			Subject:   user.Id.String(),
			// Audience:  []string{"somebody_else"},
		},
	})
	return token.SignedString([]byte(config.Conf.JwtSecret))
}

func (s *authImpl) NewAuthTokens(
	user schema.User,
	oauth_name string,
	fresh bool,
) (authTokens schema.AuthTokens, err error) {
	category := ""
	if oauth_name != "" {
		category = "oauth"
	}
	accessToken, err := s.NewAccessToken(user, oauth_name, category, fresh)
	if err != nil {
		return
	}
	refreshToken, err := s.NewRefreshToken(user, oauth_name)
	if err != nil {
		return
	}
	return schema.AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "bearer",
	}, err
}

func (s *authImpl) CreateAuthTokens(
	user schema.User,
	oauthName string,
	fresh bool,
) (authTokens schema.AuthTokens, err error) {
	authTokens, err = s.NewAuthTokens(user, oauthName, fresh)
	if err != nil {
		return
	}
	// Set cookie: access token
	s.c.Cookie(&fiber.Cookie{
		Name:  "access_token_cookie",
		Value: authTokens.AccessToken,
		Expires: time.Now().
			Add(time.Duration(config.Conf.JwtExpireSeconds) * time.Second),
		HTTPOnly: true,
		SameSite: "lax",
	})
	// Set cookie: refresh token
	s.c.Cookie(&fiber.Cookie{
		Name:  "refresh_token_cookie",
		Value: authTokens.RefreshToken,
		Expires: time.Now().
			Add(time.Duration(config.Conf.JwtRefreshExpireSeconds) * time.Second),
		HTTPOnly: true,
		SameSite: "lax",
	})
	return
}

func (s *authImpl) RegisterNewUser(userCreate *schema.UserCreate) (
	user schema.User, err error,
) {
	hashedPassword, err := schema.HashPassword(*userCreate.Password)
	if err != nil {
		return
	}
	ip := s.c.IP()
	userModel := model.User{
		Username:       *userCreate.Username,
		Email:          *userCreate.Email,
		StudentId:      "",
		RealName:       "",
		IsActive:       false,
		HashedPassword: hashedPassword,
		RegisterIP:     ip,
		LoginIP:        ip,
	}
	err = db.Create(&userModel).Error
	if err != nil {
		return
	}
	return convert.To[schema.User](userModel)
}

func (s *authImpl) Login(loginForm *schema.OAuth2PasswordRequestForm) (
	authTokens schema.AuthTokens, err error,
) {
	userModel := model.User{Username: *loginForm.Username}
	err = db.First(&userModel).Error
	if err != nil {
		err = schema.NewBizError(schema.UserNotFoundError)
		return
	}
	if !schema.VerifyPassword(
		*loginForm.Password,
		userModel.HashedPassword,
	) {
		err = schema.NewBizError(
			schema.UsernamePasswordError,
			"incorrect password",
		)
		return
	}
	logrus.Debugf("user login: %+v", userModel)
	userModel.LoginAt = time.Now()
	userModel.LoginIP = s.c.IP()
	if err = db.Save(&userModel).Error; err != nil {
		return
	}
	user, err := convert.To[schema.User](userModel)
	if err != nil {
		return
	}
	return Auth(s.c).CreateAuthTokens(user, "", true)
}

func (s *authImpl) Logout() {
	s.c.Cookie(&fiber.Cookie{
		Name: "access_token_cookie",
		// Set expiry date to the past
		Expires:  time.Now().Add(-(time.Hour * 2)),
		HTTPOnly: true,
		SameSite: "lax",
	})
	s.c.Cookie(&fiber.Cookie{
		Name: "refresh_token_cookie",
		// Set expiry date to the past
		Expires:  time.Now().Add(-(time.Hour * 2)),
		HTTPOnly: true,
		SameSite: "lax",
	})
}
