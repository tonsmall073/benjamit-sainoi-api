package dto

import "time"

type GetProfileResponseModel struct {
	Data        *GetProfileDataListResponseModel `json:"data"`
	MessageDesc string                           `json:"messageDesc"`
	StatusCode  int                              `json:"statusCode"`
}

type GetProfileDataListResponseModel struct {
	Uuid          string    `json:"uuid"`
	PrefixName    string    `json:"prefixName"`
	Firstname     string    `json:"firstname"`
	Lastname      string    `json:"lastname"`
	Nickname      string    `json:"nickname"`
	Birthday      time.Time `json:"birthday"`
	Email         string    `json:"email"`
	LineId        string    `json:"lineId"`
	MobilePhoneNo string    `json:"mobilePhoneNo"`
	HomePhoneNo   string    `json:"homePhoneNo"`
}
