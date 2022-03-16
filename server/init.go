package server

import (
	"flag"
	routes "gofun/app/route"
	"gofun/conf"
	"gofun/server/middleware"
	"gofun/tools"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var AppPath string
var RunMode string

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
	isLimit := tools.VersionCompare(goVersion, goLimitVersion, ">=")
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
	} else if !tools.InArray(RunMode, []string{"dev", "test", "pre", "prod"}) {
		log.Println("err mode", RunMode)
		os.Exit(200)
	}

	// 获取命令行参数
	log.Println("run mode", RunMode)
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
	viper.SetDefault("AppPath", AppPath)
	if tools.InArray(RunMode, []string{"prod", "pre"}) {
		viper.SetDefault("RunMode", gin.ReleaseMode)
	} else {
		viper.SetDefault("RunMode", gin.DebugMode)
	}
	viper.Unmarshal(&conf.Config)
	// fmt.Printf("%+v", conf.Config)

	// 监听配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file:%s Op:%s\n", e.Name, e.Op)
	})
}

// 启动服务
func StartServer(HttpServer *gin.Engine) {
	// 配置文件自检
	chkConfig()
	// 初始化配置
	initConfig()

	// Gin服务
	HttpServer = gin.Default()

	// 捕捉接口运行耗时（必须排第一）
	HttpServer.Use(middleware.Runtime)

	// 设置全局ctx参数（必须排第二）
	HttpServer.Use(middleware.CommonParam)

	// 拦截应用500报错，使之可视化
	HttpServer.Use(middleware.AppError500)

	// Gin运行时：release、debug、test
	gin.SetMode(conf.Config.RunMode)

	// 静态路径
	HttpServer.Static("/assets", "./assets")

	// 注册路由
	routes.RegisterRoutes(HttpServer)

	// 模板目录
	HttpServer.LoadHTMLGlob(AppPath + "/app/view/*")

	// // 注册其他路由，可以自定义
	// routes.RouterRegister(HttpServer)
	// //Router.Api(HttpServer) // 面向Api
	// //Router.Web(HttpServer) // 面向模版输出

	// // 初始化定时器（立即运行定时器）
	// Task.TimeInterval(0, 0, "0")

	// // 实际访问网址和端口
	host := "127.0.0.1:" + strconv.Itoa(conf.Config.AppPort)

	// 终端提示
	log.Println(
		"\n App Success! \n\n " +
			" \n " +
			"访问地址示例：http://" + host + " >>> \n " +
			"1) 运行模式：" + conf.Config.RunMode + " \n " +
			"2) 运行目录：" + conf.Config.AppPath +
			"")

	err := HttpServer.Run(host)
	if err != nil {
		log.Println("http服务初始化失败，error：", err.Error())
		os.Exit(200)
	}

	return
}
