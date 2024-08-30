package customer

import (
	"time"
	"01-xyz-finance/internal/customerlimit"
	"01-xyz-finance/internal/transaction"
)
	


type Customer struct {
	ID           uint      `gorm:"primaryKey"`
	NIK          string    `gorm:"type:varchar(255)" form:"NIK" json:"NIK"`
	FullName     string    `gorm:"type:varchar(255)" form:"FullName" json:"FullName"`
	LegalName    string    `gorm:"type:varchar(255)" form:"LegalName" json:"LegalName"`
	PlaceOfBirth string    `gorm:"type:varchar(255)" form:"PlaceOfBirth" json:"PlaceOfBirth"`
	DateOfBirth  time.Time    `gorm:"type:date" form:"DateOfBirth" json:"DateOfBirth" time_format:"2006-01-02"`
	Salary       float64   `form:"Salary" json:"Salary"`
	IDCardPhoto  string    `gorm:"type:varchar(255)" form:"IDCardPhoto" json:"IDCardPhoto"`
	SelfiePhoto  string    `gorm:"type:varchar(255)" form:"SelfiePhoto" json:"SelfiePhoto"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	CreatedAt  time.Time `gorm:"autoCreateTime"` 
	CustomerLimit []customerlimit.CustomerLimit `gorm:"foreignKey:CustomerID" json:"customer_limit,omitempty"`
	Transaction  []transaction.Transaction `gorm:"foreignKey:CustomerID" json:"transaction,omitempty"`
}

