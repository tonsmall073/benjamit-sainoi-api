package prefix

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(route fiber.Router) {
	group := route.Group("/prefix")
	group.Get("/getAll", getAllPrefix)
}
