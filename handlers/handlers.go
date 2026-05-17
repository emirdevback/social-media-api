package handlers

import (
	"encoding/json"
	"net/http"
	"social-media-api/models"
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
