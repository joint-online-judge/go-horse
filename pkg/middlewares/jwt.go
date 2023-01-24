package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/app/schemas"
	"github.com/joint-online-judge/go-horse/pkg/configs"
)

func JWT() fiber.Handler {
	return jwtware.New(jwtware.Config{
		// SigningMethod: config.Config.JwtAlgorithm,
		SigningMethod: "HS256",
		SigningKey:    []byte(configs.Conf.JwtSecret),
		SuccessHandler: func(c *fiber.Ctx) error {
			tokenString := c.Locals("user").(*jwt.Token).Raw
			token, err := jwt.ParseWithClaims(
				tokenString,
				&schemas.JWTClaims{},
				func(t *jwt.Token) (interface{}, error) { return []byte(configs.Conf.JwtSecret), nil },
			)
			if err != nil {
				return err
			}
			claims := token.Claims.(*schemas.JWTClaims)
			c.Locals("jwt", claims)
			return c.Next()
		},
	})
}

func NewAccessToken(user schemas.User, category, oauth_name string, fresh bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, schemas.JWTClaims{
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
				time.Now().Add(time.Duration(configs.Conf.JwtExpireSeconds) * time.Second),
			),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    uuid.New().String(),
			ID:        user.Id.String(),
			Subject:   user.Id.String(),
			// Audience:  []string{"somebody_else"},
		},
	})
	return token.SignedString([]byte(configs.Conf.JwtSecret))
}
