package entity

// Website 定义网址信息结构体
// website 包括 url + tag
type Website struct {
	UrlId int    `json:"urlId" gorm:"primaryKey" gorm:"autoIncrement" `
	Url   string `json:"url"`
	Tag   string `json:"tag"`
}
