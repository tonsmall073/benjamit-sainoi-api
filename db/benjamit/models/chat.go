package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MessageTypeEnum string

const (
	TEXT  MessageTypeEnum = "TEXT"
	IMAGE MessageTypeEnum = "IMAGE"
	EMOJI MessageTypeEnum = "EMOJI"
)

type Chat struct {
	gorm.Model
	UUID        uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4()" json:"uuid"`
	Message     string          `gorm:"not null"`
	MessageType MessageTypeEnum `gorm:"not null;default:'TEXT'"`
	UserId      int             `gorm:"null;default:null"`
	User        User            `gorm:"foreignKey:UserId" json:"user"`
	ChannelName string          `gorm:"not null"`
	ReadStatus  bool            `gorm:"default:false"`
	ClientId    string          `gorm:"null"`
}
