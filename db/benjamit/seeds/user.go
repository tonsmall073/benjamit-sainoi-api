package seeds

import (
	"bjm/db/benjamit/models"
	"bjm/utils"
	"bjm/utils/enums"
)

func User() []models.User {
	data := []models.User{
		{
			Username:      "tonsmall",
			Password:      "$2a$12$P47EMd8pvduJrXul64zDhu8GTzZiVH8NpHsD6Pf25MENAZxF26CnS", //ton1234
			PrefixId:      1,
			Nickname:      "ต้น",
			Firstname:     "วิศรุต",
			Lastname:      "รูปเขียน",
			Birthday:      utils.ConvTime("1992-01-11T21:22:23.000Z"),
			Email:         "tonsmall073@gmail.com",
			LineId:        "tonsmall073",
			MobilePhoneNo: "0819999999",
			HomePhoneNo:   "029999999",
			Role:          enums.USER,
		},
		{
			Username:      "admin",
			Password:      "$2a$12$P47EMd8pvduJrXul64zDhu8GTzZiVH8NpHsD6Pf25MENAZxF26CnS", //ton1234
			PrefixId:      1,
			Nickname:      "ต้น",
			Firstname:     "วิศรุต",
			Lastname:      "รูปเขียน",
			Birthday:      utils.ConvTime("1992-01-11T21:22:23.000Z"),
			Email:         "admin@gmail.com",
			LineId:        "tonsmall073",
			MobilePhoneNo: "0819999999",
			HomePhoneNo:   "029999999",
			Role:          enums.ADMIN,
		},
	}
	return data
}
