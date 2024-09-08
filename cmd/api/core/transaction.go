package core

import (
	"github.com/gin-gonic/gin"
)

// TransactionHandler exposes the handlers for the transactions services.
type TransactionHandler interface {
	Delete(c *gin.Context)
	Insert(c *gin.Context)
}
