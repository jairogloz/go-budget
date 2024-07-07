package account

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-budget/cmd/gin/core"
	"net/http"
)

// Delete deletes an account.
func (h Handler) Delete(c *gin.Context) {

	// Retrieve the user ID from the context
	user, err := h.ctxHdl.GetUser(c.Request.Context())
	if err != nil {
		// todo: log error
		c.JSON(http.StatusInternalServerError, gin.H{"error": core.ErrMsgInternalServerError})
		return
	}

	accountID := c.Param("id")

	if err := h.service.Delete(user.ID, accountID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, nil)
}
