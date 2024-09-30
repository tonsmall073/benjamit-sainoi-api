package user

import "github.com/gofiber/fiber/v2"

func getUsers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Get all users",
	})
}

func getUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{
		"message": "Get user by ID",
		"id":      id,
	})
}

func createUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "User created",
	})
}

func updateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{
		"message": "User updated",
		"id":      id,
	})
}

func deleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{
		"message": "User deleted",
		"id":      id,
	})
}
