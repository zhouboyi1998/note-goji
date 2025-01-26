package main

import (
	"encoding/json"
	"goji.io/v3"
	"goji.io/v3/pat"
	"log"
	"net/http"
)

func main() {
	// 新建 Goji 实例
	app := goji.NewMux()

	// Hello World
	app.HandleFunc(pat.Get("/hello/:name"), func(w http.ResponseWriter, r *http.Request) {
		name := pat.Param(r, "name")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Hello, " + name,
		})
	})

	// 启动服务
	log.Println("Server listen on :18084")
	http.ListenAndServe(":18084", app)
}
