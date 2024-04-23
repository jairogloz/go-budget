package main

import (
	"github.com/gin-gonic/gin"
	ginCore "github.com/jairogloz/go-budget/cmd/gin/core"
	transactionHandler "github.com/jairogloz/go-budget/cmd/gin/handlers/transaction"
	domainCore "github.com/jairogloz/go-budget/pkg/domain/core"
	transactionService "github.com/jairogloz/go-budget/pkg/domain/services/transaction"
	"github.com/jairogloz/go-budget/pkg/mongo"
	"github.com/jairogloz/go-budget/pkg/mongo/transaction"
	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
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

	server := ginCore.Server{
		Router:         router,
		TransactionHdl: txHandler,
	}

	// ============= BACKEND ROUTES =============
	server.Router.POST("/transactions", server.TransactionHdl.Insert)

	// ============= TEMPLATE ROUTES =============
	server.Router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	server.Router.GET("/accounts", func(c *gin.Context) {
		c.HTML(200, "accounts.tmpl", gin.H{
			"accounts": []domainCore.Account{
				{Name: "Savings", ID: "savings", CurrentBalance: decimal.NewFromFloat(100.0)},
				{Name: "Credit", ID: "credit", CurrentBalance: decimal.NewFromFloat(-100.0)},
			},
		})
	})

	server.Router.GET("/accounts/:id", func(c *gin.Context) {
		c.HTML(200, "account.tmpl", gin.H{
			"account": domainCore.Account{
				Name:           "Savings",
				CurrentBalance: decimal.NewFromInt32(100),
			},
			"transactions": []domainCore.Transaction{
				{ID: "1", Amount: decimal.NewFromFloat(100.0), Description: "Initial deposit", Category: strPtr("deposit")},
				{ID: "2", Amount: decimal.NewFromFloat(-10.0), Description: "Coquita", Category: strPtr("food")},
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
