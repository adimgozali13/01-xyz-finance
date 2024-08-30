package customerlimit

import (
	"net/http"
	"strconv"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetAll(c *gin.Context) {
	customerLimits, err := h.service.GetAllCustomerLimits()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data" : customerLimits,
		"succes" : true,
	})
}

func (h *Handler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	customerLimit, err := h.service.GetCustomerLimitByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer limit not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data" : customerLimit})
}

func (h *Handler) Create(c *gin.Context) {
	var customerLimit CustomerLimit

	if err := c.ShouldBind(&customerLimit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingLimit, _ := h.service.GetCustomerLimitByTerm(customerLimit.Term, customerLimit.CustomerID)
	
	if existingLimit != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Customer limit with a term of %d months already exists", customerLimit.Term),
		})
		return
	}

	if err := h.service.CreateCustomerLimit(&customerLimit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, customerLimit)
}

func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var customerLimit CustomerLimit
	if err := c.ShouldBindJSON(&customerLimit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customerLimit.ID = uint(id)
	if err := h.service.UpdateCustomerLimit(&customerLimit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customerLimit)
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteCustomerLimit(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Customer limit deleted"})
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	customerLimits := r.Group("/customer-limits")
	{
		customerLimits.GET("", handler.GetAll)
		customerLimits.GET("/:id", handler.GetByID)
		customerLimits.POST("/create", handler.Create)
		customerLimits.PUT("/:id", handler.Update)
		customerLimits.DELETE("/:id", handler.Delete)
	}
}
