package utils

import "github.com/gofiber/fiber/v2"

func FiberResponseJson(c *fiber.Ctx, body interface{}, status int) error {
	if status <= 0 {
		return c.Status(500).JSON(fiber.Map{
			"messageDesc": "",
			"statusCode":  HttpStatusCodes[500],
		}) // ส่ง response ด้วย status 500
	}

	return c.Status(status).JSON(body)
}
