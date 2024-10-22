package dto

import (
	"bjm/db/benjamit/models"
	"time"
)

type SendRequestModel struct {
	Message     string                 `json:"message" validate:"required"`
	MessageType models.MessageTypeEnum `json:"messageType" validate:"required,messageTypeEnum"`
	ChannelName string                 `json:"channelName" validate:"required"`
}

type SendResponseModel struct {
	Data        *SendDataListResponseModel `json:"data"`
	MessageDesc string                     `json:"messageDesc"`
	StatusCode  int                        `json:"statusCode"`
}

type SendDataListResponseModel struct {
	Uuid        string                 `json:"uuid"`
	Message     string                 `json:"message"`
	MessageType models.MessageTypeEnum `json:"messageType"`
	ChannelName string                 `json:"channelName"`
	Fullname    string                 `json:"fullname"`
	Nickname    string                 `json:"nickname"`
	CreatedAt   time.Time              `json:"createdAt"`
	ReadStatus  bool                   `json:"readStatus"`
}
