package jwt

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(username string, uuid string, role string) (string, error) {
	getTime, getTimeErr := strconv.Atoi(os.Getenv("EXPIRE_HOUR_TOKEN_JWT"))
	var expire time.Duration = time.Hour * 72
	if getTimeErr == nil {
		expire = time.Hour * time.Duration(getTime)
	}
	claims := jwt.MapClaims{
		"username": username,
		"uuid":     uuid,
		"role":     role,
		"exp":      time.Now().Add(expire).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	res, err := token.SignedString([]byte(os.Getenv("SECRET_KEY_TOKEN_JWT")))

	return res, err
}
