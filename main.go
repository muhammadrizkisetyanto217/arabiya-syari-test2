package main

import (
	"arabiya-syari-api/config"
	"arabiya-syari-api/routes"
)

func main() {
	// Koneksi ke database
	config.ConnectDB()

	// Setup router
	r := routes.SetupRouter()

	// Jalankan server
	r.Run(":8080")
}
