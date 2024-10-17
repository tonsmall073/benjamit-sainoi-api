package seeds

import "bjm/db/benjamit/models"

func UnitType() []models.UnitType {
	data := []models.UnitType{
		{Name: "ชิ้น", NameEn: "Pieces"},
		{Name: "กิโลกรัม", NameEn: "Kilograms"},
		{Name: "ลิตร", NameEn: "Liters"},
		{Name: "เมตร", NameEn: "Meters"},
		{Name: "กล่อง", NameEn: "Boxes"},
		{Name: "เซ็ต", NameEn: "Sets"},
		{Name: "แพ็ค", NameEn: "Packs"},
	}
	return data
}
