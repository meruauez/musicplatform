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

	// Защищенные маршруты, требующие JWT авторизации
	auth := r.Group("/")
	auth.Use(middlewares.JWTAuthMiddleware())

	// Профиль текущего пользователя
	auth.GET("/profile", handlers.GetCurrentUser)
	auth.PUT("/profile", handlers.UpdateCurrentUser)

	// Работа с артистами
	auth.GET("/artists", handlers.GetArtists)
	auth.POST("/artists", handlers.CreateArtist)
	auth.DELETE("/artists/:id", handlers.DeleteArtist)
	auth.GET("/artists/:id", handlers.GetArtistByID)
	// Дополнительный эндпоинт: Получить все песни артиста
	auth.GET("/artists/:id/songs", handlers.GetSongsByArtist)

	// Работа с жанрами

	// Genre
	auth.GET("/genres", handlers.GetGenres)
	auth.POST("/genres", handlers.CreateGenre)
	auth.DELETE("/genres/:id", handlers.DeleteGenre)
	auth.GET("/genres/:id", handlers.GetGenreByID)
	// Дополнительный эндпоинт: Получить все песни в жанре
	auth.GET("/genres/:id/songs", handlers.GetSongsByGenre)

	// Работа с песнями
	auth.GET("/songs", handlers.GetSongs) // Получить все песни с пагинацией и фильтром
	auth.POST("/songs", handlers.CreateSong)
	auth.PUT("/songs/:id", handlers.UpdateSong)
	auth.DELETE("/songs/:id", handlers.DeleteSong)
	auth.GET("/songs/:id", handlers.GetSongByID)
	// Дополнительный эндпоинт: Получить все песни по ключевому слову в названии
	auth.GET("/songs/search", handlers.SearchSongs)

	r.Run(":8080")
}
