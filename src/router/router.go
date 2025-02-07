package router

import (
	"goji.io/v3"
	"goji.io/v3/pat"
	"note-goji/src/controller"
)

// RegisterRouter 注册路由
func RegisterRouter(j *goji.Mux) {
	// 添加命令相关子路由
	s := goji.SubMux()
	s.HandleFunc(pat.Get("/:commandId"), controller.One)
	s.HandleFunc(pat.Get("/"), controller.List)
	s.HandleFunc(pat.Post("/"), controller.Insert)
	s.HandleFunc(pat.Post("/batch"), controller.InsertBatch)
	s.HandleFunc(pat.Put("/"), controller.Update)
	s.HandleFunc(pat.Put("/batch"), controller.UpdateBatch)
	s.HandleFunc(pat.Delete("/:commandId"), controller.Delete)
	s.HandleFunc(pat.Delete("/batch"), controller.DeleteBatch)
	s.HandleFunc(pat.Get("/select/:commandName"), controller.Select)
	s.HandleFunc(pat.Get("/name-list"), controller.NameList)
	// 新建命令路由组
	j.Handle(pat.New("/command/*"), s)
}
