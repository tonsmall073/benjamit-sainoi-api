package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UnitType struct {
	gorm.Model
	UUID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();unique" json:"uuid"`
	Name   string    `gorm:"not null"`
	NameEn string    `gorm:"null;default:null"`
}
