package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-budget/cmd/gin/core"
	core2 "github.com/jairogloz/go-budget/pkg/core"
	"github.com/shopspring/decimal"
	"log"
)

func main() {

	router := gin.Default()

	router.LoadHTMLGlob("pkg/templates/*")

	server := core.Server{
		Router: router,
	}

	server.Router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	server.Router.GET("/accounts", func(c *gin.Context) {
		c.HTML(200, "accounts.tmpl", gin.H{
			"accounts": []core2.Account{
				{Name: "Savings", ID: "savings", CurrentBalance: decimal.NewFromFloat(100.0)},
				{Name: "Credit", ID: "credit", CurrentBalance: decimal.NewFromFloat(-100.0)},
			},
		})
	})

	server.Router.GET("/accounts/:id", func(c *gin.Context) {
		c.HTML(200, "account.tmpl", gin.H{
			"account": core2.Account{
				Name:           "Savings",
				CurrentBalance: decimal.NewFromInt32(100),
			},
			"transactions": []core2.Transaction{
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
