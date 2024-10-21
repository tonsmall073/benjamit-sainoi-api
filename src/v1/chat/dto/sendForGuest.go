package dto

import (
	"bjm/db/benjamit/models"
	"time"
)

type SendForGuestRequestModel struct {
	ClientId    string                 `json:"clientId" validate:"required"`
	Message     string                 `json:"message" validate:"required"`
	MessageType models.MessageTypeEnum `json:"messageType" validate:"required,messageTypeEnum"`
	Fullname    string                 `json:"fullname" validate:"required"`
	Nickname    string                 `json:"nickname" validate:"required"`
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
