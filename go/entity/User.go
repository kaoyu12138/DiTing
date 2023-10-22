package entity

// User 定义用户信息结构体
type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
	DelStatus int    `json:"delStatus"`
}
