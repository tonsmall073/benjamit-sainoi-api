package chat

import (
	auth "bjm/auth/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/jsorb84/ssefiber"
)

func Setup(route fiber.Router, sse *ssefiber.FiberSSEApp) {
	group := route.Group("/chat")
	group.Post("/send", func(c *fiber.Ctx) error { return sendForGuest(c, sse) })
	group.Get("/events/:channelName", func(c *fiber.Ctx) error { return eventChat(c, sse) })

	groupAuth := group.Group("/user", auth.UseGuard)
	groupAuth.Post("/send", func(c *fiber.Ctx) error { return send(c, sse) })
}
