package request

type LoginStruct struct {
	Phone string `json:"phone"`
	Password string `json:"password"`
}

type RegisterStruct struct {
	Phone string `json:"phone"`
	Username string `json:"username"`
	Avatar string `json:"avatar"`
	Password string `json:"password"`
}

type GetInfoStruct struct {
	Id int `json:"id"`
	Phone string `json:"phone"`
	Username string `json:"username"`
}