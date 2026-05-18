package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Post struct {
	ID        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Body      string `json:"body"`
	LikeCount int    `json:"like_count"`
}

var Users []User
var Posts []Post
