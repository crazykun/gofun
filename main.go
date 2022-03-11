package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// 运行环境
var runMode string

func main() {

	flag.StringVar(&runMode, "m", "dev", "0-dev, 1-生产，2-测试，3-预上线")
	flag.Parse()

	if runMode != "" {
		runMode = os.Getenv("RUN_MODE")
		if runMode == "1" {
			runMode = "product"
		} else if runMode == "2" {
			runMode = "test"
		} else if runMode == "3" {
			runMode = "pre"
		} else {
			runMode = "dev"
		}
	}

	// 获取命令行参数
	fmt.Println("run mode", runMode)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()

}
