package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionTypeEnum string
type EntrySourceEnum string

const (
	DEBIT  TransactionTypeEnum = "DEBIT"
	CREDIT TransactionTypeEnum = "CREDIT"
)

const (
	SYSTEM EntrySourceEnum = "SYSTEM"
	MANUAL EntrySourceEnum = "MANUAL"
)

type IncomeAndExpense struct {
	gorm.Model
	UUID                  uuid.UUID           `gorm:"type:uuid;default:uuid_generate_v4();unique" json:"uuid"`
	Amount                float64             `gorm:"not null"`                                          // จำนวนเงิน
	Description           string              `gorm:"null"`                                              // รายละเอียดของรายการ
	TransactionDate       time.Time           `gorm:"not null;type:timestamp;default:CURRENT_TIMESTAMP"` // วันที่ทำรายการ
	TransactionType       TransactionTypeEnum `gorm:"not null;type:transaction_type_enum"`               // ประเภท (debit หรือ credit)
	EntrySource           EntrySourceEnum     `gorm:"not null;type:entry_source_enum;default:'MANUAL'"`
	UserId                int                 `gorm:"not null" json:"user_id"`                  // ID ผู้ใช้ (ถ้ามี)
	User                  User                `gorm:"foreignKey:UserId" json:"user"`            // ความสัมพันธ์กับ User (ถ้ามี)
	ReferProductId        int                 `gorm:"null"`                                     // ID สินค้า (ถ้ามี)
	Product               Product             `gorm:"foreignKey:ReferProductId" json:"product"` // ความสัมพันธ์กับ Product (ถ้ามี)
	ReferProductSellingId int                 `gorm:"null"`
	ProductSelling        ProductSelling      `gorm:"foreignKey:ReferProductSellingId" json:"product_selling"`
}
