package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {

	name := c.Param ("name")[1:]

	if name == "" {name="hello"}

	c.JSON(http.StatusOK, gin.H{
		"page":  "hello",
		"hello": name,
	})
}
