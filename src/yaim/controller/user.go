package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)
import "yaim/model/jsonmodel"

const USERIDKEY = "userid"

//PATH /user
type Usercontroller struct {
	//控制器依赖注入
	//动态绑定方式 每一个请求都有差异
	//注入的字段名必须大写
	Ctx  iris.Context
	Sess *sessions.Session

	//静态绑定方式 所有的控制器共用一个实例
}

// 函数名 第一个字段为方法名 第二个字段为控制器路径
// Method POST
// Path /user/register
func (c *Usercontroller) PostRegister() {
	var registerUser jsonmodel.TestForm

	//解析JSON
	//curl -X POST --data {\"userid\":\"baoyuli\"} -H "Content-Type:application/json" http://localhost:8090/user/register
	if err := c.Ctx.ReadJSON(&registerUser); err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest) //400

		_, _ = c.Ctx.JSON(iris.Map{
			"message": "Error",
			"Error":   err.Error(),
		})
		return
	}
	_, _ = c.Ctx.JSON(iris.Map{
		"message": "Success",
		"data":    registerUser.Email,
	})
}

// 通过Session 获取用户id
func (c *Usercontroller) getuserid() string {
	userID := c.Sess.GetStringDefault(USERIDKEY, "")
	return userID
}

//Method POST
// Path /user/login
func (c *Usercontroller) PostLogin() {
	var loginUser jsonmodel.TestForm

	if err := c.Ctx.ReadJSON(&loginUser); err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest) //400

		_, _ = c.Ctx.JSON(iris.Map{
			"message": "Error",
			"data":   err.Error(),
		})
	} else if loginUser.Email != "123" || loginUser.Password != "123" {
		_, _ = c.Ctx.JSON(iris.Map{
			"message": "Error",
			"data":    "Wrong UserName",
		})
	} else {
		c.Sess.Set(USERIDKEY, loginUser.Email)
		_, _ = c.Ctx.JSON(iris.Map{
			"message": "Success",
			"data":    iris.Map{
				"userid":loginUser.Email,
				"Cookie":"YaimSession="+ c.Sess.ID(),
			},
		})
	}
}

func (c Usercontroller) GetLogout() {
	user := c.getuserid()
	c.Sess.Destroy()
	_, _ = c.Ctx.JSON(iris.Map{
		"message": "Success",
		"data":    user,
	})
}
