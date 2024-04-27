package account

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-budget/cmd/gin/core"
	"net/http"
)

// Create creates a new account.
func (h Handler) Create(c *gin.Context) {

	// Retrieve the user ID from the context
	userID := c.Request.Context().Value(core.UserIDKey).(string)
	if userID == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user ID not found in the context"})
		return
	}

	var accountCreate core.AccountCreate
	if err := c.ShouldBindJSON(&accountCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainAccount := accountCreate.ToDomain(userID)
	err := h.service.Create(domainAccount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, domainAccount)

}
