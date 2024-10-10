package user

import "github.com/gofiber/fiber/v2"

func Setup(route fiber.Router) {
	group := route.Group("/user")
	group.Post("create", createUser)
}
