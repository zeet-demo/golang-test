package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.gohtml", gin.H{
		"page": "home",
	})
}
