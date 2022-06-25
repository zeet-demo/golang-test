package main

import (
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
		r.Run("127.0.0.1:8080")
	}()
}
