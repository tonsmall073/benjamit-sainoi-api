package jwt

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func DecodeToken(c *fiber.Ctx) jwt.MapClaims {
	authHeader := c.Get("Authorization")
	getToken := strings.TrimPrefix(authHeader, "Bearer ")
	res, err := VerifyToken(getToken)
	if err != nil {
		log.Println("[ERROR] decode token : " + err.Error())
		return jwt.MapClaims{}
	}
	return res
}

func DecodeTokenByTokenStr(token string) (jwt.MapClaims, error) {
	getToken := strings.TrimPrefix(token, "Bearer ")
	res, err := VerifyToken(getToken)
	if err != nil {
		return jwt.MapClaims{}, err
	}
	return res, nil
}
