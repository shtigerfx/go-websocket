FROM golang:1.21 AS build-dist
ENV GOPROXY='https://proxy.golang.org,direct'
WORKDIR /data/release
COPY . .
RUN go build

# 更换系统镜像
FROM alpine:latest as prod

# 安装 GLIBC 和其他运行时库
RUN apk --no-cache add ca-certificates libc6-compat

WORKDIR /data/go-websocket
COPY --from=build-dist /data/release/go-websocket ./
COPY --from=build-dist /data/release/conf /data/go-websocket/conf

EXPOSE 7800

CMD ["/data/go-websocket/go-websocket","-c","./conf/app.ini"]