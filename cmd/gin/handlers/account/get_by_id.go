package account

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-budget/cmd/gin/core"
	"net/http"
)

func (h Handler) GetById(c *gin.Context) {
	// Retrieve the user ID from the context
	user, err := h.ctxHdl.GetUser(c.Request.Context())
	if err != nil {
		// todo: log error
		c.JSON(http.StatusInternalServerError, gin.H{"error": core.ErrMsgInternalServerError})
		return
	}

	accountID := c.Param("id")
	account, err := h.service.GetByID(user.ID, accountID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, account)
}
