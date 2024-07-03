FROM golang:1.17-alpine AS build

WORKDIR /app

COPY . .

RUN go build -o main .

FROM alpine:latest

COPY --from=build /app/main /app/main

EXPOSE 8080

# 设置容器启动命令
CMD ["/app/main"]
