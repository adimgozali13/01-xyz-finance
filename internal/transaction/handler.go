package transaction

import (
	"net/http"
	"strconv"
	"time"
	"math/rand"

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
	transactions, err := h.service.GetAllTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func (h *Handler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	transaction, err := h.service.GetTransactionByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func generateRandomString() string {
	rand.Seed(time.Now().UnixNano())
	randomNumbers := ""
	for i := 0; i < 12; i++ {
		randomNumbers += strconv.Itoa(rand.Intn(10)) 
	}
	return "XYZ-" + randomNumbers
}

func (h *Handler) Create(c *gin.Context) {
	var transaction Transaction
	if err := c.ShouldBind(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction.ContractNumber = generateRandomString()

	currentDate := time.Now()
	year, month, day := currentDate.Date()
	newMonth := int(month) + transaction.Term
	newYear := year
	if newMonth > 12 {
		newMonth = newMonth - 12
		newYear++
	}

	if newMonth <= 0 {
		newMonth = 12 + newMonth
		newYear--
	}

	billingDate := time.Date(newYear, time.Month(newMonth), day, currentDate.Hour(), currentDate.Minute(), currentDate.Second(), currentDate.Nanosecond(), currentDate.Location())
	otr := transaction.InstallmentAmount + transaction.AdminFee + ((transaction.InterestAmount / 100) * transaction.InstallmentAmount)

	limit,_ := h.service.GetLimitTerm(transaction.Term, transaction.CustomerID)

	if limit == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No loan limit assigned to this customer",
		})
		return
	}
	
	if otr > limit.Amount {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "OTR amount exceeds the loan limit.",
		})
		return
	}

	newAmount := float64(limit.Amount) - float64(otr)


	h.service.UpdateTermLimitAmount(transaction.Term, transaction.CustomerID, newAmount)


	transaction.CustomerLimitID = limit.ID
	transaction.OTR = otr
	transaction.BillingDate = billingDate
	transaction.Status = "Unpaid"

	if err := h.service.CreateTransaction(&transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}


func (h *Handler) PaidTransactionCustomer(c *gin.Context) {
	var transaction Transaction
	if err := c.ShouldBind(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cNumber := transaction.ContractNumber


	trx, climit, err := h.service.PaidTransactionAmount(cNumber)
	if err != nil {
		if err.Error() == "Customer has already paid" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transaction": trx,
		"customerLimit": climit,
	})
}

func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var transaction Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction.ID = uint(id)
	if err := h.service.UpdateTransaction(&transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteTransaction(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Transaction deleted"})
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	transactions := r.Group("/transactions")
	{
		transactions.GET("", handler.GetAll)
		transactions.GET("/:id", handler.GetByID)
		transactions.POST("/create", handler.Create)
		transactions.POST("/paid", handler.PaidTransactionCustomer)
		transactions.PUT("/:id", handler.Update)
		transactions.DELETE("/:id", handler.Delete)
	}
}
