package dto

import (
	"bjm/utils/enums"
	"time"
)

type CreateListRequestModel struct {
	Quantity                int                       `json:"quantity"`
	Amount                  float64                   `json:"amount" validate:"required"`
	Description             string                    `json:"description" validate:"required"`
	TransactionDate         time.Time                 `json:"transactionDate" validate:"required"`
	ReferProductUuid        string                    `json:"referProductUuid"`
	ReferProductSellingUuid string                    `json:"referProductSellingUuid"`
	TransactionType         enums.TransactionTypeEnum `json:"transactionType" validate:"required,transactionTypeEnum"`
	ReferProductStatus      bool                      `json:"referProductStatus"`
}

type CreateListResponseModel struct {
	StatusCode  int                              `json:"statusCode"`
	MessageDesc string                           `json:"messageDesc"`
	Data        *CreateListDataListResponseModel `json:"data"`
}

type CreateListDataListResponseModel struct {
	Amount             float64                                        `json:"amount"`
	Description        string                                         `json:"description"`
	TransactionDate    time.Time                                      `json:"transactionDate"`
	ProductData        *CreateListProductDataListResponseModel        `json:"productData"`
	ProductSellingData *CreateListProductSellingDataListResponseModel `json:"productSellingData"`
}

type CreateListProductDataListResponseModel struct {
	Name string `json:"name"`
}

type CreateListProductSellingDataListResponseModel struct {
	SellPrice    float64                                  `json:"SellPrice"`
	CostPrice    float64                                  `json:"CostPrice"`
	UnitTypeData *CreateListUnitTypeDataListResponseModel `json:"unitTypeData"`
}

type CreateListUnitTypeDataListResponseModel struct {
	Name   string `json:"name"`
	NameEn string `json:"nameEn"`
}
