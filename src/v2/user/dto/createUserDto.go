package dto

type CreateUserRequestModel struct {
	Nickname  string `json:"nickname"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type CreateUserResponseModel struct {
	Data        *CreateUserDataListResponseModel `json:"data"`
	MessageDesc string                           `json:"messageDesc"`
	StatusCode  int                              `json:"statusCode"`
}

type CreateUserDataListResponseModel struct {
	Nickname  string `json:"nickname"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
