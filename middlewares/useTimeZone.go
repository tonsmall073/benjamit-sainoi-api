package middlewares

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UseTimeZone(app *fiber.App) {
	defaultStr := "Asia/Bangkok"

	if getTimeZone := os.Getenv("SERVER_TIME_ZONE"); getTimeZone != "" {
		defaultStr = getTimeZone
	}
	loc, err := time.LoadLocation(defaultStr)
	if err != nil {
		panic(err)
	}

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("timezone", loc)
		return c.Next()
	})
}
