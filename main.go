package main

import (
	"gofun/internal/cache"
	"gofun/internal/db"
	"gofun/server"

	"github.com/gin-gonic/gin"
)

// 运行环境
var HttpServer *gin.Engine

// 初始化
func main() {

	// 初始化服务
	server.StartServer(HttpServer)

	// 关闭服务
	defer db.CloseMysqlPool()
	defer cache.CloseRedisPool()
}
