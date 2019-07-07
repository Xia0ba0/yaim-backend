package controller

import (
	"crypto/md5"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"net/url"
	"yaim/config"
	"yaim/service/mailservice"
)

import "yaim/model/jsonmodel"
import "yaim/service/userservice"

//PATH /user
type Usercontroller struct {
	//控制器依赖注入
	//动态绑定方式 每一个请求都有差异
	//注入的字段名必须大写
	Ctx  iris.Context
	Sess *sessions.Session

	UserService *userservice.UserServiceProvider
	MailService *mailservice.MailServiceProvider

	//静态绑定方式 所有的控制器共用一个实例
}

//动态路由配置
func (c *Usercontroller) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle("GET", "/verification", "Verification")
}

// 函数名 第一个字段为方法名 第二个字段为控制器路径
// Method POST
// Path /user/register
func (c *Usercontroller) PostRegister() {
	var registerUser jsonmodel.RegisterForm

	if err := c.Ctx.ReadJSON(&registerUser); err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest) //400

		_, _ = c.Ctx.JSON(iris.Map{
			"message": "Error",
			"Error":   err.Error(),
		})
		return
	}

	if err := c.UserService.Adduser(&registerUser); err != nil {

		_, _ = c.Ctx.JSON(iris.Map{
			"message": "Error",
			"Error":   err.Error(),
		})

		return
	}

	c.MailService.SendToken(registerUser.Email)

	_, _ = c.Ctx.JSON(iris.Map{
		"message": "Success",
		"data":    "Please Check your Verification Email",
	})
}

func (c *Usercontroller) Verification() string {
	user := c.Ctx.URLParamDefault("user","no")
	token := c.Ctx.URLParamDefault("token","no")

	if user=="no" || token=="no"{
		return "ParamError"
	}

	user, _ = url.QueryUnescape(user)
	if !c.UserService.Checkuser(user){
		return "ParamError"
	}

	if c.UserService.CheckVerification(user){
		return "Already Verified"
	}

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(user))
	cipherStr := fmt.Sprintf("%x",md5Ctx.Sum([]byte(config.TokenKey)))

	if token != cipherStr{
		return "Verify Failed"
	}

	_ = c.UserService.Verificate(user)
	return "Success"
}


// 通过Session 获取用户id
func (c *Usercontroller) getuserid() string {
	userID := c.Sess.GetStringDefault(config.UserIdKey, "")
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
			"Error":   err.Error(),
		})
		return
	}

	if err := c.UserService.CheckIdentity(loginUser.Email, loginUser.Password); err != nil {
		_, _ = c.Ctx.JSON(iris.Map{
			"message": "Error",
			"Error":   err.Error(),
		})
		return
	}

	c.Sess.Set(config.UserIdKey, loginUser.Email)
	_, _ = c.Ctx.JSON(iris.Map{
		"message": "Success",
		"data": iris.Map{
			"userid": loginUser.Email,
			"Cookie": config.CookieName + "=" + c.Sess.ID(),
		},
	})
}

func (c Usercontroller) GetLogout() {
	user := c.getuserid()
	c.Sess.Destroy()
	_, _ = c.Ctx.JSON(iris.Map{
		"message": "Success",
		"data":    user,
	})
}
