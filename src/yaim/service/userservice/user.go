package userservice

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"yaim/model/jsonmodel"
	"yaim/model/ormmodel"
)

type UserServiceProvider struct {
	engine *xorm.Engine
}

func NewProvider(engine *xorm.Engine) *UserServiceProvider {
	return &UserServiceProvider{
		engine: engine,
	}
}

func (service *UserServiceProvider) Adduser(reisteruser *jsonmodel.RegisterForm) error {
	if service.Checkuser(reisteruser.Email) {
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

func (service *UserServiceProvider) Verificate(email string) error {
	user := &ormmodel.User{Validate: "yes"}
	_, _ = service.engine.Id(email).Update(user)
	return nil
}

func (service *UserServiceProvider) CheckIdentity(email string, password string) error {
	if !service.Checkuser(email) {
		return errors.New("no such user")
	}

	sha256Inst := sha256.New()
	sha256Inst.Write([]byte(password))
	hashpassword := fmt.Sprintf("%x", sha256Inst.Sum([]byte("")))

	user := &ormmodel.User{Email: email, Password: hashpassword}
	has, _ := service.engine.Get(user)

	if !has {
		return errors.New("wrong password")
	}

	if user.Validate == "" {
		return errors.New("no email verification")
	}

	return nil
}

func (service *UserServiceProvider) CheckVerification(email string) bool{
	user := &ormmodel.User{Email: email}
	_, _ = service.engine.Get(user)

	return user.Validate != ""
}

func (service *UserServiceProvider) Checkuser(email string) bool {
	user := &ormmodel.User{Email: email}
	has, _ := service.engine.Get(user)
	return has
}
