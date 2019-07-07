package controller

import (
	"crypto/md5"
	"fmt"
	"github.com/kataras/iris"
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

	//静态绑定方式 所有的控制器共用一个实例
	UserService *userservice.UserServiceProvider
	MailService *mailservice.MailServiceProvider
}

// 函数名 第一个字段为方法名 第二个字段为控制器路径
// Method POST
// Path /user/register
// function: 用户注册 并且发送验证邮件
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

	//c.MailService.SendToken(registerUser.Email)

	_, _ = c.Ctx.JSON(iris.Map{
		"message": "Success",
		"data":    "Please Check your Verification Email",
	})
}

// 函数名 第一个字段为方法名 第二个字段为控制器路径
// Method GET
// Path /user/verification
// function: 用户邮箱验证
func (c *Usercontroller) GetVerification() string {
	user := c.Ctx.URLParamDefault("user", "no")
	token := c.Ctx.URLParamDefault("token", "no")

	if user == "no" || token == "no" {
		return "ParamError"
	}

	user, _ = url.QueryUnescape(user)
	if !c.UserService.Checkuser(user) {
		return "ParamError"
	}

	if c.UserService.CheckVerification(user) {
		return "Already Verified"
	}

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(user))
	cipherStr := fmt.Sprintf("%x", md5Ctx.Sum([]byte(config.TokenKey)))

	if token != cipherStr {
		return "Verify Failed"
	}

	_ = c.UserService.Verificate(user)
	return "Success"
}

// Method POST
// Path /user/login
// function 用户登录
func (c *Usercontroller) PostLogin() {
	var loginUser jsonmodel.LoginForm

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

// Method POST
// Path /user/logout
// function 用户注销
func (c *Usercontroller) GetLogout() {
	user := c.getuserid()
	c.Sess.Destroy()
	_, _ = c.Ctx.JSON(iris.Map{
		"message": "Success",
		"data":    user,
	})
}

// Method POST
// Path /user/key
// function 用户上传公钥
func (c *Usercontroller) PostKey() {
	var keyForm jsonmodel.UploadKeyForm

	if err := c.Ctx.ReadJSON(&keyForm); err != nil {
		_, _ = c.Ctx.JSON(iris.Map{
			"message": "Error",
			"Error":   err.Error(),
		})
		return
	}

	userid := c.getuserid()
	if err := c.UserService.Updatepubkey(userid, keyForm.Key); err != nil {
		_, _ = c.Ctx.JSON(iris.Map{
			"message": "Error",
			"Error":   err.Error(),
		})
		return
	}

	_, _ = c.Ctx.JSON(iris.Map{
		"message": "Success",
		"data":    "Upload public key success",
	})
}

// Method POST
// Path /user/address
// function 用户上传地址
func (c *Usercontroller) PostAddress() {
	var addrForm jsonmodel.UpdateAddrForm

	if err := c.Ctx.ReadJSON(&addrForm); err != nil {
		_, _ = c.Ctx.JSON(iris.Map{
			"message": "Error",
			"Error":   err.Error(),
		})
		return
	}

	userid := c.getuserid()
	if err := c.UserService.UpdateNetAddr(userid, addrForm.Ip, addrForm.Port); err != nil {
		_, _ = c.Ctx.JSON(iris.Map{
			"message": "Error",
			"Error":   err.Error(),
		})
		return
	}

	_, _ = c.Ctx.JSON(iris.Map{
		"message": "Success",
		"data":    "Update network address success",
	})
}

// 通过Session 获取用户id
func (c *Usercontroller) getuserid() string {
	userID := c.Sess.GetStringDefault(config.UserIdKey, "")
	return userID
}
