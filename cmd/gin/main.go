package main

import (
	"github.com/gin-gonic/gin"
	ginCore "github.com/jairogloz/go-budget/cmd/gin/core"
	accHandler "github.com/jairogloz/go-budget/cmd/gin/handlers/account"
	transactionHandler "github.com/jairogloz/go-budget/cmd/gin/handlers/transaction"
	"github.com/jairogloz/go-budget/cmd/gin/middleware/auth"
	accService "github.com/jairogloz/go-budget/pkg/domain/services/account"
	transactionService "github.com/jairogloz/go-budget/pkg/domain/services/transaction"
	"github.com/jairogloz/go-budget/pkg/mongo"
	"github.com/jairogloz/go-budget/pkg/mongo/account"
	"github.com/jairogloz/go-budget/pkg/mongo/category"
	"github.com/jairogloz/go-budget/pkg/mongo/transaction"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	router.LoadHTMLGlob("pkg/templates/*")

	mongoClient, disconnectFunc, err := mongo.ConnectMongoDB(os.Getenv("MONGO_URI"))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer disconnectFunc()

	txRepo := transaction.NewRepository(mongoClient)
	catRepo := category.NewRepository(mongoClient)
	txService := transactionService.NewService(txRepo, catRepo)
	txHandler := transactionHandler.NewHandler(txService)

	accountRepo := account.NewRepository(mongoClient)
	accountService := accService.NewService(accountRepo)
	accountHandler := accHandler.NewHandler(accountService)

	server := ginCore.Server{
		AccountHdl:     accountHandler,
		AccountSrv:     accountService,
		Router:         router,
		TransactionHdl: txHandler,
		TxService:      txService,
	}

	// ============= BACKEND ROUTES =============

	// Account routes
	server.Router.DELETE("/accounts/:id", auth.AuthRequired(), server.AccountHdl.Delete)
	server.Router.GET("/accounts", server.AccountHdl.List)
	server.Router.GET("/accounts/:id", server.AccountHdl.GetById)
	server.Router.POST("/accounts", auth.AuthRequired(), server.AccountHdl.Create)

	// Transaction routes
	server.Router.POST("/transactions", auth.AuthRequired(), server.TransactionHdl.Insert)
	server.Router.DELETE("/transactions/:id", auth.AuthRequired(), server.TransactionHdl.Delete)

	// ============= TEMPLATE ROUTES =============
	server.Router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	server.Router.GET("/my-accounts", auth.AuthRequired(), func(c *gin.Context) {

		// Retrieve the user ID from the context
		userID := c.Request.Context().Value(ginCore.UserIDKey).(string)
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

	server.Router.GET("/my-accounts/:id", auth.AuthRequired(), func(c *gin.Context) {

		// Retrieve the user ID from the context
		userID := c.Request.Context().Value(ginCore.UserIDKey).(string)
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

	log.Fatalln(server.Router.Run(":8080"))

}

func strPtr(s string) *string {
	return &s
}

//if useHTML {
//c.HTML(http.StatusOK, "index.tmpl", data)
//} else {
//c.JSON(http.StatusOK, data)
//}
