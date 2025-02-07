package main

import (
	"goji.io/v3"
	"log"
	"net/http"
	"note-goji/src/application"
	"note-goji/src/router"
)

func main() {
	// 新建 Goji 实例
	app := goji.NewMux()
	// 注册路由
	router.RegisterRouter(app)
	// 启动服务
	addr := application.App.Server.Host + ":" + application.App.Server.Port
	log.Println("Server listen on " + addr)
	http.ListenAndServe(addr, app)
}
