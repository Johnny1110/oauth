# 使用官方的 golang 映像作為構建階段的基礎映像
FROM golang:1.21.10 AS builder

# 設置工作目錄
WORKDIR /app

# 複製 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# resources 內有設定 yml 檔，需要帶入到 container
COPY resources/* ./resources/

# 下載依賴
RUN go mod download

# 複製源代碼
COPY . .

# 編譯 Go 應用程序
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 使用一個更小的映像作為運行階段的基礎映像
FROM alpine:latest

# 安裝必要的運行時依賴
RUN apk --no-cache add ca-certificates

# 設置工作目錄
WORKDIR /root/

# 從構建階段複製編譯好的二進制文件
COPY --from=builder /app/main .

# 複製 resources 目錄中的配置文件
COPY --from=builder /app/resources ./resources

# 確保二進制文件具有執行權限
RUN chmod +x ./main

# 暴露應用程序運行的端口
EXPOSE 8080

# 設置容器啟動時執行的命令
CMD ["./main"]
