package dto

import (
	"bjm/utils/enums"
	"time"
)

type GetAllListRequestModel struct {
	Search     string         `json:"search"`
	StartDate  time.Time      `json:"startDate"`
	EndDate    time.Time      `json:"endDate"`
	Take       int            `json:"take"`
	Skip       int            `json:"skip"`
	Sort       enums.SortEnum `json:"sort"`
	SortColumn string         `json:"sortColumn"`
}

type GetAllListResponseModel struct {
	StatusCode  int                                `json:"statusCode"`
	MessageDesc string                             `json:"messageDesc"`
	Data        []*GetAllListDataListResponseModel `json:"data"`
	TotalData   int                                `json:"totalData"`
}

type GetAllListDataListResponseModel struct {
	Amount             float64                                        `json:"amount"`
	Description        string                                         `json:"description"`
	TransactionDate    time.Time                                      `json:"transactionDate"`
	ProductData        *GetAllListProductDataListResponseModel        `json:"productData"`
	ProductSellingData *GetAllListProductSellingDataListResponseModel `json:"productSellingData"`
}

type GetAllListProductDataListResponseModel struct {
	Name string `json:"name"`
}

type GetAllListProductSellingDataListResponseModel struct {
	SellPrice    float64                                  `json:"SellPrice"`
	CostPrice    float64                                  `json:"CostPrice"`
	UnitTypeData *GetAllListUnitTypeDataListResponseModel `json:"unitTypeData"`
}

type GetAllListUnitTypeDataListResponseModel struct {
	Name   string `json:"name"`
	NameEn string `json:"nameEn"`
}
