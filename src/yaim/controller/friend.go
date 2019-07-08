package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"yaim/config"
	"yaim/model/jsonmodel"
	"yaim/service/userservice"
)

type FriendController struct {
	//控制器依赖注入
	//动态绑定方式 每一个请求都有差异
	//注入的字段名必须大写
	Ctx  iris.Context
	Sess *sessions.Session

	UserService *userservice.UserServiceProvider
}

// Method POST
// Path /friend/request
// function: 发起添加好友请求
func (c *FriendController) PostRequest() {
	var requestForm jsonmodel.AddFriendForm

	if err := c.Ctx.ReadJSON(&requestForm); err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest) //400

		_, _ = c.Ctx.JSON(iris.Map{
			"message": "Error",
			"Error":   err.Error(),
		})
		return
	}

	userid := c.getuserid()
	if err := c.UserService.AddFriendRequest(userid, requestForm.Receiver); err != nil {
		_, _ = c.Ctx.JSON(iris.Map{
			"message": "Error",
			"Error":   err.Error(),
		})
		return
	}

	_, _ = c.Ctx.JSON(iris.Map{
		"message": "Success",
		"data":    "wait for agreement",
	})
}

func (c *FriendController) PostHandlerequest() {
	var handleForm jsonmodel.HandleAddFriendForm

	if err := c.Ctx.ReadJSON(&handleForm); err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest) //400

		_, _ = c.Ctx.JSON(iris.Map{
			"message": "Error",
			"Error":   err.Error(),
		})
		return
	}

	userid := c.getuserid()
	if err := c.UserService.HandleFriendRequest(userid, handleForm.Adder, handleForm.Action); err != nil {
		_, _ = c.Ctx.JSON(iris.Map{
			"message": "Error",
			"Error":   err.Error(),
		})
		return
	}

	_, _ = c.Ctx.JSON(iris.Map{
		"message": "Success",
		"data":    "action successes",
	})
}

func (c *FriendController) PostDelete() {
	var deleteForm jsonmodel.DeleteFriendForm

	if err := c.Ctx.ReadJSON(&deleteForm); err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest) //400

		_, _ = c.Ctx.JSON(iris.Map{
			"message": "Error",
			"Error":   err.Error(),
		})
		return
	}

	userid := c.getuserid()
	if err := c.UserService.DeleteFriend(userid, deleteForm.FriendEmail); err != nil {
		_, _ = c.Ctx.JSON(iris.Map{
			"message": "Error",
			"Error":   err.Error(),
		})
		return
	}

	_, _ = c.Ctx.JSON(iris.Map{
		"message": "Success",
		"data":    "delete friend successes",
	})
}

func (c *FriendController) GetGet() {
	userid := c.getuserid()
	onlineUsers, offlineUsers := c.UserService.GetFriends(userid)

	_, _ = c.Ctx.JSON(iris.Map{
		"message": "Success",
		"data": iris.Map{
			"onlineUsers":  onlineUsers,
			"offlineUsers": offlineUsers,
		},
	})
}

// 通过Session 获取用户id
func (c *FriendController) getuserid() string {
	userID := c.Sess.GetStringDefault(config.UserIdKey, "")
	return userID
}
