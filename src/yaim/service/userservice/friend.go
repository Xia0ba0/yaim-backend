package userservice

import (
	"errors"
	"fmt"
	"yaim/model/ormmodel"
)

// 发起好友请求服务
func (service *UserServiceProvider) AddFriendRequest(email, friendEmail string) error {
	if !service.Checkuser(friendEmail) {
		return errors.New("no such user")
	}

	if has, _:= service.CheckRequest(email, friendEmail); has {
		return errors.New("repeat request")
	}

	if email==friendEmail{
		return errors.New("?")
	}

	request := new(ormmodel.Friend)
	request.Adder = email
	request.Receiver = friendEmail

	_, err := service.engine.Insert(request)

	return err
}

func (service *UserServiceProvider) HandleFriendRequest(email, friendEmail, action string) error {
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
	}else if action == "no"{
		_, err = service.engine.Id(record.Id).Delete(record)
	}else{
		err = errors.New("invalid action")
	}

	return err
}

func (service *UserServiceProvider) DeleteFriend(email, friendEmail string) error{
	has, record := service.CheckRequest(email, friendEmail)
	if !has{
		return errors.New("no such friend")
	}

	if record.Validate != "yes"{
		return errors.New("you two are not friends")
	}

	_, err := service.engine.Id(record.Id).Delete(record)
	return err
}

// 查询是否有好友请求记录, 并返回记录
func (service *UserServiceProvider) CheckRequest(email, friendEmail string) (bool, *ormmodel.Friend){
	record1 := &ormmodel.Friend{Adder: email, Receiver: friendEmail}
	record2 := &ormmodel.Friend{Adder: friendEmail, Receiver: email}
	has1, _ := service.engine.Get(record1)
	has2, _ := service.engine.Get(record2)

	if has1{
		return has1, record1
	}else if has2{
		return has2, record2
	}else{
		return false, nil
	}
}


// 获取好友列表
// email 从session中获取， 不存在注入问题
func (service *UserServiceProvider) GetFriends(email string)(){
	offlineQuery := `SELECT email, username, key
					FROM user
					WHERE (state = 'no' or state = '')
							AND
							(email in (
										SELECT receiver 
										FROM friend
										WHERE adder = "` + email +`" and validate="yes"
										)
								OR
							email in (
										SELECT adder
										FROM friend
										WHERE receiver = "` + email +`" and validate ="yes"
									)
							)`

	onlienQuery := `SELECT email, username, key
					FROM user
					WHERE (state = 'yes')
							AND
							(email in (
										SELECT receiver 
										FROM friend
										WHERE adder = "` + email +`" and validate="yes"
										)
								OR
							email in (
										SELECT adder
										FROM friend
										WHERE receiver = "` + email +`" and validate ="yes"
									)
							)`

	offline, err := service.engine.Query(offlineQuery)
	online, err := service.engine.Query(onlienQuery)

	if err!= nil{
		fmt.Println()
	}
}