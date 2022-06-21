package middleware

// 接管服务器500错误，使错误可视化

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gofun/conf"
	"gofun/pkg/log"
	"gofun/pkg/tools"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
)

type template struct {
	ErrorName   string `json:"ErrorName"`
	ErrorMsg    string `json:"ErrorMsg"`
	RequestTime string `json:"RequestTime"`
	RequestURL  string `json:"RequestURL"`
	RequestUA   string `json:"RequestUA"`
	RequestBody string `json:"RequestBody"`
	RequestIP   string `json:"RequestIP"`
	DebugStack  string `json:"DebugStack"`
}

// AppError500 抛出500错误
func AppError500(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			DebugStack := ""
			for _, v := range strings.Split(string(debug.Stack()), "\n") {
				DebugStack += v + `\n`
			}

			// 获取请求信息
			buf := make([]byte, 1024)
			n, _ := ctx.Request.Body.Read(buf)

			errorName := fmt.Sprintf("【服务异常】%s 运行异常", conf.Config.AppName)
			var t = template{
				RequestTime: tools.GetTimeDate("Y-m-d H:i:s"),
				ErrorName:   errorName,
				ErrorMsg:    fmt.Sprintf("%s", err),
				RequestURL:  ctx.Request.Method + "  " + ctx.Request.Host + ctx.Request.RequestURI,
				RequestUA:   ctx.Request.UserAgent(),
				RequestBody: string(buf[0:n]),
				RequestIP:   ctx.ClientIP(),
				DebugStack:  DebugStack,
			}

			_txt, err := json.Marshal(t)
			if err == nil {
				log.Err(string(_txt))
			}

			bot := conf.Config.Warnbot.Wx
			if bot != "" {
				warnContent := `#### %s提醒 \n> ErrorMsg: %s \n> RequestTime: %s \n> RequestURL: %s \n> RequestUA: %s \n> RequestBody: %s \n> RequestIP: %s \n\n %s\n`
				warnContent = fmt.Sprintf(warnContent, t.ErrorName, t.ErrorMsg, t.RequestTime, t.RequestURL, t.RequestUA, t.RequestBody, t.RequestIP, DebugStack)

				body := fmt.Sprintf(`{"msgtype": "markdown", "markdown": {"content": "%s"}}`, warnContent)
				http.Post(
					"https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+bot,
					"application/json",
					bytes.NewBuffer([]byte(body)))
			}

			// 返回
			ctx.JSON(500, gin.H{
				"status":  500,
				"message": "系统异常，请稍后重试！",
				"data":    gin.H{},
			})
		}
	}()
	//加载完 defer recover，继续后续接口调用并返回JSON提示
	ctx.Next()
}
