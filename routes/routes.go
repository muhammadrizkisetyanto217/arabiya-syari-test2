package routes

import (
	"arabiya-syari-api/config"
	// "arabiya-syari-api/middleware"
	"arabiya-syari-api/model"
	"net/http"

	"github.com/rs/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Middleware CORS
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Izinkan semua origin (bisa diganti dengan domain tertentu)
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Gunakan CORS Middleware di Gin
	r.Use(func(c *gin.Context) {
		corsMiddleware.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Next()
		})).ServeHTTP(c.Writer, c.Request)
	})

	// Route sederhana
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Gin + GORM + PostgreSQL + CORS"})
	})


	// Route untuk tambah user
	r.POST("/users", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		config.DB.Create(&user)
		c.JSON(http.StatusCreated, user)
	})

	// Route untuk ambil semua user
	r.GET("/users", func(c *gin.Context) {
		var users []models.User
		config.DB.Find(&users)
		c.JSON(http.StatusOK, users)
	})

	return r
}
