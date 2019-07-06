package jsonmodel

//TODO: 前后端的数据交换
type RegisterForm struct {
	Name string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type TestForm struct{
	Email string `json:"email"`
	Password string `json:"password"`
}