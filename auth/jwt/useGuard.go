package jwt

import (
	"bjm/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func UseGuard(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{
			"statusCode":  401,
			"messageDesc": utils.HttpStatusCodes[401],
		})
	}

	// ตรวจสอบว่า authHeader เริ่มต้นด้วย "Bearer "
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(401).JSON(fiber.Map{
			"statusCode":  401,
			"messageDesc": utils.HttpStatusCodes[401],
		})
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	_, err := VerifyToken(token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"statusCode":  401,
			"messageDesc": utils.HttpStatusCodes[401],
		})
	}

	return c.Next()
}

func UseAdminGuard(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{
			"statusCode":  401,
			"messageDesc": utils.HttpStatusCodes[401],
		})
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(401).JSON(fiber.Map{
			"statusCode":  401,
			"messageDesc": utils.HttpStatusCodes[401],
		})
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	res, err := VerifyToken(token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"statusCode":  401,
			"messageDesc": utils.HttpStatusCodes[401],
		})
	}
	if res["role"].(string) != "ADMIN" {
		return c.Status(403).JSON(fiber.Map{
			"statusCode":  403,
			"messageDesc": utils.HttpStatusCodes[403],
		})
	}

	return c.Next()
}
