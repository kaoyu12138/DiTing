package entity

import "time"

// Website 定义网址信息结构体
type Website struct {
	UrlId      int       `json:"urlId" gorm:"primaryKey" gorm:"autoIncrement" `
	Url        string    `json:"url"`
	Tag        string    `json:"tag"`
	RecordDate time.Time `json:"recordDate"`
}
