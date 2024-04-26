package account

import "github.com/gin-gonic/gin"

func (h Handler) GetById(c *gin.Context) {
	userId := "1"

	id := c.Param("id")
	account, err := h.service.GetByID(userId, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, account)
}
