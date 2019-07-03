package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/websocket"
	"net/http"
	"yaim/controller"
)


//TODO: session管理器 依赖注入与中间件
func main(){
	app := iris.New()
	app.Use(logger.New())

	mvc.Configure(app.Party("/websocket"),ConfigurePushMVC)
	mvc.Configure(app.Party("/user"),ConfigureUserMVC)

	_ = app.Run(iris.Addr(":8090"))
}


//此函数进行 动态/静态依赖注入与 中间件嵌入
func ConfigurePushMVC(app *mvc.Application){

	//跨域访问websocket 仅供测试使用 后面取消配置
	ws := websocket.New(websocket.Config{
		CheckOrigin:func(r *http.Request) bool{
			return true
		},
	})

	//动态注入websocket连接 到 controller
	app.Register(ws.Upgrade)
	app.Handle(new(controller.PushController))
}
func ConfigureUserMVC(app *mvc.Application){
	app.Handle(new(controller.Usercontroller))
}