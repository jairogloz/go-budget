package transaction

import (
	"github.com/gin-gonic/gin"
	pkgCore "github.com/jairogloz/go-budget/pkg/domain/core"
	"net/http"
)

// List godoc
// @Summary List transactions
// @Description Get a list of transactions for a user within a date range
// @Tags transactions
// @Param from query string true "Start date in ISO 8601 format"
// @Param to query string true "End date in ISO 8601 format"
// @Param limit query int false "Limit the number of results"
// @Param offset query int false "Offset for pagination"
// @Param list_type query string false "Type of list"
// @Success 200 {object} []pkgCore.Transaction
// @Failure 400 {object} gin.H{"error": "from is required"}
// @Failure 400 {object} gin.H{"error": "to is required"}
// @Failure 500 {object} gin.H{"error": "user not found in the context"}
// @Failure 500 {object} gin.H{"error": "internal server error"}
// @Router /transactions [get]
func (h Handler) List(c *gin.Context) {
	user := c.Request.Context().Value(pkgCore.CtxKeyUser).(*pkgCore.User)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found in the context"})
		return
	}

	from := c.Query("from")
	if from == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "from is required"})
		return
	}
	to := c.Query("to")
	if to == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "to is required"})
		return
	}
	limit := c.Query("limit")
	if limit == "" {
		limit = "10"
	}
	offset := c.Query("offset")
	if offset == "" {
		offset = "0"
	}
	listType := c.Query("list_type")

	transactions, err := h.service.List(c.Request.Context(), user, from, to, limit, offset, listType)
	if err != nil {
		// Todo: identify different types of errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
