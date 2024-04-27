package account

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-budget/cmd/gin/core"
	"net/http"
)

// Delete deletes an account.
func (h Handler) Delete(c *gin.Context) {

	// Retrieve the user ID from the context
	userID := c.Request.Context().Value(core.UserIDKey).(string)
	if userID == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user ID not found in the context"})
		return
	}
	id := c.Param("id")

	if err := h.service.Delete(userID, id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, nil)
}
