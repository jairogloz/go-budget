package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	router := gin.Default()

	router.LoadHTMLGlob("pkg/templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	log.Fatalln(router.Run(":8080"))

}
