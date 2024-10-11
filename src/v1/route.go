package v1

import (
	"bjm/src/v1/chat"
	"bjm/src/v1/prefix"
	"bjm/src/v1/product"
	"bjm/src/v1/user"

	"github.com/gofiber/fiber/v2"
)

func UseRoute(app fiber.Router) {
	route := app.Group("/v1")
	prefix.Setup(route)
	user.Setup(route)
	product.Setup(route)
	chat.Setup(route)
}
