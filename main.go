package main

import (
	"os"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zeet-demo/golang-test/handler"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", handler.Home)
	r.GET("/hello/*name", handler.Hello)

	func() {
		log.Println("Server started")
		r.Run(":" + os.Getenv("PORT"))
	}()
}
