package models

import "time"

type Post struct {
	ID          int64     `json:"id,string" db:"post_id"` //json:"id,string"序列化时将Int64转化为string,反之同理
	AuthorID    int64     `json:"author_id,string" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

type ApiPostDetail struct {
	AuthorName    string `json:"author_name" db:"username"`
	Title         string `json:"title" db:"title"`
	Content       string `json:"content" db:"content"`
	CommunityName string `json:"community_name" db:"community_name"`
}
