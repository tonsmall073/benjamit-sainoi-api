package jwt

import (
	"bjm/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func UseUserGuard(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return utils.FiberResponseErrorJson(c, utils.HttpStatusCodes[401], 401)
	}
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return utils.FiberResponseErrorJson(c, utils.HttpStatusCodes[401], 401)
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	res, err := VerifyToken(token)
	if err != nil {
		return utils.FiberResponseErrorJson(c, utils.HttpStatusCodes[401], 401)
	}

	role := res["role"].(string)
	if role != "USER" && role != "ADMIN" {
		return utils.FiberResponseErrorJson(c, utils.HttpStatusCodes[403], 403)
	}

	return c.Next()
}

func UseAdminGuard(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return utils.FiberResponseErrorJson(c, utils.HttpStatusCodes[401], 401)
	}
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return utils.FiberResponseErrorJson(c, utils.HttpStatusCodes[401], 401)
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	res, err := VerifyToken(token)
	if err != nil {
		return utils.FiberResponseErrorJson(c, utils.HttpStatusCodes[401], 401)
	}

	if res["role"].(string) != "ADMIN" {
		return utils.FiberResponseErrorJson(c, utils.HttpStatusCodes[403], 403)
	}

	return c.Next()
}
