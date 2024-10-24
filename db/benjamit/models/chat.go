package models

import (
	"bjm/utils/enums"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	UUID        uuid.UUID             `gorm:"type:uuid;default:uuid_generate_v4();unique" json:"uuid"`
	Message     string                `gorm:"not null"`
	MessageType enums.MessageTypeEnum `gorm:"not null;type:message_type_enum;default:'TEXT'"`
	UserId      int                   `gorm:"null;default:null"`
	User        User                  `gorm:"foreignKey:UserId" json:"user"`
	ChannelName string                `gorm:"not null"`
	ReadStatus  bool                  `gorm:"default:false"`
	ClientId    string                `gorm:"null;default:null"`
}
