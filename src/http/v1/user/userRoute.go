package user

import (
	auth "bjm/auth/jwt"

	"github.com/gofiber/fiber/v2"
)

func Setup(route fiber.Router) {
	group := route.Group("/user")
	group.Post("/create", createUser)
	group.Post("/login", login)

	authGroup := group.Group("", auth.UseUserGuard)
	authGroup.Get("/profile", getProfile)
	authGroup.Put("/update", updateUser)

	authAdminGroup := group.Group("/admin", auth.UseAdminGuard)
	authAdminGroup.Delete("/delete", deleteUser)
	authAdminGroup.Get("/test", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/plain")
		return c.Status(200).SendString("test ok")
	})
}
