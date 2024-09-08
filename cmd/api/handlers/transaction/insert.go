package transaction

import (
	"github.com/gin-gonic/gin"
	pkgCore "github.com/jairogloz/go-budget/pkg/domain/core"
	"net/http"
)

// Insert inserts a new transaction into the database.
//
// @Summary Insert a new transaction
// @Description Insert a new transaction into the database
// @Tags transactions
// @Accept json
// @Produce json
// @Param transaction body core.TransactionCreateParams true "Transaction Create"
// @Success 200 {object} core.Transaction "OK"
// @Failure 400 {object} gin.H "Bad Request"
// @Failure 500 {object} gin.H "Internal Server Error"
// @Router /transactions [post]
func (h Handler) Insert(c *gin.Context) {
	user := c.Request.Context().Value(pkgCore.CtxKeyUser).(*pkgCore.User)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found in the context"})
		return
	}

	var txCreate pkgCore.TransactionCreateParams
	if err := c.ShouldBindJSON(&txCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTx, err := h.service.Insert(user, txCreate)
	if err != nil {
		// Todo: identify different types of errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newTx)

}
