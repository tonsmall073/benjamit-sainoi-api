package v2

import (
	"github.com/gofiber/fiber/v2"
)

func UseRoute(app fiber.Router) {
	route := app.Group("/v2")
	route.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
