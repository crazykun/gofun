package middleware

import (
	"gofun/pkg/logs"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// HttpLimiter 拦截http请求频率，请求限流
func HttpLimiter(_max float64) gin.HandlerFunc {
	var max float64
	var tbOptions limiter.ExpirableOptions
	tbOptions.DefaultExpirationTTL = time.Second // 默认按每秒

	if _max > 0 || _max <= 1000 {
		max = _max
	} else {
		max = 20
	}
	lmt := tollbooth.NewLimiter(max, &tbOptions) // 默认4次/秒，建议范围[1，40]

	return func(ctx *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, ctx.Writer, ctx.Request)
		if httpError != nil {
			Uri := ctx.Request.URL.Path
			// 记录日志
			logs.Warn("http_limit_log", zap.Any("Uri", Uri), zap.String("ClientIP", ctx.ClientIP()))
			ctx.JSON(429, gin.H{
				"status":  429,
				"message": "访问频率过高，请稍后再试",
				"data":    gin.H{
					//"ip": ctx.ClientIP(),
				},
			})

			ctx.Abort()
		} else {
			ctx.Next()
		}
	}
}
