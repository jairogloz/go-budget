package account

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-budget/cmd/api/core"
	"net/http"
)

// Create creates a new account.
func (h Handler) Create(c *gin.Context) {

	// Retrieve the user ID from the context
	user, err := h.ctxHdl.GetUser(c.Request.Context())
	if err != nil {
		// todo: log error
		c.JSON(http.StatusInternalServerError, gin.H{"error": core.ErrMsgInternalServerError})
		return
	}

	var accountCreate core.AccountCreate
	if err := c.ShouldBindJSON(&accountCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainAccount := accountCreate.ToDomain(user.ID)
	createdAccount, err := h.service.Create(user, *domainAccount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdAccount)

}
