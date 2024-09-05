package transactions

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"net/http"
)

func (h TransactionHandler) Transactions(c *gin.Context) {

	user := c.Request.Context().Value(core.CtxKeyUser).(*core.User)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found in the context"})
		return
	}

	// List latest transactions

	c.HTML(http.StatusOK, "transactions.tmpl", gin.H{})
}
