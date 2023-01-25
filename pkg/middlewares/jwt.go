package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
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
