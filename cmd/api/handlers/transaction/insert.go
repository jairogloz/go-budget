package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-budget/cmd/api/core"
	pkgCore "github.com/jairogloz/go-budget/pkg/domain/core"
	"net/http"
	"time"
)

// Insert inserts a new transaction into the database.
//
// @Summary Insert a new transaction
// @Description Insert a new transaction into the database
// @Tags transactions
// @Accept json
// @Produce json
// @Param transaction body core.TransactionCreate true "Transaction Create"
// @Success 201 "Created"
// @Failure 400 {object} gin.H "Bad Request"
// @Failure 500 {object} gin.H "Internal Server Error"
// @Router /transactions [post]
func (h Handler) Insert(c *gin.Context) {
	user := c.Request.Context().Value(pkgCore.CtxKeyUser).(*pkgCore.User)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found in the context"})
		return
	}

	var txCreate core.TransactionCreate
	if err := c.ShouldBindJSON(&txCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now().UTC()
	newTx := &pkgCore.Transaction{
		Amount:      txCreate.Amount,
		AccountId:   txCreate.AccountId,
		Category:    txCreate.Category,
		CreatedAt:   &now,
		Description: txCreate.Description,
		UpdatedAt:   &now,
		UserId:      user.ID,
	}

	err := h.service.Insert(newTx)
	if err != nil {
		// Todo: identify different types of errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)

}
