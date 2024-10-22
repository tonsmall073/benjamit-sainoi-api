package dto

import "time"

type CreateListRequestModel struct {
	Amount                  float64   `json:"amount"`
	Description             string    `json:"description"`
	TransactionDate         time.Time `json:"transactionDate"`
	ReferProductUuid        int       `json:"referProductUuid"`
	ReferProductSellingUuid int       `json:"referProductSellingUuid"`
}

type CreateListResponseModel struct {
	StatusCode  int                              `json:"statusCode"`
	MessageDesc string                           `json:"messageDesc"`
	Data        *CreateListDataListResponseModel `json:"data"`
}

type CreateListDataListResponseModel struct {
}
