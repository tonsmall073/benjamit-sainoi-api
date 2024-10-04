package dto

type LoginRequestModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponseModel struct {
	Data        *LoginDataListResponseModel `json:"data"`
	StatusCode  int                         `json:"statusCode"`
	MessageDesc string                      `json:"messageDesc"`
}

type LoginDataListResponseModel struct {
	AccessToken string `json:"accessToken"`
	Username    string `json:"username"`
	Nickname    string `json:"nickname"`
	PrefixName  string `json:"prefixName"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
}
