package controller

import (
	"fmt"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/websocket"
	"yaim/config"
	"yaim/service/pushservice"
	"yaim/service/userservice"
)

// Websocket连接非常持久 此控制器实例将持续存在
type PushController struct {
	//控制器依赖注入
	//动态绑定方式 每一个请求都有差异
	//注入的字段名必须大写
	Conn websocket.Connection
	Sess *sessions.Session

	//静态注入服务
	PushService *pushservice.Provider
	UserService *userservice.Provider
}

// websocket 连接只接受Get方法
func (c *PushController) Get(){
	userid := c.getuserid()

	// 用户上线 先改变用户状态
	// 然后注册连接 向好友推送上线消息
	_ = c.UserService.UpdateState(userid, "Online")
	c.PushService.Register(userid, &c.Conn)

	// 用户下线， 改变用户状态
	// 销毁连接 向好友推送下线消息
	_ = c.UserService.UpdateState(userid,"Offline")
	c.Conn.OnLeave(func(romName string){
		c.PushService.UnRegister(userid)
	})

	// 用户消息监听
	c.Conn.OnMessage(func (message []byte){
		_ = c.Conn.EmitMessage([]byte("hello world!"))
		fmt.Println(string(message))
	})


	c.Conn.Wait()
}

// 通过Session 获取用户id
func (c *PushController) getuserid() string {
	userID := c.Sess.GetStringDefault(config.UserIdKey, "")
	return userID
}