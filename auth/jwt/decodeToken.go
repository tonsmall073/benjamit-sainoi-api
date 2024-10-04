package jwt

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func DecodeToken(c *fiber.Ctx) jwt.MapClaims {
	authHeader := c.Get("Authorization")
	getToken := strings.TrimPrefix(authHeader, "Bearer ")
	res, _ := VerifyToken(getToken)
	return res
}
