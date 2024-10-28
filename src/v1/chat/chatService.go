package chat

import (
	"bjm/db/benjamit/models"
	"bjm/src/v1/chat/dto"
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

type ChatService struct {
	_context *gorm.DB
}

func (s *ChatService) Send(
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
	if resMarshalErr == nil && sse != nil {
		channel := sse.GetChannel("/chat/user/" + resModel.Data.ChannelName)
		if channel != nil {
			channel.SendEvent("message", string(resMarshal))
		}
	}

	return resModel
}

func (s *ChatService) sendForGuest(
	reqModel *dto.SendForGuestRequestModel,
	resModel *dto.SendForGuestResponseModel,
	sse *ssefiber.FiberSSEApp,
) *dto.SendForGuestResponseModel {
	insertData := models.Chat{
		Message:     reqModel.Message,
		MessageType: reqModel.MessageType,
		ChannelName: "guest",
		ClientId:    reqModel.ClientId,
	}

	insert, insertErr := s.insertChat(insertData)
	if insertErr != nil {
		resModel.MessageDesc = insertErr.Error()
		resModel.StatusCode = 500
		return resModel
	}

	s.mapSendForGuestResponseModel(insert, reqModel, resModel)
	resModel.MessageDesc = utils.HttpStatusCodes[201]
	resModel.StatusCode = 201

	resMarshal, resMarshalErr := json.Marshal(resModel)
	if resMarshalErr == nil && sse != nil {
		channel := sse.GetChannel("/chat/" + resModel.Data.ChannelName)
		if channel != nil {
			channel.SendEvent("message", string(resMarshal))
		}
	}

	return resModel
}

func (s *ChatService) WsSendForGuest(
	wsCon *websocket.Conn,
) error {
	defer wsCon.Close()
	wsClients[wsCon] = "guest"

	for {
		reqModel := &dto.SendForGuestRequestModel{}
		resModel := &dto.SendForGuestResponseModel{}
		err := wsCon.ReadJSON(reqModel)
		if err != nil {
			delete(wsClients, wsCon)
			fmt.Println("[ERROR] read json error:", err)
			return err
		}

		s.sendForGuest(reqModel, resModel, nil)

		for client, channel := range wsClients {
			if channel == "guest" {
				if err := client.WriteJSON(resModel); err != nil {
					fmt.Println("[ERROR] write json error:", err)
					client.Close()
					delete(wsClients, client)
				}
			}
		}
	}
}

func (s *ChatService) WsSend(
	uuid string,
	wsCon *websocket.Conn,
) error {
	defer wsCon.Close()
	wsClients[wsCon] = wsCon.Params("channelName")

	for {
		reqModel := &dto.SendRequestModel{}
		resModel := &dto.SendResponseModel{}
		err := wsCon.ReadJSON(reqModel)

		if err != nil {
			delete(wsClients, wsCon)
			fmt.Println("[ERROR] read json error:", err)
			return err
		}

		s.Send(reqModel, resModel, uuid, nil)

		for client, channelName := range wsClients {
			if channelName == reqModel.ChannelName {
				if err := client.WriteJSON(resModel); err != nil {
					fmt.Println("[ERROR] write json error:", err)
					client.Close()
					delete(wsClients, client)
				}
			}
		}
	}
}

func (s *ChatService) EventChat(c *fiber.Ctx, sse *ssefiber.FiberSSEApp) error {
	channelName := c.Params("channelName")

	channel := sse.GetChannel(channelName)
	if channel == nil {
		channel = sse.CreateChannel("/chat/user/"+channelName, "/chat/user/"+channelName)
		log.Println("Created chat channel: " + channelName)
	}
	return channel.ServeHTTP(c)
}

func (s *ChatService) EventChatForGuest(c *fiber.Ctx, sse *ssefiber.FiberSSEApp) error {
	channel := sse.GetChannel("guest")
	if channel == nil {
		channel = sse.CreateChannel("/chat/guest", "/chat/guest")
		log.Println("Created chat channel: guest")
	}
	return channel.ServeHTTP(c)
}

func (s *ChatService) insertChat(data models.Chat) (models.Chat, error) {
	if err := s._context.Create(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (s *ChatService) fetchUserByUuid(uuid string) (models.User, error) {
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

func (s *ChatService) conditionInsertChatData(
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

func (s *ChatService) mapSendResponseModel(
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
		Nickname:   userData.Nickname,
		CreatedAt:  chatData.CreatedAt,
		Uuid:       userData.UUID.String(),
		ReadStatus: chatData.ReadStatus,
	}
}

func (s *ChatService) mapSendForGuestResponseModel(
	chatData models.Chat,
	reqModel *dto.SendForGuestRequestModel,
	resModel *dto.SendForGuestResponseModel,

) {
	resModel.Data = &dto.SendForGuestDataListResponseModel{
		Message:     chatData.Message,
		MessageType: chatData.MessageType,
		ChannelName: "guest",
		CreatedAt:   chatData.CreatedAt,
		ClientId:    reqModel.ClientId,
		Fullname:    reqModel.Fullname,
		Nickname:    reqModel.Nickname,
		ReadStatus:  chatData.ReadStatus,
	}
}
