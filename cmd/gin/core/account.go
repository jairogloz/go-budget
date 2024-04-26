package core

import "github.com/gin-gonic/gin"

// AccountHandler exposes the handlers for the account domain.
type AccountHandler interface {
	List(c *gin.Context)
	GetById(c *gin.Context)
}
