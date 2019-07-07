package config

import "time"

const (
	ServerAddr = "http://localhost:8090"
	HostName = "http://localhost:9080"
	Port     = ":8090"

	CookieName    = "YaimSession"
	CookieExpires = 24 * time.Hour
	UserIdKey     = "userid"

	SMTPServer   = "smtp.163.com:25"
	SMTPAccount  = "m18569002382@163.com"
	SMTPPassword = "4702391byl"
	SMTPSubject = "Yaim Account Verification"

	TokenKey = "Yaim?:@$%"
)
