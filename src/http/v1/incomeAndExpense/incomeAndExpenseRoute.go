package incomeAndExpense

import (
	auth "bjm/auth/jwt"

	"github.com/gofiber/fiber/v2"
)

func SetUp(route fiber.Router) {
	group := route.Group("/incomeAndExpense")

	groupAuth := group.Group("/user", auth.UseUserGuard)
	groupAuth.Post("/create", createList)
	groupAuth.Post("/getAll", getAllList)
}
