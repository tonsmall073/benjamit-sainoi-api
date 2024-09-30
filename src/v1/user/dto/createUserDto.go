package dto

type CreateUserRequestModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserResponseModel struct {
	Data        *CreateUserDataListResponseModel `json:"data"`
	MessageDesc string                           `json:"messageDesc"`
	Status      int                              `json:"status"`
}

type CreateUserDataListResponseModel struct {
	Username string `json:"username"`
}
