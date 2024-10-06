package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func UseSwagger(app *fiber.App) {
	app.Get("/v1/swagger/*", swagger.HandlerDefault)

	app.Get("/v1/swagger.json", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/v1/swagger.json")
	})
}
