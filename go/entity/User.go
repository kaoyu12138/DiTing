package entity

// User 定义用户信息结构体
type User struct {
	Id        int    `json:"id" gorm:"primaryKey" gorm:"autoIncrement" `
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	DelStatus int    `json:"delStatus"`
}
