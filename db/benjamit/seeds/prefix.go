package seeds

import "bjm/db/benjamit/models"

func Prefix() []models.Prefix {
	data := []models.Prefix{
		{Name: "นาย"},
		{Name: "นาง"},
		{Name: "นางสาว"},
	}
	return data
}
