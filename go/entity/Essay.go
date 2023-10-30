package entity

import "time"

type Essay struct {
	EssayId      int       `json:"essayId"`
	LikeCount    int       `json:"likeCount"`
	UserName     string    `json:"userName"`
	UserAvatar   string    `json:"userAvatar"`
	EssayName    string    `json:"essayName"`
	EssayContent string    `json:"essayContent"`
	PublishDate  time.Time `json:"publishDate"`
}
