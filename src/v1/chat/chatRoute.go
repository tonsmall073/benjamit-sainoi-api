package chat

import (
	auth "bjm/auth/jwt"

	"github.com/gofiber/fiber/v2"
)

func Setup(route fiber.Router) {
	group := route.Group("/chat")
	group.Post("/send", sendForGuest)

	group.Get("/events/:channel", eventChat)

	groupAuth := group.Group("/user", auth.UseGuard)
	groupAuth.Post("/send", send)
}
