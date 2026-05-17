package main

import (
	"fmt"
	"net/http"
	"social-media-api/handlers"
)

func main() {

	http.HandleFunc("/users", handlers.UsersHandler)
	http.HandleFunc("/posts", handlers.PostHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)

	fmt.Println("Sosyal Medya Backend Motoru Paket Mantığıyla Açıldı: http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
