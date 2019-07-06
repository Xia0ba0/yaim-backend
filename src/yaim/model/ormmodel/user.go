package ormmodel

type User struct {
	Email string `xorm:"pk 'email'"`
	Password string `xorm:"varchar(1024) notnull 'password'"`
	Username string `xorm:"notnull 'username'"`
	Key string `xorm:"varchar(4096) 'key'"`
	State string `xorm:"default 'offline' 'state'"`
	Ip string `xorm:"'ip'"`
	Port int32 `xorm:"port"`
	Validate string `xorm:"default 'no' 'validate'"`
}