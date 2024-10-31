package utils

import "github.com/gofiber/fiber/v2"

func FiberResponseJson(c *fiber.Ctx, body interface{}, statusCode int) error {
	if statusCode <= 0 {
		return c.Status(500).JSON(
			&ErrorResponseModel{
				MessageDesc: HttpStatusCodes[500],
				StatusCode:  500,
			},
		)
	}

	return c.Status(statusCode).JSON(body)
}

func FiberResponseErrorJson(c *fiber.Ctx, messageDesc string, statusCode int) error {
	if statusCode <= 0 {
		return c.Status(500).JSON(
			&ErrorResponseModel{
				MessageDesc: HttpStatusCodes[500],
				StatusCode:  500,
			},
		)
	}

	return c.Status(statusCode).JSON(
		&ErrorResponseModel{
			MessageDesc: messageDesc,
			StatusCode:  statusCode,
		},
	)
}
