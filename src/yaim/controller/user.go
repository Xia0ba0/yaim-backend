package controller

import (
	"github.com/kataras/iris"
)
import "yaim/model/jsonmodel"

//PATH /user
type Usercontroller struct {
	//控制器依赖注入
	//动态绑定方式 每一个请求都有差异
	//注入的字段名必须大写
	Ctx iris.Context

	//静态绑定方式 所有的控制器共用一个实例
}

// 函数名 第一个字段为方法名 第二个字段为控制器路径
// Method POST
// Path /user/register
func (c *Usercontroller) PostRegister(){
	var registerUser jsonmodel.TestForm

	//解析JSON
	//curl -X POST --data {\"userid\":\"baoyuli\"} -H "Content-Type:application/json" http://localhost:8090/user/register
	if err := c.Ctx.ReadJSON(&registerUser); err!=nil{
		c.Ctx.StatusCode(iris.StatusBadRequest) //400

		_, _ = c.Ctx.JSON(iris.Map{
			"message":"failed",
			"Error":err.Error(),
		})
		return
	}
	_, _ = c.Ctx.JSON(iris.Map{
		"message":"success",
		"data":registerUser.UserID,
	})
}

// Method GET
// Path /user/login
func (c *Usercontroller)GetLogin() string{
	return "ok"
}

//Method POST
// Path /user/login
func (c *Usercontroller) PostLogin() string{

	return "ok"
}
