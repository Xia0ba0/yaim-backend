package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"

	"github.com/kataras/iris/middleware/logger"

	"yaim/controller"
)


//TODO: session管理器 依赖注入与中间件
func main(){
	app := iris.New()
	app.Use(logger.New())


	mvc.Configure(app.Party("/user"),userMVC)

	_ = app.Run(iris.Addr(":8090"))
}

//此函数进行 动态/静态依赖注入与 中间件嵌入
func userMVC(app *mvc.Application){
	app.Handle(new(controller.Usercontroller))
}