package seeds

import "bjm/db/benjamit/models"

func ProductSelling() []models.ProductSelling {
	data := []models.ProductSelling{
		{
			CostPrice:  80.50,
			SellPrice:  99.50,
			Stock:      10,
			ProductId:  1,
			UnitTypeId: 1,
			UserId:     1,
		},
	}
	return data
}
