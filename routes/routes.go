package routes

import (
	"arabiya-syari-api/config"
	"arabiya-syari-api/middleware"
	"arabiya-syari-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Pasang Middleware CORS
	r.Use(middleware.SetupCORS())

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
