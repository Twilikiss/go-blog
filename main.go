package main

import (
	"go-blog/common"
	"go-blog/router"
	"log"
	"net/http"
)

func init() {
	// 初始化模版加载
	common.LoadTemplates()
}

func main() {
	// 对于项目来说，一个项目只能有一个main入口
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	// 路由
	router.Router()
	log.Fatal(server.ListenAndServe())
}
