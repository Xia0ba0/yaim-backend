package pushservice

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/websocket"
)

type Provider struct {
	connections map[string]*websocket.Connection
	engine *xorm.Engine
}

func NewProvider(engine *xorm.Engine) *Provider {
	return &Provider{
		connections: make(map[string]*websocket.Connection),
		engine:engine,
	}
}

func (service *Provider) Register(email string, conn *websocket.Connection) {
	service.connections[email] = conn

	// 向在线好友推送上线消息
	friends := service.GetFriends(email)
	for _, friend := range friends{
		fmt.Println(friend)

		// go语言的坑 多写几步
		connectionPointer := service.connections[friend]
		connection := *connectionPointer
		_ = connection.EmitMessage([]byte(""))
	}
}

func (service *Provider) UnRegister(email string) {
	delete(service.connections, email)

	friends := service.GetFriends(email)
	for _, friend := range friends{
		// go语言的坑 多写几步
		connectionPointer := service.connections[friend]
		connection := *connectionPointer
		_ = connection.EmitMessage([]byte(""))
	}
}


func (service *Provider) GetFriends(email string) []string {

	onlineQuery := `SELECT email
					FROM user
					WHERE (state = "Online")
						AND
							(email in (
										SELECT receiver 
										FROM friend
										WHERE adder = "` + email + `" and validate="yes"
										)
								OR
							email in (
										SELECT adder
										FROM friend
										WHERE receiver = "` + email + `" and validate ="yes"
									)
							)`

	onlineFriends := make([]string, 0)

	// result type []map[string][]byte
	onlineData, _ := service.engine.Query(onlineQuery)

	for _, data := range onlineData {
		onlineFriends = append(onlineFriends, string(data["email"]))
	}

	return onlineFriends
}