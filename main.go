package main

import (
	// "gorm_api/db"
	"gorm_api/controller"
	"gorm_api/models"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting server.......")

	r := gin.Default()
	models.CDB()
	r.GET("/books", controller.Index)
	r.GET("/books/:Id", controller.Show)
	r.POST("/books", controller.Create)
	r.PUT("/books/:Id", controller.Update)
	r.DELETE("/books", controller.Delete)

	//FrontEnd
	r.LoadHTMLGlob("views/*") // Set the location of HTML templates

	//render read data
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	//render add data
	r.GET("/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create.html", gin.H{})
	})

	//render update data
	r.GET("/update", func(c *gin.Context) {
		c.HTML(http.StatusOK, "update.html", gin.H{})
	})

	r.Run()
}
