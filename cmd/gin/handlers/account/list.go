package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// List lists the accounts for a given user.
func (h Handler) List(c *gin.Context) {

	// Todo: replace userId with the actual user id.
	userId := "1"
	accounts, err := h.service.List(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, accounts)
}
