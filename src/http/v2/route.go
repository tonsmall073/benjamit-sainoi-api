package v2

import (
	"bjm/src/http/v2/user"

	"github.com/gofiber/fiber/v2"
)

func UseRoute(app fiber.Router) {
	route := app.Group("/v2")
	user.Setup(route)
}
