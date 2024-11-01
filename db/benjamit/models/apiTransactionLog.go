package models

import (
	"bjm/utils/enums"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ApiTransactionLog struct {
	gorm.Model
	UUID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();unique" json:"uuid"`
	Path          string    `gorm:"not null"`
	Method        string    `gorm:"not null"`
	ContentType   string
	RequestBody   string
	StatusCode    int
	ResponseBody  string
	Headers       string
	Origin        string
	InterfaceType enums.InterfaceTypeEnum `gorm:"default:HTTP;not null;type:interface_type_enum"`
}
