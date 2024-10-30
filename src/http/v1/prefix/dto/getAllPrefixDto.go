package dto

type GetAllPrefixResponseModel struct {
	Data        []*GetAllPrefixDataListResponseModel `json:"data"`
	MessageDesc string                               `json:"messageDesc"`
	StatusCode  int                                  `json:"statusCode"`
}

type GetAllPrefixDataListResponseModel struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}
