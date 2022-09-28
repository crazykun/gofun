package route

import (
	"gofun/app/controller"
	"gofun/conf"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// 首页
	router.Any("/", controller.Index)

	// 心跳接口
	router.GET("/health-check", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome "+conf.Config.AppName)
	})
	// 慢日志测试
	router.GET("/slow", func(c *gin.Context) {
		time.Sleep(time.Duration(3) * time.Second)
		c.String(http.StatusOK, "slow")
	})
	// 错误日志测试
	router.GET("/err", func(context *gin.Context) {
		panic("error test")
	})
	// 监控
	router.GET("/monitor/stat", controller.Stat)
}
