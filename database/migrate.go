package database

import (
	"arabiya-syari-api/config"
	"arabiya-syari-api/model"
	"fmt"
)

func Migrate() {
	config.ConnectDB()

	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("Gagal migrasi:", err)
	} else {
		fmt.Println("Migrasi berhasil!")
	}
}
