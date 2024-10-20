package middlewares

import (
	"bjm/utils"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func UseRequestLimit(app *fiber.App) {
	defaultMax := 20
	defaultExpiration := time.Duration(30)
	if convInt, convErr := strconv.Atoi(os.Getenv("REQUEST_SEND_LIMIT_MAX")); convErr == nil {
		if convInt > 0 {
			defaultMax = convInt
		}
	}
	if convInt, convErr := strconv.Atoi(os.Getenv("REQUEST_SEND_DELAY_LIMIT_SECONDS")); convErr == nil {
		if convInt > 0 {
			defaultExpiration = time.Duration(convInt)
		}
	}
	app.Use(limiter.New(limiter.Config{
		Max:               defaultMax,
		Expiration:        defaultExpiration * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return utils.FiberResponseErrorJson(c,
				utils.HttpStatusCodes[fiber.StatusTooManyRequests],
				fiber.StatusTooManyRequests,
			)
		},
		Next: func(c *fiber.Ctx) bool {
			return utils.IsSwaggerPath(c.Path())
		},
	}))
}
