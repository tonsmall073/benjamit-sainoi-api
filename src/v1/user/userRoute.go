package user

import (
	auth "bjm/auth/jwt"

	"github.com/gofiber/fiber/v2"
)

func Setup(route fiber.Router) {
	group := route.Group("/user")
	group.Post("/create", createUser)
	group.Post("/login", login)

	authGroup := group.Group("", auth.UseGuard)
	authGroup.Get("/profile", getProfile)
	authGroup.Put("/update", updateUser)
	authGroup.Delete("/delete", deleteUser)

}
