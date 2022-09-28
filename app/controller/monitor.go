package controller

import (
	"encoding/json"
	"gofun/pkg/tools"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rcrowley/go-metrics"
	stat "github.com/semihalev/gin-stats"
)

type GinStat struct {
	Db   DbStat
	Http metrics.Registry
}

type DbStat struct {
	MaxOpenConnections int64 `json:"MaxOpenConnections(最大连接数)"`
	OpenConnections    int64 `json:"OpenConnections(打开连接数)"`
	InUse              int   `json:"InUse(使用中的连接数)"`
	Idle               int   `json:"Idle(空闲连接数)"`
	WaitCount          int64 `json:"WaitCount(等待的连接总数)"`
	WaitDuration       int64 `json:"WaitDuration(等待新连接被阻止的总时间)"`
	MaxIdleClosed      int   `json:"MaxIdleClosed(由于达到设置的空闲连接池的最大数量而关闭的连接数)"`
	MaxIdleTimeClosed  int64 `json:"MaxIdleTimeClosed(由于达到设置的连接可空闲的最长时间而关闭的连接数)"`
	MaxLifetimeClosed  int64 `json:"MaxLifetimeClosed(由于达到设置的可重用连接的最长时间而关闭的连接数)"`
}

// 监控
func Stat(ctx *gin.Context) {
	token := ctx.DefaultQuery("token", "")
	var gs GinStat
	if token == "gofun_monitor" {
		gs.Http = stat.Report()
		info, _ := tools.Db("default").DB()
		statByte, err := json.Marshal(info.Stats())
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"status": 0, "msg": err.Error()})
			return
		}

		err = json.Unmarshal(statByte, &gs.Db)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"status": 0, "msg": err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"status": 1, "msg": "succ", "data": gs})
}
