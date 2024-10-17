package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	UUID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();unique" json:"uuid"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"null"`
	Image       []string  `gorm:"type:text[]" json:"image"`
	UserId      int       `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserId" json:"user"`
}
