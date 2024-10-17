package jwt

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	secret := []byte(os.Getenv("SECRET_KEY_TOKEN_JWT"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
