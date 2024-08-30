package customer

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
	customers, err := h.service.GetAllCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success" : true,
		"data" : customers,
	})
}

func (h *Handler) GetAllWithLimitCust(c *gin.Context) {
	customers, err := h.service.GetAllCustomersWithLimitCust()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}

func (h *Handler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	customer, err := h.service.GetCustomerByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, customer)
}



func (h *Handler) Create(c *gin.Context) {
	var customer Customer
	var errorMessage []string

	c.ShouldBind(&customer);


	existingCustomer, _ := h.service.GetCustomerByNik(customer.NIK)
	if existingCustomer != nil {
		errorMessage = append(errorMessage, fmt.Sprintf("NIK %s already exists ", existingCustomer.NIK))
	}

	fileKtp, err := c.FormFile("IDCardPhoto")
	if err != nil {
		errorMessage = append(errorMessage, "Photo KTP is required")
	}

	fileSelfie, err := c.FormFile("SelfiePhoto")
	if err != nil {
		errorMessage = append(errorMessage, "Photo Selfie is required")
	}

	if len(errorMessage) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	filePathKtp := fmt.Sprintf("uploads/KTP/%s", fileKtp.Filename)
	filePathSelfie := fmt.Sprintf("uploads/Selfie/%s", fileSelfie.Filename)

	if err := c.SaveUploadedFile(fileKtp, filePathKtp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save KTP photo"})
		return
	}

	if err := c.SaveUploadedFile(fileSelfie, filePathSelfie); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save Selfie photo"})
		return
	}

	customer.IDCardPhoto = filePathKtp
	customer.SelfiePhoto = filePathSelfie

	if err := h.service.CreateCustomer(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	c.JSON(http.StatusCreated, customer)
}




func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var customer Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer.ID = uint(id)
	if err := h.service.UpdateCustomer(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteCustomer(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Customer deleted"})
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	customers := r.Group("/customers")
	{
		customers.GET("", handler.GetAll)
		customers.GET("/:id", handler.GetByID)
		customers.POST("/create", handler.Create)
		customers.PUT("/:id", handler.Update)
		customers.DELETE("/:id", handler.Delete)
	}
}
