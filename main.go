package main

import (
	"fmt"
	"net/http"
	"social-media-api/handlers"
)

func main() {

	http.HandleFunc("/users", handlers.UsersHandler)
	http.HandleFunc("/posts", handlers.PostHandler)
	http.HandleFunc("/post", handlers.GetPostHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/create-post", handlers.CreatePostHandler)
	http.HandleFunc("/like-post", handlers.LikePostHandler)

	fmt.Println("Sosyal Medya Backend Motoru Paket Mantığıyla Açıldı: http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
