package middleware

import (
	"bytes"
	"gofun/pkg/log"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// 开始时间
		start := time.Now()
		// 请求报文
		var requestBody []byte
		if ctx.Request.Body != nil {
			var err error
			requestBody, err = ctx.GetRawData()
			if err != nil {
				log.Warn(map[string]interface{}{"err": err.Error()}, "get http request body error")
			}
			ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		}
		// 处理请求
		ctx.Next()
		// 结束时间
		end := time.Now()
		param, is_param := ctx.Get("param")
		if !is_param {
			param = ""
		}
		result, is_result := ctx.Get("result")
		if !is_result {
			result = ""
		}
		log.Info(map[string]interface{}{
			"statusCode": ctx.Writer.Status(),
			"runtime":    float64(end.Sub(start).Nanoseconds()/1e4) / 100.0,
			"clientIp":   ctx.ClientIP(),
			"method":     ctx.Request.Method,
			"uri":        ctx.Request.RequestURI,
			"param":      param,
			"result":     result,
		})

	}
}

// 日志记录到 MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// 日志记录到 ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// 日志记录到 MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
