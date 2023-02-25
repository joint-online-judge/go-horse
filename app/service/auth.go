package service

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/model"
	"github.com/joint-online-judge/go-horse/app/query"
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

func (s *authImpl) JWTUser() *schema.User {
	claims := s.c.Locals("jwt").(*schema.JWTClaims)
	user := schema.User{
		Gravatar: &claims.Gravatar,
		ID:       uuid.MustParse(claims.ID),
		IsActive: &claims.IsActive,
		Role:     &claims.Role,
		Username: claims.Username,
	}
	return &user
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
				ID:        user.ID.String(),
				Subject:   user.ID.String(),
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

func (s *authImpl) NewRefreshToken(user schema.User, oauth_name string) (string, error) {
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
			ID:        user.ID.String(),
			Subject:   user.ID.String(),
			// Audience:  []string{"somebody_else"},
		},
	})
	return token.SignedString([]byte(config.Conf.JwtSecret))
}

func (s *authImpl) NewAuthTokens(
	user schema.User,
	oauth_name string,
	fresh bool,
) (*schema.AuthTokens, error) {
	category := ""
	if oauth_name != "" {
		category = "oauth"
	}
	accessToken, err := s.NewAccessToken(user, oauth_name, category, fresh)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.NewRefreshToken(user, oauth_name)
	if err != nil {
		return nil, err
	}
	return &schema.AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "bearer",
	}, err
}

func (s *authImpl) CreateAuthTokens(
	user *schema.User,
	oauthName string,
	fresh bool,
) (*schema.AuthTokens, error) {
	token, err := s.NewAuthTokens(*user, oauthName, fresh)
	if err != nil {
		return nil, err
	}
	cookie := new(fiber.Cookie)
	cookie.Name = "refresh_token_cookie"
	cookie.Value = token.RefreshToken
	cookie.Expires = time.Now().Add(time.Duration(config.Conf.JwtRefreshExpireSeconds) * time.Second)
	s.c.Cookie(cookie)
	return token, nil
}

func (s *authImpl) RegisterNewUser(userCreate *schema.UserCreate) (*schema.User, error) {
	hashedPassword, err := schema.HashPassword(*userCreate.Password)
	if err != nil {
		return nil, err
	}
	ip := s.c.IP()
	userModel := model.User{
		Username:       *userCreate.Username,
		Email:          *userCreate.Email,
		StudentID:      "",
		RealName:       "",
		IsActive:       false,
		HashedPassword: hashedPassword,
		RegisterIP:     ip,
		LoginIP:        ip,
	}
	user, err := query.CreateObj[schema.User](&userModel)
	return &user, err
}
