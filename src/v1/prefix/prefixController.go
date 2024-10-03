package prefix

import (
	db "bjm/db/benjamit"
	"bjm/src/v1/prefix/dto"
	"bjm/utils"

	"github.com/gofiber/fiber/v2"
)

func getAllPrefix(c *fiber.Ctx) error {
	context, contextErr := db.Connect()
	if contextErr != nil {
		return utils.FiberResponseErrorJson(c, contextErr.Error(), 500)
	}
	resModel := &dto.GetAllPrefixResponseModel{}
	service := &PrefixService{context}
	serviceRes := service.GetAllPrefix(resModel)
	return utils.FiberResponseJson(c, serviceRes, serviceRes.StatusCode)
}
