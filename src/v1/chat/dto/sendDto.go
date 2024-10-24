package dto

import (
	"bjm/utils/enums"
	"time"
)

type SendRequestModel struct {
	Message     string                `json:"message" validate:"required"`
	MessageType enums.MessageTypeEnum `json:"messageType" validate:"required,messageTypeEnum"`
	ChannelName string                `json:"channelName" validate:"required"`
}

type SendResponseModel struct {
	Data        *SendDataListResponseModel `json:"data"`
	MessageDesc string                     `json:"messageDesc"`
	StatusCode  int                        `json:"statusCode"`
}

type SendDataListResponseModel struct {
	Uuid        string                `json:"uuid"`
	Message     string                `json:"message"`
	MessageType enums.MessageTypeEnum `json:"messageType"`
	ChannelName string                `json:"channelName"`
	Fullname    string                `json:"fullname"`
	Nickname    string                `json:"nickname"`
	CreatedAt   time.Time             `json:"createdAt"`
	ReadStatus  bool                  `json:"readStatus"`
}
