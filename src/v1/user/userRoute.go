package user

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(route fiber.Router) {
	group := route.Group("/user")
	group.Get("/", getUsers)
	group.Get("/:id", getUserByID)
	group.Post("/createUser", createUser)
	group.Put("/:id", updateUser)
	group.Delete("/:id", deleteUser)
}
