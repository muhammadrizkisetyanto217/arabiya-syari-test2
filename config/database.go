package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Load file .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Tidak bisa load .env file, pakai env default")
	}

	// Ambil URL database dari environment variable
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("DB_URL tidak ditemukan di environment variable")
	}

	// Koneksi ke PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	log.Println("Database Railway berhasil terhubung!")
	DB = db
}
