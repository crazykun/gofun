package middleware

import (
	"bytes"
	"gofun/pkg/logs"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
				logs.Warn(err.Error())
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
		// body, is_body := ctx.Get("body")
		// if !is_body {
		// 	body = ""
		// }
		var field []zap.Field
		field = append(field, zap.Int("statusCode", ctx.Writer.Status()))
		field = append(field, zap.Float64("runtime", float64(end.Sub(start).Nanoseconds()/1e4)/100.0))
		field = append(field, zap.String("clientIp", ctx.ClientIP()))
		field = append(field, zap.String("method", ctx.Request.Method))
		field = append(field, zap.String("uri", ctx.Request.URL.Path))
		field = append(field, zap.String("path", ctx.Request.RequestURI))
		field = append(field, zap.Any("param", param))
		// field = append(field, zap.Any("body", body))
		field = append(field, zap.Any("result", result))
		logs.Info("access_log", field...)

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
