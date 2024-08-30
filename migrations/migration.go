package migrations

import (
	"gorm.io/gorm"
	"01-xyz-finance/internal/accessapp"
	"01-xyz-finance/internal/customer"
	"01-xyz-finance/internal/customerlimit"
	"01-xyz-finance/internal/transaction"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&accessapp.AccessApp{},
		&customer.Customer{},
		&customerlimit.CustomerLimit{},
		&transaction.Transaction{},
	)
}
