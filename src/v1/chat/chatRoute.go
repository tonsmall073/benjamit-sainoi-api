package chat

import (
	auth "bjm/auth/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/jsorb84/ssefiber"
)

func Setup(route fiber.Router, sse *ssefiber.FiberSSEApp) {
	group := route.Group("/chat")
	group.Post("/send", func(c *fiber.Ctx) error { return sendForGuest(c, sse) })

	group.Get("/events/guest", func(c *fiber.Ctx) error { return eventChatGuest(c, sse) })
	group.Get("/websocket/send", websocket.New(func(wsCon *websocket.Conn) { wsSendForGuest(wsCon) }))
	group.Get("/user/events/:channelName", func(c *fiber.Ctx) error { return eventChat(c, sse) })

	groupAuth := group.Group("/user", auth.UseUserGuard)
	groupAuth.Post("/send", func(c *fiber.Ctx) error { return send(c, sse) })
	groupAuth.Get("/websocket/send/:channelName", websocket.New(func(wsCon *websocket.Conn) { wsSend(wsCon) }))

}
