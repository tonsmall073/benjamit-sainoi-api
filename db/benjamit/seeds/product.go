package seeds

import "bjm/db/benjamit/models"

func Product() []models.Product {
	data := []models.Product{
		{
			Name:        "Test",
			Description: "tester system",
			Image:       nil,
			UserId:      1,
		},
	}
	return data
}
