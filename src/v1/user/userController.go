package user

import (
	db "bjm/db/benjamit"
	"bjm/src/v1/user/dto"
	"bjm/utils"

	"github.com/gofiber/fiber/v2"
)

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
	reqModel := &dto.CreateUserRequestModel{}
	resModel := &dto.CreateUserResponseModel{}
	err := c.BodyParser(reqModel)
	if err != nil {
		return utils.FiberResponseErrorJson(c, err.Error(), 400)
	}

	context, _ := db.Connect()
	service := &UserService{context}
	serviceRes := service.CreateUser(reqModel, resModel)
	return utils.FiberResponseJson(c, serviceRes, serviceRes.StatusCode)
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
