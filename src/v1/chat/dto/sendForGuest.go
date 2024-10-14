package dto

import (
	"bjm/db/benjamit/models"
	"time"
)

type SendForGuestRequestModel struct {
	Message     string                 `json:"message"`
	MessageType models.MessageTypeEnum `json:"messageType"`
}

type SendForGuestResponseModel struct {
	Data        *SendForGuestDataListResponseModel `json:"data"`
	MessageDesc string                             `json:"messageDesc"`
	StatusCode  int                                `json:"statusCode"`
}

type SendForGuestDataListResponseModel struct {
	Message     string                 `json:"message"`
	MessageType models.MessageTypeEnum `json:"messageType"`
	ChannelName string                 `json:"channelName"`
	CreatedAt   time.Time              `json:"createdAt"`
}
