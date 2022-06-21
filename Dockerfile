# 基础镜像
FROM docker.io/alpine:3.13.5
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
# 时区文件
COPY localtime /etc/localtime
# 用RUN_MODE来区分环境 (0-开发环境,1-生产,2-测试,3-预上线)
ENV RUN_MODE 1
# 代码目录
WORKDIR /data/work/gofun
# 拷贝程序和配置文件
COPY main ./gofun
COPY conf ./conf
COPY cmd ./cmd
# 拷贝其他文件，如静态文件、模版等目录，在下面逐条新增COPY指令
COPY public ./public
# COPY static ./static
# 程序监听的端口，根据实际情况修改
EXPOSE 9000
CMD ["./main"]
