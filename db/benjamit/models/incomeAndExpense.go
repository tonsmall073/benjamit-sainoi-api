package models

import (
	"bjm/utils/enums"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IncomeAndExpense struct {
	gorm.Model
	UUID                  uuid.UUID                 `gorm:"type:uuid;default:uuid_generate_v4();unique" json:"uuid"`
	Amount                float64                   `gorm:"not null"`                                          // จำนวนเงิน
	Description           string                    `gorm:"null;default:null"`                                 // รายละเอียดของรายการ
	TransactionDate       time.Time                 `gorm:"not null;type:timestamp;default:CURRENT_TIMESTAMP"` // วันที่ทำรายการ
	TransactionType       enums.TransactionTypeEnum `gorm:"not null;type:transaction_type_enum"`               // ประเภท (debit หรือ credit)
	EntrySource           enums.EntrySourceEnum     `gorm:"not null;type:entry_source_enum;default:'MANUAL'"`
	UserId                int                       `gorm:"not null" json:"user_id"`                  // ID ผู้ใช้ (ถ้ามี)
	User                  User                      `gorm:"foreignKey:UserId" json:"user"`            // ความสัมพันธ์กับ User (ถ้ามี)
	ReferProductId        int                       `gorm:"null;default:null"`                        // ID สินค้า (ถ้ามี)
	Product               Product                   `gorm:"foreignKey:ReferProductId" json:"product"` // ความสัมพันธ์กับ Product (ถ้ามี)
	ReferProductSellingId int                       `gorm:"null;default:null"`
	ProductSelling        ProductSelling            `gorm:"foreignKey:ReferProductSellingId" json:"product_selling"`
	Quantity              int                       `gorm:"null;default:null"`
}
