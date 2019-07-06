package ormmodel

type Offline struct{
	Id int64
	Type string `xorm:"notnull 'type'"`
	Sender string `xorm:"notnull 'sender'"`
	Receiver string `xorm:"notnull 'receiver'"`
	Content string `xorm:"varchar(4096)"`
}