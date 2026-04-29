package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  c.Request.URL.Path,
		"addr": c.ClientIP(),
	})
}
