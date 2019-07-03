package controller

import (
	"fmt"
	"github.com/kataras/iris/websocket"
)

// Websocket连接非常持久 此控制器实例将持续存在
type PushController struct {
	//控制器依赖注入
	//动态绑定方式 每一个请求都有差异
	//注入的字段名必须大写
	Conn websocket.Connection
}


// websocket 连接只接受Get方法
func (c *PushController) Get(){
	c.Conn.OnPong(func(){
		fmt.Println("a client has connected")
	})
	c.Conn.OnLeave(func(romName string){
		fmt.Println("a client has disconnected")
	})
	c.Conn.Wait()
}