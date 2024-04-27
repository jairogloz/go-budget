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
	if transactionCreate.Category != nil && transactionCreate.Category.IsNew != nil {
		isNewCategory = *transactionCreate.Category.IsNew
	}

	// Todo: take userId from authentication
	domainTx := transactionCreate.ToDomain("1")
	updatedAccount, err := h.service.Insert(domainTx, isNewCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	txCreateResponse := core.TransactionCreateResponse{
		Transaction: *domainTx,
		Account:     *updatedAccount,
	}

	c.JSON(http.StatusOK, txCreateResponse)
}
