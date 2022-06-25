package main

import (
	"os"
	"log"
	"html/template"
	
	"github.com/gin-gonic/gin"
	"github.com/zeet-demo/golang-test/handler"
)

var tmpl *template.Template

func main() {
	r := gin.New()
	r.Use (gin.Logger())
	r.Use (gin.Recovery())
	r.LoadHTMLGlob("templates/*.gohtml")

	
	r.GET("/", handler.Home)
	r.GET("/hello/*name", handler.Hello)

	func() {
		log.Println("Server started")
		r.Run(":" + os.Getenv("PORT"))
		// r.Run (":8080")
	}()
}
