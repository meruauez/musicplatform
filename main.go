package main

import (
	"musicplatform/config"
	"musicplatform/handlers"
	"musicplatform/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()

	// Auth endpoints
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Защищённые маршруты
	auth := r.Group("/")
	auth.Use(middlewares.JWTAuthMiddleware())

	// Роуты для артистов
	auth.GET("/artists", handlers.GetArtists)          // Получить всех артистов
	auth.POST("/artists", handlers.CreateArtist)       // Создать нового артиста
	auth.DELETE("/artists/:id", handlers.DeleteArtist) // Удалить артиста по ID
	auth.GET("/artists/:id", handlers.GetArtistByID)   // Получить артиста по ID

	// Роуты для жанров
	auth.GET("/genres", handlers.GetGenres)          // Получить все жанры
	auth.POST("/genres", handlers.CreateGenre)       // Создать новый жанр
	auth.DELETE("/genres/:id", handlers.DeleteGenre) // Удалить жанр по ID
	auth.GET("/genres/:id", handlers.GetGenreByID)   // Получить жанр по ID

	// Роуты для песен с пагинацией и фильтрами
	auth.GET("/songs", handlers.GetSongs)          // Получить все песни с пагинацией и фильтром
	auth.POST("/songs", handlers.CreateSong)       // Создать новую песню
	auth.PUT("/songs/:id", handlers.UpdateSong)    // Обновить песню по ID
	auth.DELETE("/songs/:id", handlers.DeleteSong) // Удалить песню по ID
	auth.GET("/songs/:id", handlers.GetSongByID)   // Получить песню по ID

	// Запускаем сервер
	r.Run(":8080")
}
