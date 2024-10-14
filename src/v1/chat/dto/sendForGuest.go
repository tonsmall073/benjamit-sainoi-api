package dto

import (
	"bjm/db/benjamit/models"
	"time"
)

type SendForGuestRequestModel struct {
	ClientId    string                 `json:"clientId"`
	Message     string                 `json:"message"`
	MessageType models.MessageTypeEnum `json:"messageType"`
	Fullname    string                 `json:"fullname"`
	Nickname    string                 `json:"nickname"`
}

type SendForGuestResponseModel struct {
	Data        *SendForGuestDataListResponseModel `json:"data"`
	MessageDesc string                             `json:"messageDesc"`
	StatusCode  int                                `json:"statusCode"`
}

type SendForGuestDataListResponseModel struct {
	ClientId    string                 `json:"clientId"`
	Message     string                 `json:"message"`
	MessageType models.MessageTypeEnum `json:"messageType"`
	ChannelName string                 `json:"channelName"`
	CreatedAt   time.Time              `json:"createdAt"`
	Fullname    string                 `json:"fullname"`
	Nickname    string                 `json:"nickname"`
	ReadStatus  bool                   `json:"readStatus"`
}
