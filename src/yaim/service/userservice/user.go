package userservice

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"yaim/model/jsonmodel"
	"yaim/model/ormmodel"
)

type UserServiceProvider struct{
	engine *xorm.Engine
}

func NewProvider(engine *xorm.Engine) *UserServiceProvider {
	return &UserServiceProvider{
		engine:engine,
	}
}

func (service *UserServiceProvider) Adduser(reisteruser *jsonmodel.RegisterForm) error{
	if service.Checkuser(reisteruser.Email){
		return errors.New("already exists")
	}

	user := new(ormmodel.User)
	user.Email = reisteruser.Email
	user.Username = reisteruser.Name


	sha256Inst := sha256.New()
	sha256Inst.Write([]byte(reisteruser.Password))
	user.Password = fmt.Sprintf("%x", sha256Inst.Sum([]byte("")))

	_, err := service.engine.Insert(user)

	return err
}

func (service *UserServiceProvider) Checkuser(email string) bool{
	user := &ormmodel.User{Email:email}
	has, _ := service.engine.Get(user)
	return has
}