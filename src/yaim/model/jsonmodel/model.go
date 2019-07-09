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

type OnlineUserForm struct{
	Email string `json:"email"`
	Username string `json:"username"`
	Key string `json:"key"`
	Ip string `json:"ip"`
	Port int `json:"port"`
}

type OfflineUserForm struct{
	Email string `json:"email"`
	Username string `json:"username"`
	Key string `json:"key"`
}