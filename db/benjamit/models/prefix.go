package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Prefix struct {
	gorm.Model
	UUID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"uuid"`
	Name string    `gorm:"not null"`
}
