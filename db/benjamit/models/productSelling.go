package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductSelling struct {
	gorm.Model
	UUID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();unique" json:"uuid"`
	SellPrice  float64   `gorm:"not null;default:0"`
	CostPrice  float64   `gorm:"not null;default:0"`
	Stock      int       `gorm:"not null;default:0"`
	ProductId  int       `gorm:"not null"`
	Product    Product   `gorm:"foreignKey:ProductId" json:"product"`
	UnitTypeId int       `gorm:"not null"`
	UnitType   UnitType  `gorm:"foreignKey:UnitTypeId" json:"unit_type"`
	UserId     int       `gorm:"not null"`
	User       User      `gorm:"foreignKey:UserId" json:"user"`
}
