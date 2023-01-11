# golang框架模板

## 框架使用
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
   4. 热更新测试
      fresh


```


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


