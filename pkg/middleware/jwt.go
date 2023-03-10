package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joint-online-judge/go-horse/app/schema"
	"github.com/joint-online-judge/go-horse/pkg/config"
)

func JWT() fiber.Handler {
	return jwtware.New(jwtware.Config{
		// SigningMethod: config.Config.JwtAlgorithm,
		SigningMethod: "HS256",
		SigningKey:    []byte(config.Conf.JwtSecret),
		TokenLookup:   "header:Authorization,cookie:access_token_cookie",
		SuccessHandler: func(c *fiber.Ctx) error {
			tokenString := c.Locals("user").(*jwt.Token).Raw
			token, err := jwt.ParseWithClaims(
				tokenString,
				&schema.JWTClaims{},
				func(t *jwt.Token) (any, error) { return []byte(config.Conf.JwtSecret), nil },
			)
			if err != nil {
				return err
			}
			claims := token.Claims.(*schema.JWTClaims)
			c.Locals("jwt", claims)
			return c.Next()
		},
	})
}
