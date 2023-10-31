package entity

import "time"

type Essay struct {
	EssayId      int       `json:"essayId" gorm:"primaryKey" gorm:"autoIncrement" `
	LikeCount    int       `json:"likeCount"`
	UserName     string    `json:"userName"`
	UserAvatar   string    `json:"userAvatar"`
	EssayName    string    `json:"essayName"`
	EssayContent string    `json:"essayContent"`
	EssayAvatar  string    `json:"essayAvatar"`
	PublishDate  time.Time `json:"publishDate"`
}
