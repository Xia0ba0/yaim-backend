package jsonmodel

//TODO: 前后端的数据交换
type RegisterForm struct {
	ID       int64
	Username string
	Password string
	Email    string
}

type TestForm struct{
	UserID int64 `json:"number"`
}