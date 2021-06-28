package main

import (
	"net/http"
	"web-practice/controllers"
	"web-practice/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello"})
	})

	models.ConnectDtebase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/book", controllers.CreateBook)

	r.Run()
}
