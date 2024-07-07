package account

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-budget/cmd/gin/core"
	"net/http"
)

// List lists the accounts for a given user.
func (h Handler) List(c *gin.Context) {

	// Retrieve the user ID from the context
	user, err := h.ctxHdl.GetUser(c.Request.Context())
	if err != nil {
		// todo: log error
		c.JSON(http.StatusInternalServerError, gin.H{"error": core.ErrMsgInternalServerError})
		return
	}
	
	accounts, err := h.service.List(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, accounts)
}
