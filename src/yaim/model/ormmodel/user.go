package ormmodel

//TODO: 数据库的构造与ORM模型
type User struct {
	ID        int64
	Username  string
	Password  string
	Email     string
}