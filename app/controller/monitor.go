package controller

import (
	"gofun/pkg/tools"
	"net/http"

	"github.com/gin-gonic/gin"
	stat "github.com/semihalev/gin-stats"
)

// 数据库监控
func Db(ctx *gin.Context) {
	token := ctx.DefaultQuery("token", "")
	var data interface{}
	if token == "gofun_monitor" {
		info, _ := tools.Db("default").DB()
		data = info.Stats()
	}

	ctx.JSON(http.StatusOK, gin.H{"status": 1, "msg": "succ", "db": data})
}

// 访问统计
func Stat(ctx *gin.Context) {
	token := ctx.DefaultQuery("token", "")
	var gin interface{}
	if token == "gofun_monitor" {
		gin = stat.Report()
	}
	ctx.JSON(http.StatusOK, gin)
}
