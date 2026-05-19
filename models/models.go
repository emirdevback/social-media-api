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

type DeletePostRequest struct {
	AdminPassword string `json:"admin_password"`
}

type ChangePasswordRequest struct {
	Username     string `json:"nusername"`
	Password     string `json:n_password`
	New_password string `json:"new_password"`
}

var Users []User
var Posts []Post
