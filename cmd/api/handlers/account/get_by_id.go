package account

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-budget/cmd/api/core"
	"net/http"
)

// GetById retrieves an account by its ID
// @Summary Get account by ID
// @Description Retrieve an account by its ID for the authenticated user
// @Tags accounts
// @Param id path string true "Account ID"
// @Success 200 {object} core.Account
// @Failure 500 {object} gin.H{"error": string}
// @Router /accounts/{id} [get]
func (h Handler) GetById(c *gin.Context) {
	// Retrieve the user ID from the context
	user, err := h.ctxHdl.GetUser(c.Request.Context())
	if err != nil {
		// todo: log error
		c.JSON(http.StatusInternalServerError, gin.H{"error": core.ErrMsgInternalServerError})
		return
	}

	accountID := c.Param("id")
	account, err := h.service.GetByID(c.Request.Context(), user.ID, accountID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, account)
}
