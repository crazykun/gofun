package route

import (
	"gofun/app/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", controller.Index)
	// router.GET("/test", controller.Test)
	router.GET("/health-check", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gofun Server")
	})
}
