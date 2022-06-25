package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {

	name := c.Param ("name")[1:]
	
	c.HTML(http.StatusOK, "hello.gohtml", gin.H{
		"page": "home",
		"name": name,
	})
}
