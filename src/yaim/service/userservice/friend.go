package userservice

import (
	"errors"
	"strconv"
	"yaim/model/jsonmodel"
	"yaim/model/ormmodel"
)

// 发起好友请求服务
func (service *Provider) AddFriendRequest(email, friendEmail string) error {
	if !service.Checkuser(friendEmail) {
		return errors.New("no such user")
	}

	if has, _ := service.CheckRequest(email, friendEmail); has {
		return errors.New("repeat request")
	}

	if email == friendEmail {
		return errors.New("?")
	}

	request := new(ormmodel.Friend)
	request.Adder = email
	request.Receiver = friendEmail

	_, err := service.engine.Insert(request)

	return err
}

func (service *Provider) HandleFriendRequest(email, friendEmail, action string) error {
	record := &ormmodel.Friend{Adder: friendEmail, Receiver: email}
	has, _ := service.engine.Get(record)

	if !has {
		return errors.New("you don't have such request")
	}
	if record.Validate != "" {
		return errors.New("already agreed")
	}

	var err error = nil
	if action == "yes" {
		record.Validate = "yes"
		_, err = service.engine.Id(record.Id).Update(record)
	} else if action == "no" {
		_, err = service.engine.Id(record.Id).Delete(record)
	} else {
		err = errors.New("invalid action")
	}

	return err
}

func (service *Provider) DeleteFriend(email, friendEmail string) error {
	has, record := service.CheckRequest(email, friendEmail)
	if !has {
		return errors.New("no such friend")
	}

	if record.Validate != "yes" {
		return errors.New("you two are not friends")
	}

	_, err := service.engine.Id(record.Id).Delete(record)
	return err
}

// 查询是否有好友请求记录, 并返回记录
func (service *Provider) CheckRequest(email, friendEmail string) (bool, *ormmodel.Friend) {
	record1 := &ormmodel.Friend{Adder: email, Receiver: friendEmail}
	record2 := &ormmodel.Friend{Adder: friendEmail, Receiver: email}
	has1, _ := service.engine.Get(record1)
	has2, _ := service.engine.Get(record2)

	if has1 {
		return has1, record1
	} else if has2 {
		return has2, record2
	} else {
		return false, nil
	}
}

// 获取好友列表
// email 从session中获取， 不存在注入问题
func (service *Provider) GetFriends(email string) ([]*jsonmodel.OnlineUserForm, []*jsonmodel.OfflineUserForm) {

	onlineQuery := `SELECT email, username, user.key, ip, port
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

	offlineQuery := `SELECT email, username, user.key
					FROM user
					WHERE (state = "Offline" or state = "")
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
	onlineUsers := make([]*jsonmodel.OnlineUserForm, 0)
	offlineUsers := make([]*jsonmodel.OfflineUserForm, 0)

	// result type []map[string][]byte
	onlineData, _ := service.engine.Query(onlineQuery)
	offlineData, _ := service.engine.Query(offlineQuery)

	for _, data := range onlineData {
		user := new(jsonmodel.OnlineUserForm)
		user.Email = string(data["email"])
		user.Username = string(data["username"])
		user.Key = string(data["key"])
		user.Ip = string(data["ip"])
		user.Port, _ = strconv.Atoi(string(data["port"]))

		onlineUsers = append(onlineUsers, user)
	}

	for _, data := range offlineData {
		user := new(jsonmodel.OfflineUserForm)
		user.Email = string(data["email"])
		user.Username = string(data["username"])
		user.Key = string(data["key"])

		offlineUsers = append(offlineUsers, user)
	}

	return onlineUsers, offlineUsers
}
