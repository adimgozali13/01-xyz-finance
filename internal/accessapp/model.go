package accessapp

import (
	"time"
)

type AccessApp struct {
	ID        uint   `gorm:"primaryKey"`
	Domain    string `gorm:"unique"`
	ApiKey    string
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	CreatedAt  time.Time `gorm:"autoCreateTime"` 
}
