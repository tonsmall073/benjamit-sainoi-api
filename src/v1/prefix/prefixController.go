package prefix

import (
	db "bjm/db/benjamit"
	"bjm/src/v1/prefix/dto"
	"bjm/utils"

	"github.com/gofiber/fiber/v2"
)

// @Tags Prefix
// @Description Fetch all prefixes
// @Produce json
// @Success 200 {object} dto.GetAllPrefixResponseModel
// @Failure 400 {object} utils.ErrorResponseModel "invalid input"
// @Failure 500 {object} utils.ErrorResponseModel "internal server error"
// @Router /v1/prefix/getAll [get]
func getAllPrefix(c *fiber.Ctx) error {
	context, contextErr := db.Connect()
	if contextErr != nil {
		return utils.FiberResponseErrorJson(c, contextErr.Error(), 500)
	}
	defer db.ConnectClose(context)

	resModel := &dto.GetAllPrefixResponseModel{}
	service := &PrefixService{context}
	serviceRes := service.GetAllPrefix(resModel)
	return utils.FiberResponseJson(c, serviceRes, serviceRes.StatusCode)
}
