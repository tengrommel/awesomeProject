package main

import (
	"github.com/kataras/iris"
)

func main() {
	// Initialize app
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h2>Hello world this is Iris</h2>")
	})
	app.Get("/text", func(ctx iris.Context) {
		ctx.Text("Hello world this is text format")
	})
	app.Get("/write", func(ctx iris.Context) {
		ctx.WriteString("Hello world this is string")
	})
	app.Get("/api", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"Msg": "Hello this is JSON"})
	})
	app.Handle("GET", "/simple", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"msg": "ok"})
	})
	app.Run(iris.Addr(":5000"))
}
