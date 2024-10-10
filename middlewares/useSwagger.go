package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func UseSwagger(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/swagger.json", func(c *fiber.Ctx) error {
		err := c.SendFile("./docs/swagger.json")
		if err != nil {
			return c.Status(500).SendString("Failed to serve swagger.json: " + err.Error())
		}
		return nil
	})
}
