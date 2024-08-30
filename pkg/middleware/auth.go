package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"01-xyz-finance/internal/accessapp"
)

func APIKeyAuth(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("API-Key")
		domain := c.Request.Host

		if apiKey == "" || domain == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "API key or domain missing"})
			c.Abort()
			return
		}

		var access accessapp.AccessApp
		if err := db.Where("api_key = ? AND domain = ?", apiKey, domain).First(&access).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key or domain"})
			c.Abort()
			return
		}

		c.Next()
	}
}
