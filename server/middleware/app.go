package middleware

// 处理App运行时的一些必要事件

import (
	"fmt"
	"gofun/pkg/log"
	"gofun/pkg/tools"
	"time"

	"github.com/gin-gonic/gin"
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
		template := `{"name":"slow","uri":"%s","runtime":%f, "ip":"%s", "param":"%s", "body":"%s", "result":"%s"}`
		template = fmt.Sprintf(template, _host, statLatency, ctx.ClientIP(), _param, _body, _result)
		log.Log(template)
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
	buf := make([]byte, 1024)
	n, _ := ctx.Request.Body.Read(buf)
	ctx.Set("body", string(buf[0:n]))

	// 返回值
	ctx.Set("result", "")
	ctx.Next()
}
