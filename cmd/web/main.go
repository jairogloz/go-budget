package main

import (
	"github.com/gin-gonic/gin"
	ginCore "github.com/jairogloz/go-budget/cmd/api/core"
	accHandler "github.com/jairogloz/go-budget/cmd/api/handlers/account"
	transactionHandler "github.com/jairogloz/go-budget/cmd/api/handlers/transaction"
	"github.com/jairogloz/go-budget/cmd/api/middleware/auth"
	"github.com/jairogloz/go-budget/cmd/web/handlers/transactions"
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
	"net/http"
)

func main() {

	config, err := core.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %s", err.Error())
	}

	router := gin.Default()

	router.LoadHTMLGlob("cmd/web/templates/*")

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

	// ============= Transactions =============
	txHdl := transactions.TransactionHandler{}
	server.Router.GET("/transactions", txHdl.Transactions)

	server.Router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	server.Router.GET("/my-accounts", func(c *gin.Context) {

		// Retrieve the user ID from the context
		userID := c.Request.Context().Value(core.CtxKeyUser).(string)
		if userID == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user ID not found in the context"})
			return
		}

		accounts, err := server.AccountSrv.List(userID)
		if err != nil {
			log.Printf("Error getting accounts: %v", err)
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}

		var totalBalance float64
		for _, a := range accounts {
			totalBalance += a.Balance
		}

		c.HTML(200, "accounts.tmpl", gin.H{
			"accounts":     accounts,
			"totalBalance": totalBalance,
		})
	})

	server.Router.GET("/my-accounts/:id", func(c *gin.Context) {

		// Retrieve the user ID from the context
		userID := c.Request.Context().Value(core.CtxKeyUser).(string)
		if userID == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user ID not found in the context"})
			return
		}

		// get id from path
		id := c.Param("id")

		queriedAccount, err := server.AccountSrv.GetByID(userID, id)
		if err != nil {
			log.Printf("Error getting account: %v", err)
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}

		txs, err := server.TxService.FindByAccountID(userID, id)
		if err != nil {
			log.Printf("Error getting transactions: %v", err)
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}

		c.HTML(200, "account.tmpl", gin.H{
			"account":      queriedAccount,
			"transactions": txs,
		})
	})

	log.Fatalln(server.Router.Run(":8081"))
}
