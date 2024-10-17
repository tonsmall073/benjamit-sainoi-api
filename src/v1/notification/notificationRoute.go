package notification

import (
	auth "bjm/auth/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/jsorb84/ssefiber"
)

func Setup(route fiber.Router, sse *ssefiber.FiberSSEApp) {
	group := route.Group("/notification")

	group.Get("/user/events/:token", func(c *fiber.Ctx) error { return eventNoti(c, sse) })

	groupAuth := group.Group("/user", auth.UseUserGuard)
	groupAuth.Post("/create", func(c *fiber.Ctx) error { return createNoti(c, sse) })
}
