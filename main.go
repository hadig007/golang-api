package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "Hello World",
	})
}

func main() {
	fmt.Println("Hello World")

	r := gin.Default()
	r.GET("/", getHome)

	r.Run()
}
