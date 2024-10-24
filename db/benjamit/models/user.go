package models

import (
	"bjm/utils/enums"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();unique" json:"uuid"`
	Username      string         `gorm:"unique;not null"`
	Password      string         `gorm:"not null"`
	PrefixId      int            `gorm:"not null" json:"prefix_id"`
	Prefix        Prefix         `gorm:"foreignKey:PrefixId" json:"prefix"`
	Firstname     string         `gorm:"not null"`
	Lastname      string         `gorm:"not null"`
	Nickname      string         `gorm:"default:null"`
	Birthday      time.Time      `gorm:"type:timestamp"`
	Email         string         `gorm:"unique;not null"`
	LineId        string         `gorm:"default:null"`
	MobilePhoneNo string         `gorm:"default:null"`
	HomePhoneNo   string         `gorm:"default:null"`
	Role          enums.RoleEnum `gorm:"not null;type:role_enum;default:'USER'"`
}
