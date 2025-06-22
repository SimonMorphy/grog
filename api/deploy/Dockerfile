# 第一阶段：构建应用
FROM golang:1.24.2-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件
COPY ../go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY .. .

# 构建应用，开启CGO会导致生成的二进制文件依赖于特定的C库
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 第二阶段：创建轻量级运行时镜像
FROM alpine:latest

# 安装必要的证书
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /root/

# 从builder阶段复制构建好的应用
COPY --from=builder /app/main .

# 暴露应用端口
EXPOSE 8080

# 运行应用
CMD ["./main"]