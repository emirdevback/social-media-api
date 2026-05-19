package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-media-api/database"
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

	for _, u := range models.Users {
		if u.Username == yeniKullanici.Username {
			http.Error(w, "Zaten bu ad ile kayıt olmus kullanıcı var lütfen farklı ad seçiniz", http.StatusBadRequest)
			return
		}
	}

	yeniKullanici.ID = len(models.Users) + 1
	models.Users = append(models.Users, yeniKullanici)

	database.SaveData()
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

	database.SaveData()

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

	database.SaveData()

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
			database.SaveData()
			return
		}
	}

	if !postBulundu { // postBulundu false ise
		http.Error(w, "Post Bulunamadı", http.StatusBadRequest)
	}
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Sadece POST metodu desteklenir", http.StatusMethodNotAllowed)
		return
	}

	var req models.DeletePostRequest
	err := json.NewDecoder(r.Body).Decode(&req) // r.url.query.get ile aynı mantık r.body urlyi &re e atıo & pointer demek rame kaydedio
	if err != nil {
		http.Error(w, "Geçersiz JSON verisi", http.StatusBadRequest)
		return
	}

	const gizliSifre = "captan123"

	if req.AdminPassword != gizliSifre {
		http.Error(w, "Hatalı Admin Şifresi! Erişim Engellendi.", http.StatusUnauthorized) // 401 Hatası
		return
	}

	stringId := r.URL.Query().Get("id")
	intId, err := strconv.Atoi(stringId)

	if err != nil {
		http.Error(w, "Lütfen geçerli bir post ID'si giriniz", http.StatusBadRequest)
		return
	}

	postBulundu := false

	for i, p := range models.Posts {
		if p.ID == intId {
			postBulundu = true
			models.Posts = append(models.Posts[:i], models.Posts[i+1:]...)
			fmt.Fprintf(w, "Admin, %d numaralı postu kaldırdı", intId)
			database.SaveData()
			return
		}
	}

	if !postBulundu { // yani false ise
		http.Error(w, "Silinmek istenen post numarası bulunamadı", http.StatusNotFound)
	}

}

func ChangePassword(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, " Sadece POST metodu desteklenir ", http.StatusMethodNotAllowed)
		return
	}

	var sifreDegisenKullanici models.ChangePasswordRequest

	err := json.NewDecoder(r.Body).Decode(&sifreDegisenKullanici)

	if err != nil {
		http.Error(w, " Doğru şekilde girdiğinizi kontrol edin ", http.StatusBadRequest)
		return
	}

	for i, u := range models.Users {
		if u.Username == sifreDegisenKullanici.Username {
			if u.Password == sifreDegisenKullanici.Password {
				models.Users[i].Password = sifreDegisenKullanici.New_password
				fmt.Fprintf(w, "Şifreniz başarıyla değiştirildi")
				database.SaveData()
				return
			} else {
				http.Error(w, "Hatalı Şifre ", http.StatusBadRequest)
				return
			}
		}
	}

	http.Error(w, "Kullanıcı Bulunamadı ", http.StatusNotFound)

}
