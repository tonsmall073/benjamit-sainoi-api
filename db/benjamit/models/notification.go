package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	UUID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();unique" json:"uuid"`
	Title        string    `gorm:"not null"`
	Description  string    `gorm:"not null"`
	UserId       int       `gorm:"null"`
	User         User      `gorm:"foreignKey:UserId" json:"user"`
	ReadStatus   bool      `gorm:"default:false"`
	SendToUserId int       `gorm:"null"`
	SendToUser   User      `gorm:"foreignKey:SendToUserId" json:"send_to_user"`
}
