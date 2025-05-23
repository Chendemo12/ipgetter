# 阶段一：构建镜像
FROM golang:1.23 AS builder

WORKDIR /app
# 将源代码复制到容器中
COPY . .

# 编译可执行程序
ENV GOPROXY=https://goproxy.cn,direct
ENV CGO_ENABLED=0
RUN go mod download
RUN go build -tags=jsoniter -ldflags="-s -w" -gcflags='-l -l -l -m' -o ipgetter


# 阶段二：构建最终镜像
FROM alpine:latest

# 指定镜像名称和版本号
LABEL maintainer="ipgetter"
LABEL version="1.0.0"

# 复制编译后的可执行程序到最终镜像中
COPY --from=builder /app/ipgetter .

# 设置环境变量
ENV DEBUG=1
ENV LISTEN_PORT=7290

# 暴露端口
EXPOSE 7290

# 运行可执行程序
CMD ["./ipgetter"]
