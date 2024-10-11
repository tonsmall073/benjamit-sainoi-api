package chat

import (
	"bjm/db/benjamit/models"
	"bjm/src/v1/chat/dto"
	"bjm/utils"
	"encoding/json"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jsorb84/ssefiber"
	"gorm.io/gorm"
)

type ChatService struct {
	_context *gorm.DB
}

func (s ChatService) Send(
	reqModel *dto.SendRequestModel,
	resModel *dto.SendResponseModel,
	uuid string,
	sse *ssefiber.FiberSSEApp,
) *dto.SendResponseModel {
	var userData models.User
	var userDataErr error

	if uuid != "" {
		userData, userDataErr = s.fetchUserByUuid(uuid)
		if userDataErr != nil {
			resModel.MessageDesc = userDataErr.Error()
			resModel.StatusCode = 400
			return resModel
		}
	}

	insertData := s.conditionInsertChatData(reqModel, userData, uuid)

	insert, insertErr := s.insertChat(insertData)
	if insertErr != nil {
		resModel.MessageDesc = insertErr.Error()
		resModel.StatusCode = 500
		return resModel
	}

	s.mapSendResponseModel(insert, userData, resModel)
	resModel.MessageDesc = utils.HttpStatusCodes[201]
	resModel.StatusCode = 201

	resMarshal, resMarshalErr := json.Marshal(resModel)
	if resMarshalErr == nil {
		channel := sse.GetChannel("/chat/" + resModel.Data.ChannelName)
		if channel != nil {
			channel.SendEvent("message", string(resMarshal))
		}
	}

	return resModel
}

func (s ChatService) EventChat(c *fiber.Ctx, sse *ssefiber.FiberSSEApp) error {
	channelName := c.Params("channelName")

	channel := sse.GetChannel(channelName)
	if channel == nil {
		channel = sse.CreateChannel("/chat/"+channelName, "/chat/"+channelName)
		log.Println("Created chat channel: " + channelName + "\n")
	}
	return channel.ServeHTTP(c)
}

func (s ChatService) insertChat(data models.Chat) (models.Chat, error) {
	if err := s._context.Create(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (s ChatService) fetchUserByUuid(uuid string) (models.User, error) {
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

func (s ChatService) conditionInsertChatData(
	reqModel *dto.SendRequestModel,
	userData models.User,
	uuid string,
) models.Chat {
	var res models.Chat
	if uuid != "" {
		res = models.Chat{
			Message:     reqModel.Message,
			MessageType: reqModel.MessageType,
			ChannelName: reqModel.ChannelName,
			UserId:      int(userData.ID),
		}
	} else {
		res = models.Chat{
			Message:     reqModel.Message,
			MessageType: reqModel.MessageType,
			ChannelName: reqModel.ChannelName,
		}
	}
	return res
}

func (s ChatService) mapSendResponseModel(
	chatData models.Chat,
	userData models.User,
	resModel *dto.SendResponseModel,

) {
	resModel.Data = &dto.SendDataListResponseModel{
		Message:     chatData.Message,
		MessageType: chatData.MessageType,
		ChannelName: chatData.ChannelName,
		Fullname: utils.ConcatFullname(
			userData.Prefix.Name,
			userData.Firstname,
			userData.Lastname,
			"",
		),
		Nickname:  userData.Nickname,
		CreatedAt: chatData.CreatedAt,
	}
}
