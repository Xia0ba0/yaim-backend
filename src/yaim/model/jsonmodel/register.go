package jsonmodel

//TODO: 前后端的数据交换
type RegisterForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type TestForm struct{
	UserID string `json:"userid"`
}