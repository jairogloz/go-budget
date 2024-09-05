package core

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-budget/pkg/domain/ports"
)

// Server is the core server struct that holds the required components to serve
// the application using gin.
type Server struct {
	AccountHdl     AccountHandler
	AccountSrv     ports.AccountService
	Router         *gin.Engine
	TransactionHdl TransactionHandler
	TxService      ports.TransactionService
}
