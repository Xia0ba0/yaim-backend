package mailservice

import (
	"crypto/md5"
	"fmt"
	"net/smtp"
	"net/url"
	"yaim/config"
)

type MailServiceProvider struct {
	smtpHost string
	smtpAcount string
	smtpPassword string
	smtpSubject string
}

func NewProvider(smtpHost, smtpAcount, smtpPassword, smtpSubject string) *MailServiceProvider{
	return &MailServiceProvider{
		smtpHost:smtpHost,
		smtpAcount:smtpAcount,
		smtpPassword:smtpPassword,
		smtpSubject:smtpSubject,
	}
}

func (service *MailServiceProvider) SendToken(to string){
	token:= config.ServerAddr + "/user/verification"
	token += "?user=" + url.QueryEscape(to)

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(to))
	cipherStr := fmt.Sprintf("%x",md5Ctx.Sum([]byte(config.TokenKey)))

	token += "&token=" + url.QueryEscape(cipherStr)
	service.Send(to, token)
}

func (service *MailServiceProvider) Send(to, body string) {
	/*for server to run*/
	auth := &loginAuth{service.smtpAcount, service.smtpPassword}

	html := `<html><body><h3>`
	html += body
	html += `</h3></body></html>`

	content := "To: " + to + "\r\n"
	content += "From: " + service.smtpAcount + "\r\n"
	content += "Subject: "+ service.smtpSubject + "\r\n"
	content += "Content-Type: text/html; charset=UTF-8"
	content += "\r\n\r\n"
	content += html

	msg := []byte(content)
	go smtp.SendMail(service.smtpHost, auth, service.smtpAcount, []string{to}, msg)
}


/*
	下面的代码没有具体功能
	仅仅为了能够让MailServer跑起来
*/
type loginAuth struct {
	username, password string
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	// return "LOGIN", []byte{}, nil
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		}
	}
	return nil, nil
}
