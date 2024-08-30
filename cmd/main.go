package main

import (
	"github.com/gin-gonic/gin"
	"01-xyz-finance/config"
	"01-xyz-finance/pkg/database"
	"01-xyz-finance/pkg/middleware"
	"01-xyz-finance/migrations"
	"01-xyz-finance/internal/customer"
	"01-xyz-finance/internal/customerlimit"
	"01-xyz-finance/internal/transaction"
)

func main() {
	cfg := config.LoadConfig()
	db := database.ConnectDB(cfg)

	// Run migrations
	migrations.RunMigrations(db)

	r := gin.Default()
	r.Use(middleware.APIKeyAuth(db))

	// Initialize routes
	customer.RegisterRoutes(r, db)
	customerlimit.RegisterRoutes(r, db)
	transaction.RegisterRoutes(r, db)

	r.Run(":8080")
}
