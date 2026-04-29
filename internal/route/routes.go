package route

import (
	"github.com/ProArash/go-movie-reseller/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/test", handler.TestHandler)
}
