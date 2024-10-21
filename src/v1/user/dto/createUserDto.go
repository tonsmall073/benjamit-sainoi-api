package dto

import (
	"time"
)

type CreateUserRequestModel struct {
	Username      string    `json:"username" validate:"required,min=7"`
	Password      string    `json:"password" validate:"required,min=8"`
	PrefixUuid    string    `json:"prefixUuid" validate:"required"`
	Nickname      string    `json:"nickname" validate:"required"`
	Firstname     string    `json:"firstname" validate:"required"`
	Lastname      string    `json:"lastname" validate:"required"`
	Birthday      time.Time `json:"birthday" validate:"required"`
	Email         string    `json:"email" validate:"required,email"`
	LindId        string    `json:"lindid" validate:"required"`
	MobilePhoneNo string    `json:"mobilePhoneNo" validate:"required,phone"`
	HomePhoneNo   string    `json:"homePhoneNo"`
}

type CreateUserResponseModel struct {
	Data        *CreateUserDataListResponseModel `json:"data"`
	MessageDesc string                           `json:"messageDesc"`
	StatusCode  int                              `json:"statusCode"`
}

type CreateUserDataListResponseModel struct {
	Username   string    `json:"username"`
	PrefixName string    `json:"prefixName"`
	Nickname   string    `json:"nickname"`
	Firstname  string    `json:"firstname"`
	Lastname   string    `json:"lastname"`
	Birthday   time.Time `json:"birthday"`
}
