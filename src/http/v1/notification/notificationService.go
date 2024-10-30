package notification

import (
	"bjm/db/benjamit/models"
	"bjm/src/http/v1/notification/dto"
	"bjm/utils"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/jsorb84/ssefiber"
	"gorm.io/gorm"
)

var wsClients = make(map[*websocket.Conn]string)

type NotificationService struct {
	_context *gorm.DB
}

func (s *NotificationService) CreateNoti(
	reqModel *dto.CreateNotiRequestModel,
	resModel *dto.CreateNotiResponseModel,
	uuid string,
	sse *ssefiber.FiberSSEApp,
) *dto.CreateNotiResponseModel {

	userData, userDataErr := s.fetchUserByUuid(uuid)
	if userDataErr != nil {
		resModel.MessageDesc = userDataErr.Error()
		resModel.StatusCode = 400
		return resModel
	}
	sendToUserData, sendToUserDataErr := s.fetchUserByUuid(reqModel.SendToUserUuid)
	if sendToUserDataErr != nil {
		resModel.MessageDesc = "send to '" + reqModel.SendToUserUuid + "' " + sendToUserDataErr.Error()
		resModel.StatusCode = 400
		return resModel
	}

	indertData := models.Notification{
		Title:        reqModel.Title,
		Description:  reqModel.Description,
		UserId:       int(userData.ID),
		SendToUserId: int(sendToUserData.ID),
	}
	insert, insertErr := s.insertNoti(indertData)
	if insertErr != nil {
		resModel.MessageDesc = insertErr.Error()
		resModel.StatusCode = 500
		return resModel
	}

	s.mapCreateNotiResponseModel(insert, userData, sendToUserData, resModel)

	resModel.MessageDesc = utils.HttpStatusCodes[201]
	resModel.StatusCode = 201

	resMarshal, resMarshalErr := json.Marshal(resModel)
	if resMarshalErr == nil && sse != nil {
		channel := sse.GetChannel("/noti/user/" + reqModel.SendToUserUuid)
		if channel != nil {
			channel.SendEvent("message", string(resMarshal))
		}
	}

	return resModel
}

func (s *NotificationService) WsCreateNoti(
	uuid string,
	wsCon *websocket.Conn,
) error {
	defer wsCon.Close()
	wsClients[wsCon] = uuid

	for {
		reqModel := &dto.CreateNotiRequestModel{}
		resModel := &dto.CreateNotiResponseModel{}
		err := wsCon.ReadJSON(reqModel)
		if err != nil {
			delete(wsClients, wsCon)
			fmt.Println("[ERROR] read json error:", err)
			return err
		}

		s.CreateNoti(reqModel, resModel, uuid, nil)

		for client, uuid := range wsClients {
			if uuid == reqModel.SendToUserUuid {
				if err := client.WriteJSON(resModel); err != nil {
					fmt.Println("[ERROR] write json error:", err)
					client.Close()
					delete(wsClients, client)
				}
			}
		}
	}
}

func (s *NotificationService) EventNoti(c *fiber.Ctx, sse *ssefiber.FiberSSEApp, uuid string) error {

	channel := sse.GetChannel(uuid)
	if channel == nil {
		channel = sse.CreateChannel("/noti/user/"+uuid, "/noti/user/"+uuid)
		log.Println("Created notificationService channel: " + uuid)
	}
	return channel.ServeHTTP(c)
}

func (s *NotificationService) fetchUserByUuid(uuid string) (models.User, error) {
	user := models.User{}
	result := s._context.Preload("Prefix").Where("uuid = ? AND deleted_at IS NULL", uuid).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return user, errors.New("user information not found")
		}
		return user, result.Error
	}
	return user, nil
}

func (s *NotificationService) insertNoti(data models.Notification) (models.Notification, error) {
	if err := s._context.Create(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (s *NotificationService) mapCreateNotiResponseModel(
	inserNotiData models.Notification,
	userData models.User,
	sendToUserData models.User,
	resModel *dto.CreateNotiResponseModel,
) {
	resModel.Data = &dto.CreateNotiDataListResponseModel{
		Title:       inserNotiData.Title,
		Description: inserNotiData.Description,
		Fullname: utils.ConcatFullname(
			userData.Prefix.Name,
			userData.Firstname,
			userData.Lastname,
			"",
		),
		Nickname:       userData.Nickname,
		CreatedAt:      inserNotiData.CreatedAt,
		UpdatedAt:      inserNotiData.UpdatedAt,
		ReadStatus:     inserNotiData.ReadStatus,
		SendToUserUuid: sendToUserData.UUID.String(),
		UserUuid:       userData.UUID.String(),
	}
}
