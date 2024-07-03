package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 设置路由处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "欢迎访问我的第一个Go Web后端！")
	})

	// 设置监听端口
	port := "8080"
	fmt.Printf("服务器启动，监听端口：%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("启动服务器失败： %v\n", err)
	}
}
