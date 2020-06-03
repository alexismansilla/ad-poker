package main

import (
	"github.com/foolin/gin-template"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
}

func main() {
	router := gin.Default()

	//new template engine
	router.HTMLRender = gintemplate.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "Hi"})
	})

	router.GET("/index", func(c *gin.Context) {
		c.HTML(200, "templates/index.html", gin.H{"title": "Hola"})
	})
	router.Run(":3000")
}
