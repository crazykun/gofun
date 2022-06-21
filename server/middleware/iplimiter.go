package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// 基于IP地址的请求限制器。它可以与redis和滑动窗口机制一起使用。
func NewRateLimiterMiddleware(redisClient *redis.Client, key string, limit int, slidingWindow time.Duration) gin.HandlerFunc {

	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprint("error init redis", err.Error()))
	}

	return func(ctx *gin.Context) {
		now := time.Now().UnixNano()
		userCntKey := fmt.Sprint(ctx.ClientIP(), ":", key)

		redisClient.ZRemRangeByScore(userCntKey,
			"0",
			fmt.Sprint(now-(slidingWindow.Nanoseconds()))).Result()

		reqs, _ := redisClient.ZRange(userCntKey, 0, -1).Result()

		if len(reqs) >= limit {
			ctx.JSON(http.StatusTooManyRequests, gin.H{
				"status":  http.StatusTooManyRequests,
				"message": "too many request",
				"data":    gin.H{},
			})

			ctx.Abort()
		}

		ctx.Next()
		redisClient.ZAddNX(userCntKey, redis.Z{Score: float64(now), Member: float64(now)})
		redisClient.Expire(userCntKey, slidingWindow)
	}

}
