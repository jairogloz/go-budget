package main

import (
	"github.com/gin-gonic/gin"
	ginCore "github.com/jairogloz/go-budget/cmd/gin/core"
	accHandler "github.com/jairogloz/go-budget/cmd/gin/handlers/account"
	transactionHandler "github.com/jairogloz/go-budget/cmd/gin/handlers/transaction"
	domainCore "github.com/jairogloz/go-budget/pkg/domain/core"
	accService "github.com/jairogloz/go-budget/pkg/domain/services/account"
	transactionService "github.com/jairogloz/go-budget/pkg/domain/services/transaction"
	"github.com/jairogloz/go-budget/pkg/mongo"
	"github.com/jairogloz/go-budget/pkg/mongo/account"
	"github.com/jairogloz/go-budget/pkg/mongo/transaction"
	"github.com/joho/godotenv"
	"log"
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
	txService := transactionService.NewService(txRepo)
	txHandler := transactionHandler.NewHandler(txService)

	accountRepo := account.NewRepository(mongoClient)
	accountService := accService.NewService(accountRepo)
	accountHandler := accHandler.NewHandler(accountService)

	server := ginCore.Server{
		AccountHdl:     accountHandler,
		Router:         router,
		TransactionHdl: txHandler,
	}

	// ============= BACKEND ROUTES =============
	server.Router.GET("/accounts", server.AccountHdl.List)
	server.Router.GET("/accounts/:id", server.AccountHdl.GetById)
	server.Router.POST("/transactions", server.TransactionHdl.Insert)

	// ============= TEMPLATE ROUTES =============
	server.Router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	server.Router.GET("/my-accounts", func(c *gin.Context) {
		c.HTML(200, "accounts.tmpl", gin.H{
			"accounts": []domainCore.Account{
				{Name: "Savings", ID: "savings", Balance: 100.0},
				{Name: "Credit", ID: "credit", Balance: -100.0},
			},
		})
	})

	server.Router.GET("/my-accounts/:id", func(c *gin.Context) {
		c.HTML(200, "account.tmpl", gin.H{
			"account": domainCore.Account{
				Name:    "Savings",
				Balance: 100,
			},
			"transactions": []domainCore.Transaction{
				{ID: "1", Amount: 100.0, Description: "Initial deposit", Category: strPtr("deposit")},
				{ID: "2", Amount: -10.0, Description: "Coquita", Category: strPtr("food")},
			},
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
