package transaction

import (
	"time"
	"01-xyz-finance/internal/customerlimit"
)

type Transaction struct {
	ID               uint                          `gorm:"primaryKey"`
	CustomerID         uint           			   `gorm:"not null"`
	CustomerLimitID    uint   					  `gorm:"not null"`
	ContractNumber   string                          
	OTR              float64                      
	AdminFee         float64
	InstallmentAmount float64
	InterestAmount   float64
	AssetName        string
	Status        string
	Term        int
	BillingDate        time.Time
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	CreatedAt  time.Time `gorm:"autoCreateTime"` 
	CustomerLimit    customerlimit.CustomerLimit `gorm:"foreignKey:CustomerLimitID" json:"customer_limit,omitempty"`
}
