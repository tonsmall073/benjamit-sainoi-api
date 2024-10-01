package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"uuid"`
	Username  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Firstname string    `gorm:"not null"`
	Lastname  string    `gorm:"not null"`
	Birthday  time.Time `json:"birthday"`
	PrefixId  int       `gorm:"not null" json:"prefix_id"`
	Prefix    Prefix    `gorm:"foreignKey:PrefixId" json:"prefix"`
}
