package core

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-budget/pkg/domain/core"
)

// Server is the core server struct that holds the required components to serve
// the application using gin.
type Server struct {
	AccountHdl     AccountHandler
	AccountSrv     core.AccountService
	Router         *gin.Engine
	TransactionHdl TransactionHandler
}
