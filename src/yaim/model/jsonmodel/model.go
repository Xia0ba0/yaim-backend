package jsonmodel

type RegisterForm struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UploadKeyForm struct {
	Key string `json:"key"`
}

type UpdateAddrForm struct {
	Ip   string `json:"ip"`
	Port int32  `json:"port"`
}

type AddFriendForm struct {
	Receiver string `json:"receiver"`
}

type HandleAddFriendForm struct {
	Adder  string `json:"adder"`
	Action string `json:"action"`
}

type DeleteFriendForm struct{
	FriendEmail string `json:"friendEmail"`
}