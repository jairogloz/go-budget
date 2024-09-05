package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginCore "github.com/jairogloz/go-budget/cmd/api/core"
	accHandler "github.com/jairogloz/go-budget/cmd/api/handlers/account"
	transactionHandler "github.com/jairogloz/go-budget/cmd/api/handlers/transaction"
	"github.com/jairogloz/go-budget/cmd/api/middleware/auth"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"github.com/jairogloz/go-budget/pkg/domain/services/access_control"
	accService "github.com/jairogloz/go-budget/pkg/domain/services/account"
	"github.com/jairogloz/go-budget/pkg/domain/services/app_context"
	transactionService "github.com/jairogloz/go-budget/pkg/domain/services/transaction"
	"github.com/jairogloz/go-budget/pkg/mongo"
	"github.com/jairogloz/go-budget/pkg/mongo/account"
	"github.com/jairogloz/go-budget/pkg/mongo/category"
	"github.com/jairogloz/go-budget/pkg/mongo/transaction"
	"log"
	"time"
)

func main() {

	config, err := core.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %s", err.Error())
	}

	router := gin.Default()

	// Configure CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"}, // Adjust as needed
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	mongoClient, disconnectFunc, err := mongo.ConnectMongoDB(config.MongoURI)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer disconnectFunc()

	ctxHdl := app_context.NewHandler()

	txRepo := transaction.NewRepository(mongoClient)
	catRepo := category.NewRepository(mongoClient)
	txService := transactionService.NewService(txRepo, catRepo)
	txHandler := transactionHandler.NewHandler(txService)

	accountRepo := account.NewRepository(mongoClient)
	accountService := accService.NewService(accountRepo)
	accountHandler := accHandler.NewHandler(accountService, ctxHdl)

	server := ginCore.Server{
		AccountHdl:     accountHandler,
		AccountSrv:     accountService,
		Router:         router,
		TransactionHdl: txHandler,
		TxService:      txService,
	}

	accessCtrlService := access_control.NewService()
	authHdl := auth.NewHandler(accessCtrlService)

	router.Use(authHdl.AuthRequired())

	// ============= BACKEND ROUTES =============

	// Account routes
	server.Router.DELETE("/accounts/:id", server.AccountHdl.Delete)
	server.Router.GET("/accounts", server.AccountHdl.List)
	server.Router.GET("/accounts/:id", server.AccountHdl.GetById)
	server.Router.POST("/accounts", server.AccountHdl.Create)

	// Transaction routes
	server.Router.POST("/transactions", server.TransactionHdl.Insert)
	server.Router.DELETE("/transactions/:id", server.TransactionHdl.Delete)

	log.Fatalln(server.Router.Run(":8080"))

}
