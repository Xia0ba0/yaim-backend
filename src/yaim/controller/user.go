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
			"message": "JSON parse failed",
			"Error":   err.Error(),
		})
		return
	}
	_, _ = c.Ctx.JSON(iris.Map{
		"message": "success",
		"data":    registerUser.UserID,
	})
}

// 通过Session 获取用户id
func (c *Usercontroller) getuserid() string {
	userID := c.Sess.GetStringDefault(USERIDKEY, "")
	return userID
}

func (c *Usercontroller) ifloggedin() bool {
	if c.getuserid() != "" {
		return true
	} else {
		return false
	}
}

//Method POST
// Path /user/login
func (c *Usercontroller) PostLogin() {
	var loginUser jsonmodel.TestForm

	if err := c.Ctx.ReadJSON(&loginUser); err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest) //400

		_, _ = c.Ctx.JSON(iris.Map{
			"message": "failed",
			"Error":   err.Error(),
		})
	} else if loginUser.UserID != "baoyuli" {
		_, _ = c.Ctx.JSON(iris.Map{
			"message": "Wrong Username",
			"data":    "",
		})
	} else {
		c.Sess.Set(USERIDKEY, loginUser.UserID)

		_, _ = c.Ctx.JSON(iris.Map{
			"message": "Success",
			"data":    loginUser.UserID,
		})
	}
}

func (c Usercontroller) GetLogin() {
	if c.ifloggedin() {
		_, _ = c.Ctx.JSON(iris.Map{
			"message": "Already logged in",
			"data":    c.getuserid(),
		})
	} else {
		_, _ = c.Ctx.JSON(iris.Map{
			"message": "Haven't logged in",
			"data":    "",
		})
	}
}

func (c Usercontroller) GetLogout(){
	if c.ifloggedin(){
		user := c.getuserid()
		c.Sess.Destroy()
		_, _ = c.Ctx.JSON(iris.Map{
			"message": "Already logged out",
			"data":    user,
		})
	}else{
		_, _ = c.Ctx.JSON(iris.Map{
			"message": "Haven't logged in",
			"data":    "",
		})
	}
}
