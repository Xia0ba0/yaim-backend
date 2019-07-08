package userservice

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"yaim/model/jsonmodel"
	"yaim/model/ormmodel"
)

type Provider struct {
	engine *xorm.Engine
}

func NewProvider(engine *xorm.Engine) *Provider {
	return &Provider{
		engine: engine,
	}
}

// 添加用户服务
func (service *Provider) Adduser(reisteruser *jsonmodel.RegisterForm) error {
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

// 用户认证服务
func (service *Provider) Verificate(email string) error {
	user := &ormmodel.User{Validate: "yes"}
	_, err := service.engine.Id(email).Update(user)
	return err
}

// 更新公钥服务
func (service *Provider) Updatepubkey(email, pubkey string) error{
	user := &ormmodel.User{Key:pubkey}
	_, err := service.engine.Id(email).Update(user)
	return err
}

// 更新IP和端口服务
func (service *Provider) UpdateNetAddr(email string, ip string, port int) error{
	user := &ormmodel.User{Ip:ip, Port:port}
	_, err := service.engine.Id(email).Update(user)
	return err
}

// 改变用户在线状态
func (service *Provider) UpdateState(email string, state string) error{
	user := &ormmodel.User{State:state}
	_, err := service.engine.Id(email).Update(user)
	return err
}

// 用户认证服务
func (service *Provider) CheckIdentity(email string, password string) error {
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

// 用户认证检查服务
func (service *Provider) CheckVerification(email string) bool{
	user := &ormmodel.User{Email: email}
	_, _ = service.engine.Get(user)

	return user.Validate != ""
}

// 用户检索服务
func (service *Provider) Checkuser(email string) bool {
	user := &ormmodel.User{Email: email}
	has, _ := service.engine.Get(user)
	return has
}