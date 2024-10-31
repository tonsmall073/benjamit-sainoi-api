package utils

type ErrorResponseModel struct {
	MessageDesc string `json:"messageDesc"`
	StatusCode  int    `json:"statusCode"`
}
