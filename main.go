package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK!",
		})
	})

	if err := router.Run(":4000"); err != nil {
		panic(err)
	}
}
