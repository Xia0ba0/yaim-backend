package userservice

import (
	"errors"
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
