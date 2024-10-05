package user

import (
	auth "bjm/auth/jwt"
	db "bjm/db/benjamit"
	"bjm/src/v1/user/dto"
	"bjm/utils"

	"github.com/gofiber/fiber/v2"
)

func createUser(c *fiber.Ctx) error {
	reqModel := &dto.CreateUserRequestModel{}
	resModel := &dto.CreateUserResponseModel{}
	err := c.BodyParser(reqModel)
	if err != nil {
		return utils.FiberResponseErrorJson(c, err.Error(), 400)
	}
	context, contextErr := db.Connect()
	if contextErr != nil {
		return utils.FiberResponseErrorJson(c, contextErr.Error(), 500)
	}

	service := &UserService{context}
	serviceRes := service.CreateUser(reqModel, resModel)
	return utils.FiberResponseJson(c, serviceRes, serviceRes.StatusCode)
}

func login(c *fiber.Ctx) error {
	reqModel := &dto.LoginRequestModel{}
	resModel := &dto.LoginResponseModel{}
	err := c.BodyParser(reqModel)
	if err != nil {
		return utils.FiberResponseErrorJson(c, err.Error(), 400)
	}
	context, contextErr := db.Connect()
	if contextErr != nil {
		return utils.FiberResponseErrorJson(c, contextErr.Error(), 500)
	}

	service := &UserService{context}
	serviceRes := service.Login(reqModel, resModel)
	return utils.FiberResponseJson(c, serviceRes, serviceRes.StatusCode)
}

func getProfile(c *fiber.Ctx) error {
	data := auth.DecodeToken(c)
	resModel := &dto.GetProfileResponseModel{}
	context, contextErr := db.Connect()
	if contextErr != nil {
		return utils.FiberResponseErrorJson(c, contextErr.Error(), 500)
	}
	service := &UserService{context}
	serviceRes := service.GetProfile(data["uuid"].(string), resModel)
	return utils.FiberResponseJson(c, serviceRes, serviceRes.StatusCode)
}

func updateUser(c *fiber.Ctx) error {
	data := auth.DecodeToken(c)
	return c.JSON(fiber.Map{
		"uuid":        data["uuid"],
		"messageDesc": "User updated",
	})
}

func deleteUser(c *fiber.Ctx) error {
	data := auth.DecodeToken(c)
	return c.JSON(fiber.Map{
		"uuid":        data["uuid"],
		"messageDesc": "User updated",
	})
}
