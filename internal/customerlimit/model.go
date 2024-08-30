package customerlimit

import (
	"time"
)

type CustomerLimit struct {
	ID         uint      `gorm:"primaryKey"`
	CustomerID uint      `gorm:"not null"`
	Term       int       
	Amount     float64   `gorm:"type:decimal(10,2);not null"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	CreatedAt  time.Time `gorm:"autoCreateTime"` 
}
