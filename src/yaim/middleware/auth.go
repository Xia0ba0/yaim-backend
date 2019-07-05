package middleware

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
)

const USERIDKEY = "userid"

func NewAuther(sessionManager *sessions.Sessions, noAuthPath map[string]string) context.Handler{

	return func(ctx context.Context){
		if method, path :=noAuthPath[ctx.Path()]; path&&method==ctx.Method() {
			ctx.Next()
		}else{
			sess := sessionManager.Start(ctx)
			userID := sess.GetStringDefault(USERIDKEY, "")

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
	ctx.Header("Access-Control-Allow-Origin", "http://localhost:9080")
	ctx.Header("Access-Control-Allow-Credentials","true")
	ctx.Header("liliang","hhh")

	ctx.Next()
}