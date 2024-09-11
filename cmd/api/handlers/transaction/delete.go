package transaction

import (
	"github.com/gin-gonic/gin"
	pkgCore "github.com/jairogloz/go-budget/pkg/domain/core"
	"net/http"
)

// Delete removes a transaction from the database.
//
// @Summary Delete a transaction
// @Description Delete a transaction from the database
// @Tags transactions
// @Param id path string true "Transaction ID"
// @Success 204 "No Content"
// @Failure 500 {object} gin.H "Internal Server Error"
// @Router /transactions/{id} [delete]
func (h Handler) Delete(c *gin.Context) {
	// Retrieve the user ID from the context
	user := c.Request.Context().Value(pkgCore.CtxKeyUser).(*pkgCore.User)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user ID not found in the context"})
		return
	}
	id := c.Param("id")

	if err := h.service.Delete(user.ID, id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, nil)
}
