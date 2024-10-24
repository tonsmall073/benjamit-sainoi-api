package notification

import (
	auth "bjm/auth/jwt"
	db "bjm/db/benjamit"
	"bjm/src/v1/notification/dto"
	"bjm/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/jsorb84/ssefiber"
)

// @Tags Notification
// @Description Create notification
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body dto.CreateNotiRequestModel true "notification request"
// @Success 201 {object} dto.CreateNotiResponseModel "created"
// @Failure 400 {object} utils.ErrorResponseModel "invalid input"
// @Failure 401 {object} utils.ErrorResponseModel "unauthorized"
// @Failure 500 {object} utils.ErrorResponseModel "internal server error"
// @Router /v1/notification/user/create [post]
func createNoti(c *fiber.Ctx, sse *ssefiber.FiberSSEApp) error {
	reqModel := &dto.CreateNotiRequestModel{}

	if err := c.BodyParser(reqModel); err != nil {
		return utils.FiberResponseErrorJson(c, err.Error(), 400)
	}

	if err := utils.Validate.Struct(reqModel); err != nil {
		return utils.FiberResponseErrorJson(c, err.Error(), 400)
	}

	context, contextErr := db.Connect()
	defer db.ConnectClose(context)
	if contextErr != nil {
		return utils.FiberResponseErrorJson(c, contextErr.Error(), 500)
	}

	getUuid := auth.DecodeToken(c)["uuid"].(string)

	resModel := &dto.CreateNotiResponseModel{}
	service := &NotificationService{context}
	serviceRes := service.CreateNoti(reqModel, resModel, getUuid, sse)

	return utils.FiberResponseJson(c, serviceRes, serviceRes.StatusCode)
}

// @Tags Notification
// @Description Event notification
// @Param token path string true "token"
// @Success 201 {object} dto.CreateNotiResponseModel "created"
// @Failure 401 {object} utils.ErrorResponseModel "unauthorized"
// @Failure 500 {object} utils.ErrorResponseModel "internal server error"
// @Router /v1/notification/user/events/{token} [get]
func eventNoti(c *fiber.Ctx, sse *ssefiber.FiberSSEApp) error {
	token := c.Params("token")
	tokenData, tokenDataErr := auth.DecodeTokenByTokenStr(token)
	if tokenDataErr != nil {
		return utils.FiberResponseErrorJson(c, tokenDataErr.Error(), 500)
	}
	getUuid := tokenData["uuid"].(string)
	if getUuid == "" {
		return utils.FiberResponseErrorJson(c, utils.HttpStatusCodes[401], 401)
	}

	service := &NotificationService{}
	serviceRes := service.EventNoti(c, sse, getUuid)
	return serviceRes
}
