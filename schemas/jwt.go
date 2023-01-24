package schemas

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

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

func JWT(c *fiber.Ctx) *JWTClaims {
	return c.Locals("jwt").(*JWTClaims)
}
