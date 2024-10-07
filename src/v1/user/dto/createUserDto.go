package dto

import (
	"time"
)

type CreateUserRequestModel struct {
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	PrefixUuid    string    `json:"prefixUuid"`
	Nickname      string    `json:"nickname"`
	Firstname     string    `json:"firstname"`
	Lastname      string    `json:"lastname"`
	Birthday      time.Time `json:"birthday"`
	Email         string    `json:"email"`
	LindId        string    `json:"lindid"`
	MobilePhoneNo string    `json:"mobilePhoneNo"`
	HomePhoneNo   string    `json:"homePhoneNostring"`
}

type CreateUserResponseModel struct {
	Data        *CreateUserDataListResponseModel `json:"data"`
	MessageDesc string                           `json:"messageDesc"`
	StatusCode  int                              `json:"statusCode"`
}

type CreateUserDataListResponseModel struct {
	Username   string `json:"username"`
	PrefixName string `json:"prefixName"`
	Nickname   string `json:"nickname"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
}
