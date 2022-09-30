package middleware

// 处理App运行时的一些必要事件

import (
	"bytes"
	"gofun/pkg/logs"
	"gofun/pkg/tools"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 接口运行耗时
func Runtime(ctx *gin.Context) {
	start := float64(time.Now().UnixNano()) / 1000000  // ms
	ctx.Set("reptime", tools.GetTimeDate("i.s.ms.ns")) // 设置公共参数。此参数范围是每个请求的公共参数，不是超全局参数，超全局参数请用globalData。

	// 等其他中间件先执行
	ctx.Next()

	// 获取运行耗时，ms
	end := float64(time.Now().UnixNano()) / 1000000
	runtime := end - start

	// 设置公共参数（设置ctx每次请求的全局值）
	ctx.Set("runtime", runtime)
	//fmt.Println("本次运行耗时=", runtime, "ms")

	// 进入耗时治理服务
	afterRequest(ctx)

	// 计时完成，中断所有后续函数调用
	ctx.Abort()
}

// 请求结束后处理
func afterRequest(ctx *gin.Context) {

	_host, _ := ctx.Get("host")
	_runtime, _ := ctx.Get("runtime")
	_param, _ := ctx.Get("param")
	_body, _ := ctx.Get("body")
	_result, _ := ctx.Get("result")

	// 记录慢日志
	statLatency := tools.StringToFloat(tools.ValueInterfaceToString(_runtime))
	// 超过3s都记录下来
	if statLatency > 3*1000 {
		logs.Warn("slow_log_warn",
			zap.Any("host", _host),
			zap.Float64("runtime", statLatency),
			zap.String("ip", ctx.ClientIP()),
			zap.Any("param", _param),
			zap.Any("body", _body),
			zap.Any("result", _result),
		)
	}

	ctx.Next()
}

// Param 设置全局参数
func CommonParam(ctx *gin.Context) {
	// 设置公共参数
	ctx.Set("app_name", "gofun")
	ctx.Set("host", ctx.Request.Host+ctx.Request.RequestURI)
	// url参数
	ctx.Set("param", ctx.Request.URL.RawQuery)

	// 获取body参数
	buf, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.Set("body", string(buf))
	}
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(buf))

	// 返回值
	ctx.Set("result", "")
	ctx.Next()
}
