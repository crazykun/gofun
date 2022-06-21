package middleware

import (
	"fmt"
	"gofun/pkg/log"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
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
			template := `{"name":"httpLimiter","uri":"%s", "ip":"%s", "result": "访问频率过高"}`
			Uri := ctx.Request.RequestURI
			template = fmt.Sprintf(template, Uri, ctx.ClientIP())
			log.Log(template)

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
