# golang框架模板

## 框架说明

| 功能 | 组件	| 备注
| ------------------- | ------------------- | ------------------- |
| 基础框架 |gin |github.com/gin-gonic/gin |
| 环境管理 |自研 | 可动态区分测试,正式,预上线等 |
| 配置读取 |viper | 动态监听, 可以远程读取etcd/consul等 |
| 路由规则 | | |
| 日志组件 | zap|github.com/uber-go/zap |
| 错误处理 | | |
| 类型转换 | | |
| 缓存管理 | | |
| 链路跟踪 | | |
| 热重载 |air,fresh,bee,gowatch,gin自带,realize | |
| 优雅重启 |zerodown |github.com/douglarek/zerodown |
| 连接池 | | |
| 依赖注入 |wire |github.com/google/wire/ |
| Redis |go-redis |https://github.com/go-redis/redis |
| orm |Gorm | https://github.com/go-gorm/gorm/|
| rpc | | |
| 文档生成 | | |

## 推荐组件

## 热重载
实现代码变更自动重新加载执行

| 名称 | 地址	| 备注
| ------------------- | ------------------- | ------------------- |
| air |https://github.com/cosmtrek/air | 彩色日志, 二进制构建 |
| fresh |https://github.com/gravityblast/fresh | 自动启动web |
| bee |https://github.com/beego/bee | beego框架的热编译工具 |
| gowatch |https://github.com/silenceper/gowatch | 监听当前目录下的相关文件变动，进行实时编译 |
| gin |https://github.com/codegangsta/gin | 网络应用代理模式 |
| realize |https://github.com/oxequa/realize | 高性能实时刷新, 多个项目同时代理 |


## 目录结构介绍

1. 总体目录

```

project  
└───app                   应用目录
│   └───controller        控制器目录
│   └───logic             逻辑层
│   └───model             数据层
│       │   ...
│   └───route             路由
└───assets                静态文件
└───config                配置目录
│   └───dev                  开发环境配置
│   └───test                 测试环境配置
│   └───pre                  预发布环境配置
│   └───prod                 生产环境配置
└───log                   日志目录
└───vendor                第三方包
│   README.md   
│   ...

```


2. 常用命令

```golang
   1. 下载依赖
      go mod tidy
      go mod vendor
   2. 服务启动  
      go run main.go
      go run main.go -mode prod
   3. 编译
      go build ./ 
      go build -o main.exe main.go


```