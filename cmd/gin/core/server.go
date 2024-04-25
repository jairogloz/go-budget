package core

import "github.com/gin-gonic/gin"

// Server is the core server struct that holds the required components to serve
// the application using gin.
type Server struct {
	AccountHdl     AccountHandler
	Router         *gin.Engine
	TransactionHdl TransactionHandler
}
