package router

import (
	"vue_demo_backend/handler"

	"github.com/kataras/iris/v12"
)

// RegisterRouter 注册路由
func RegisterRouter(app *iris.Application) {
	api := app.Party("/api")
	{
		api.Get("/insert", handler.InsertTmpPost)
		api.Get("/posts", handler.GetPosts)
		api.Get("/post/{id:uint}", handler.GetPost)
		api.Post("/post", handler.InsertPost)
		api.Get("/delete/{id:uint}", handler.DeletePost)
	}
}
