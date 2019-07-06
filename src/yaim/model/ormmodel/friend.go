package ormmodel

type Friend struct{
	Id int64
	Adder string `xorm:"notnull 'adder'"`
	Receiver string `xorm:"'receiver'"`
	Validate string `xorm:"default 'no' 'validate'"`
}