package utils

import "github.com/gofiber/fiber/v2"

func FiberResponseJson(c *fiber.Ctx, body interface{}, statusCode int) error {
	if statusCode <= 0 {
		return c.Status(500).JSON(fiber.Map{
			"messageDesc": HttpStatusCodes[500],
			"statusCode":  500,
		}) // ส่ง response ด้วย status 500
	}

	return c.Status(statusCode).JSON(body)
}

func FiberResponseErrorJson(c *fiber.Ctx, messageDesc string, statusCode int) error {
	if statusCode <= 0 {
		return c.Status(500).JSON(fiber.Map{
			"messageDesc": HttpStatusCodes[500],
			"statusCode":  500,
		}) // ส่ง response ด้วย status 500
	}

	return c.Status(statusCode).JSON(
		fiber.Map{
			"messageDesc": messageDesc,
			"statusCode":  statusCode,
		})
}
