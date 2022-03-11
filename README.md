# golang框架模板

## 框架说明

| 功能 | 组件	| 备注
| ------------------- | ------------------- | ------------------- |
| 基础框架 |gin |github.com/gin-gonic/gin |
| 配置管理 |自研 | 可动态区分测试,正式,预上线等 |
| 路由规则 | | |
| 日志组件 | | |
| 错误处理 | | |
| 类型转换 | | |
| 缓存管理 | | |
| 链路跟踪 | | |
| 热重载 | | |
| 优雅重启 |zerodown |github.com/douglarek/zerodown |
| 连接池 | | |
| orm | | |
| rpc | | |
| 文档生成 | | |


## 目录结构介绍

1. 总体目录

```

project  
└───App                   应用目录
│   └───Common            公共文件
│       |  function.go     公共方法
│   └───Controller        控制器目录
│   └───Logic             逻辑层
│   └───Model             数据层
│   └───Library           封装的应用
│       │   ...
└───Config                配置目录
│   └───dev    
│   └───test             
│   └───pro             
└───Log                   日志目录
└───Temp                  存放临时文件，进程通信也放在这里
└───vendor                第三方包
│   README.md   
│   ...

```


2. 常用命令

```golang
   1. 服务启动  
      go run main.go
      go build ./ 

```