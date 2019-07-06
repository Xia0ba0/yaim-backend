package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/websocket"
)
import (
	"yaim/controller"
	"yaim/middleware"
	"yaim/model/ormmodel"
	"yaim/service/userservice"
)

import (
	"net/http"
	"time"
)

var sessManager = sessions.New(sessions.Config{
	Cookie:  "YaimSession",
	Expires: 24 * time.Hour,
})
var engine, dberr = xorm.NewEngine("mysql", "root:1005@/test?charset=utf8")

func init() {
	if dberr != nil {
		fmt.Println(dberr.Error())
		return
	}

	dberr = engine.Sync2(new(ormmodel.Friend), new(ormmodel.User), new(ormmodel.Offline))
	if dberr != nil {
		defer engine.Close()

		fmt.Println(dberr.Error())
		return
	}
}

func main() {
	defer engine.Close()

	app := iris.New()
	app.Use(logger.New())

	app.Use(middleware.Allowall)

	mvc.Configure(app.Party("/websocket"), ConfigurePushMVC)
	mvc.Configure(app.Party("/user"), ConfigureUserMVC)

	_ = app.Run(iris.Addr(":8090"))
}

//此函数进行 动态/静态依赖注入与 中间件嵌入
func ConfigurePushMVC(app *mvc.Application) {

	//跨域访问websocket 仅供测试使用 后面取消配置
	ws := websocket.New(websocket.Config{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	})

	noAuthPath := make(map[string]string)

	//DEV
	noAuthPath["/websocket"] = "GET"
	app.Router.Use(middleware.NewAuther(sessManager, noAuthPath))

	//动态注入websocket连接 到 controller
	//动态注入 session 到 controller
	app.Register(
		ws.Upgrade,
		sessManager.Start)

	app.Handle(new(controller.PushController))
}
func ConfigureUserMVC(app *mvc.Application) {
	noAuthPath := make(map[string]string)
	noAuthPath["/user/login"] = "POST"
	noAuthPath["/user/register"] = "POST"
	app.Router.Use(middleware.NewAuther(sessManager, noAuthPath))

	// 动态注入session
	app.Register(sessManager.Start)

	// 静态注入service
	service := userservice.NewProvider(engine)
	app.Handle(&controller.Usercontroller{Service:service})
}
