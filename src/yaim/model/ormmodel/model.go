package ormmodel

type User struct {
	Email string `xorm:"pk 'email'"`
	Password string `xorm:"varchar(1024) notnull 'password'"`
	Username string `xorm:"notnull 'username'"`
	Key string `xorm:"varchar(4096) 'key'"`
	State string `xorm:"default 'offline' 'state'"`
	Ip string `xorm:"'ip'"`
	Port int `xorm:"port"`
	Validate string `xorm:"default 'no' 'validate'"`
}

type Friend struct{
	Id int64
	Adder string `xorm:"notnull 'adder'"`
	Receiver string `xorm:"'receiver'"`
	Validate string `xorm:"default 'no' 'validate'"`
}

type Offline struct{
	Id int64
	Type string `xorm:"notnull 'type'"`
	Sender string `xorm:"notnull 'sender'"`
	Receiver string `xorm:"notnull 'receiver'"`
	Content string `xorm:"varchar(4096)"`
}