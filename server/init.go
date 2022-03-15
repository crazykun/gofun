package server

import (
	"flag"
	"gofun/app/util"
	"log"
	"os"
	"runtime"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var AppPath string
var RunMode string
var Config string

// go环境配置检测
func init() {
	AppPath, _ = os.Getwd()

	// go版本检测
	chkVersion()

	// 初始化数据库
	initDb()

	// 初始化缓存
	initCache()

	// 初始化日志
	initLog()
}

// go版本检测
func chkVersion() {
	// 基础环境自检
	log.Println("系统Go环境自检...")

	var goVersion string = runtime.Version()
	var goLimitVersion string = "1.15.0"
	isLimit := util.VersionCompare(goVersion, goLimitVersion, ">=")
	if isLimit {
		log.Println("Go版本符合要求 >>> 当前Go版本：", goVersion, " 最小要求版本："+goLimitVersion)
	} else {
		log.Println("Go版本太低，服务自动中断！", "当前Go版本："+goVersion, " 最小要求版本："+goLimitVersion, "请升级Go版本！")
		os.Exit(200)
	}
}

// 初始化数据库
func initDb() {

}

// 初始化缓存
func initCache() {

}

// 初始化日志
func initLog() {

}

// 配置文件自检
func chkConfig() {
	flag.StringVar(&RunMode, "mode", "dev", "dev-开发环境, prod-生产，test-测试，pre-预上线")
	flag.Parse()

	if RunMode == "" {
		// 从环境变量中获取
		RunMode = os.Getenv("RUN_MODE")
		if RunMode == "1" {
			RunMode = "product"
		} else if RunMode == "2" {
			RunMode = "test"
		} else if RunMode == "3" {
			RunMode = "pre"
		} else {
			RunMode = "dev"
		}
	} else if !util.In_array(RunMode, []string{"dev", "test", "pre", "prod"}) {
		log.Println("err mode", RunMode)
		os.Exit(200)
	}

	// 获取命令行参数
	log.Println("run mode", RunMode)

	// 配置文件自检
	initConfig()

}

// 初始化配置
func initConfig() {
	viper.SetConfigName(RunMode)            //配置文件名
	viper.SetConfigType("ini")              //配置文件类型
	viper.AddConfigPath(AppPath + "/conf/") //执行go run对应的路径配置
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.Unmarshal(&Config)
	if util.In_array(RunMode, []string{"prod", "pre"}) {
		// 关闭gin的debug模式
		gin.SetMode(gin.ReleaseMode)
	}
	// 监听配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file:%s Op:%s\n", e.Name, e.Op)
	})
}

// 启动服务
func StartServer(HttpServer *gin.Engine) {
	chkConfig()

	// Gin服务
	HttpServer = gin.Default()

	// 测试访问IP
	host := "127.0.0.1:8080"

	err := HttpServer.Run(host)
	if err != nil {
		log.Println("http服务遇到错误，运行中断，error：", err.Error())
		log.Println("提示：注意端口被占时应该首先更改对外暴露的端口，而不是微服务的端口。")
		os.Exit(200)
	}

	return
}
