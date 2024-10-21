package dto

import "time"

type CreateNotiRequestModel struct {
	SendToUserUuid string `json:"sendToUserUuid" validate:"required"`
	Title          string `json:"title" validate:"required"`
	Description    string `json:"description" validate:"required"`
}

type CreateNotiResponseModel struct {
	Data        *CreateNotiDataListResponseModel `json:"data"`
	MessageDesc string                           `json:"messageDesc"`
	StatusCode  int                              `json:"statusCode"`
}

type CreateNotiDataListResponseModel struct {
	UserUuid       string    `json:"userUuid"`
	SendToUserUuid string    `json:"sendToUserUuid"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Fullname       string    `json:"fullname"`
	Nickname       string    `json:"nickname"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	ReadStatus     bool      `json:"readStatus"`
}
