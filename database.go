package database

import (
	"encoding/json"
	"fmt"
	"os"
	"social-media-api/models"
)

const (
	usersFile = "database/users.json"
	postsFile = "database/posts.json"
)

func SaveData() {

	// kullanıcıları json a çeviriyoruz ve diske yazıyoruz
	userBytes, err := json.MarshalIndent(models.Users, "", "   ")

	if err == nil {
		_ = os.WriteFile(usersFile, userBytes, 0644)
	}

	// postları json a çeviriyoruz ve diske yazıyoruz
	postsBytes, err := json.MarshalIndent(models.Posts, "", "   ")
	if err == nil {
		_ = os.WriteFile(postsFile, postsBytes, 0644)
	}

	fmt.Println("💾 [DATABASE] Tüm veriler başarıyla diske çivilendi!")

}

func LoadData() {

	// kullanıcı dosyalarını okuyup yüklüyoruz
	usersBytes, err := os.ReadFile(usersFile)

	if err == nil {
		_ = json.Unmarshal(usersBytes, &models.Users)
	}

	// post dosyalarını okuyup yüklüyoruz
	postsBytes, err := os.ReadFile(postsFile)

	if err == nil {
		_ = json.Unmarshal(postsBytes, &models.Posts)
	}

	fmt.Printf("📂 [DATABASE] Eski veriler yüklendi! (%d Kullanıcı, %d Post aktif)\n", len(models.Users), len(models.Posts))
}
