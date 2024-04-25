package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-budget/cmd/gin/core"
	"net/http"
)

// Insert inserts a new transaction into the database.
func (h Handler) Insert(c *gin.Context) {
	var transactionCreate core.TransactionCreate
	if err := c.ShouldBindJSON(&transactionCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isNewCategory := false
	if transactionCreate.Category != nil {
		isNewCategory = transactionCreate.Category.IsNew
	}

	domainTx := transactionCreate.ToDomain()
	err := h.service.Insert(domainTx, isNewCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, domainTx)
}