package main

import (
	"vue_demo_backend/router"

	"github.com/kataras/iris/v12"
)

func main() {
	app := newApp()

	router.RegisterRouter(app)

	app.Listen(":8866")
}

func newApp() *iris.Application {
	app := iris.Default()
	return app
}
