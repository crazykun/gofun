package middleware

// 路由优先级熔断式中间件

import (
	"fmt"
	"gofun/pkg/log"
	"gofun/pkg/tools"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RankingLimiter 限制接口请求数据，做到主动降载，超阀值会自动熔断
// 优先级数字越大，优先级越低，规定：40代表cpu使用率超40%时拦截；95代表cpu使用率超过95%时拦截；250%代表cpu使用率超过250%时拦截。
// 数字最低是=40；最高=100*核心数（是否会因为多核超过100%没测过，暂时定为100%上限）；0表示不做任何熔断。
func RankingLimiter(ranking int64) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 请求链接信息
		// 读取每次请求的请求全局参数
		_host, _ := ctx.Get("host")

		// 读取超全局变量即可
		var _cpuNum interface{} = tools.GetGlobalData("cpu_num")
		var cpuNum int64 = _cpuNum.(int64)
		var _cpuPercent interface{} = tools.GetGlobalData("cpu_percent")
		var __cpuPercent float64 = _cpuPercent.(float64)
		var cpuPercent int64 = int64(math.Floor(__cpuPercent))

		var alertRanking string = "[80, 100]"
		//if ranking >= 80 && ranking <= 100*cpuNum { // 正确范围
		if ranking >= 80 && ranking <= 100 { // 正确范围
			ctx.Next()
		} else if ranking == 0 { // 0不做任何熔断操作
			ranking = 10000000
			ctx.Next()
		} else if ranking < 80 { // 使用默认值
			ranking = 80
			ctx.Next()
		} else {
			fmt.Println("优先级范围错误。"+alertRanking+" => ", cpuPercent)
			ctx.JSON(http.StatusOK, gin.H{
				"status":  0,
				"message": "优先级范围错误",
				"data": gin.H{
					"Route-Ranking": ranking,
					"CPU-Num":       cpuNum,
				},
			})

			ctx.Abort()
		}

		// 直接熔断，等待x秒定时周期后看CPU占用率是否恢复
		if cpuPercent > ranking {
			template := `{"name":"rankingLimiter","uri":"%s", "ip":"%s", "result": "达到熔断标准"}`
			template = fmt.Sprintf(template, _host, ctx.ClientIP())
			log.Log(template)

			ctx.JSON(http.StatusTooManyRequests, gin.H{
				"status":  429,
				"message": "通道拥挤，请稍后再试",
				"data": gin.H{
					"Route-Ranking": ranking,
					"CPU-Percent":   cpuPercent,
				},
			})
			ctx.Abort()
		} else {
			ctx.Next()
		}

	}
}
