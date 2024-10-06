package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	UUID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"uuid"`
	Path         string    `gorm:"not null"`
	Method       string    `gorm:"not null"`
	ContentType  string
	RequestBody  string
	StatusCode   int
	ResponseBody string
	Headers      string
	Origin       string
}
