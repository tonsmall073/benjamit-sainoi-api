package dto

import (
	types "bjm/utils"
	"time"
)

type Date time.Time
type CreateUserRequestModel struct {
	Username   string     `json:"username"`
	Password   string     `json:"password"`
	PrefixUuid string     `json:"prefixUuid"`
	Firstname  string     `json:"firstname"`
	Lastname   string     `json:"lastname"`
	Birthday   types.Date `json:"birthday"`
}

type CreateUserResponseModel struct {
	Data        *CreateUserDataListResponseModel `json:"data"`
	MessageDesc string                           `json:"messageDesc"`
	Status      int                              `json:"status"`
}

type CreateUserDataListResponseModel struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
