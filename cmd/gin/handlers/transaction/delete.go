package transaction

import (
	"github.com/gin-gonic/gin"
	pkgCore "github.com/jairogloz/go-budget/pkg/domain/core"
	"net/http"
)

// Delete removes a transaction from the database.
func (h Handler) Delete(c *gin.Context) {
	// Retrieve the user ID from the context
	userID := c.Request.Context().Value(pkgCore.CtxKeyUser).(string)
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
