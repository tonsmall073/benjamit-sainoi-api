package v1

import (
	"bjm/src/http/v1/chat"
	"bjm/src/http/v1/incomeAndExpense"
	"bjm/src/http/v1/notification"
	"bjm/src/http/v1/prefix"
	"bjm/src/http/v1/product"
	"bjm/src/http/v1/user"

	"github.com/gofiber/fiber/v2"
	"github.com/jsorb84/ssefiber"
)

func UseRoute(app fiber.Router, sse *ssefiber.FiberSSEApp) {
	route := app.Group("/v1")
	prefix.Setup(route)
	user.Setup(route)
	product.Setup(route)
	chat.Setup(route, sse)
	notification.Setup(route, sse)
	incomeAndExpense.SetUp(route)
}
