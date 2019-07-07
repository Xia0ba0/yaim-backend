package middleware

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
)

import(
	"yaim/config"
)

func NewAuther(sessionManager *sessions.Sessions, noAuthPath map[string]string) context.Handler{

	return func(ctx context.Context){
		if method, path :=noAuthPath[ctx.Path()]; path&&method==ctx.Method() {
			ctx.Next()
		}else{
			sess := sessionManager.Start(ctx)
			userID := sess.GetStringDefault(config.UserIdKey, "")

			if userID == ""{
				ctx.StatusCode(iris.StatusForbidden)
				_, _ = ctx.JSON(iris.Map{
					"message":"Error",
					"data":"No Authentication",
				})
			}else{
				ctx.Next()
			}
		}
	}
}

func Allowall(ctx context.Context){
	ctx.Header("Access-Control-Allow-Origin", config.HostName)
	ctx.Header("Access-Control-Allow-Credentials","true")

	ctx.Next()
}