package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-media-api/models"
	"strconv"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// models paketinin içindeki Users listesini encode ediyoruz
	json.NewEncoder(w).Encode(models.Users)

}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// models paketinin içindeki Post listesini encode ediyoruz
	json.NewEncoder(w).Encode(models.Posts)

}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Sadece Post metodu desteklenir", http.StatusMethodNotAllowed)
		return
	}

	var yeniKullanici models.User

	err := json.NewDecoder(r.Body).Decode(&yeniKullanici)

	if err != nil {
		http.Error(w, "Geçersiz JSON verisi", http.StatusBadRequest)
		return
	}

	yeniKullanici.ID = len(models.Users) + 1
	models.Users = append(models.Users, yeniKullanici)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(yeniKullanici)

}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Sadece POST metodu desteklenir", http.StatusMethodNotAllowed)
		return
	}

	var yeniPost models.Post

	err := json.NewDecoder(r.Body).Decode(&yeniPost)

	if err != nil {
		http.Error(w, "Geçersiz JSON verisi", http.StatusBadRequest)
		return
	}

	kullaniciVarMi := false

	for _, u := range models.Users {
		if u.ID == yeniPost.UserId {
			kullaniciVarMi = true
			break
		}
	}

	if !kullaniciVarMi { // yani false ise
		http.Error(w, "ID ile eşleşen kullanıcı bulunamadı", http.StatusBadRequest) // http.StatusBadRequest = 400
		return
	}

	yeniPost.ID = len(models.Posts) + 1
	yeniPost.LikeCount = 0

	models.Posts = append(models.Posts, yeniPost)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(yeniPost) // decoder okumaya encoder yazmaya yarıo

}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Sadece GET metodu desteklenir", http.StatusMethodNotAllowed)
		return
	}

	idString := r.URL.Query().Get("id")

	idSayi, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Geçersiz ID formatı, sayı olmalı", http.StatusBadRequest)
		return
	}

	postBulundu := false
	var bulunanPost models.Post

	for _, p := range models.Posts {
		if p.ID == idSayi {
			bulunanPost = p
			postBulundu = true
		}
	}

	if !postBulundu {
		http.Error(w, "Aradığınız post bulunamadı", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK) // 200
	json.NewEncoder(w).Encode(bulunanPost)

}

func LikePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Sadece POST metodu desteklenir", http.StatusMethodNotAllowed)
		return
	}

	stringId := r.URL.Query().Get("id")
	intId, err := strconv.Atoi(stringId)

	if err != nil {
		http.Error(w, "Lütfen sadece sayı giriniz", http.StatusMethodNotAllowed)
		return
	}

	postBulundu := false

	for i, p := range models.Posts {
		if p.ID == intId {
			postBulundu = true
			models.Posts[i].LikeCount++
			fmt.Fprintf(w, "Post Başarıyla Beğenildi")
			return
		}
	}

	if !postBulundu { // postBulundu false ise
		http.Error(w, "Post Bulunamadı", http.StatusBadRequest)
	}
}
