package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/config"
	"github.com/joint-online-judge/go-horse/schemas"
)

func JWT() fiber.Handler {
	return jwtware.New(jwtware.Config{
		// SigningMethod: config.Config.JwtAlgorithm,
		SigningMethod: "HS256",
		SigningKey:    []byte(config.Config.JwtSecret),
	})
}

type JWTClaims struct {
	Type      string `json:"type"`
	Fresh     bool   `json:"fresh"`
	Csrf      string `json:"csrf"`
	Category  string `json:"category"`
	Username  string `json:"username"`
	Gravatar  string `json:"gravatar"`
	Role      string `json:"role"`
	IsActive  bool   `json:"isActive"`
	OauthName string `json:"oauthName"`
	jwt.RegisteredClaims
}

func NewAccessToken(user schemas.User, category, oauth_name string, fresh bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		Type:      "access",
		Fresh:     fresh,
		Csrf:      "test", // FIXME: do we need it as we can use csrf middleware?
		Category:  category,
		Username:  user.Username,
		Gravatar:  *user.Gravatar,
		Role:      *user.Role,
		IsActive:  *user.IsActive,
		OauthName: oauth_name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(time.Duration(config.Config.JwtExpireSeconds) * time.Second),
			),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    uuid.New().String(),
			ID:        user.Id.String(),
			Subject:   user.Id.String(),
			// Audience:  []string{"somebody_else"},
		},
	})
	return token.SignedString([]byte(config.Config.JwtSecret))
}
