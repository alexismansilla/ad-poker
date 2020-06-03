package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Hi",
		})
	})
	router.Run(":3000")

	fmt.Println("Holamundo")
}
