package main

import (
	"musicplatform/config"
	"musicplatform/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Подключаемся к базе данных
	config.ConnectDB()

	// Инициализируем роутер
	r := gin.Default()

	// Роуты для артистов
	r.GET("/artists", handlers.GetArtists)
	r.POST("/artists", handlers.CreateArtist)
	r.DELETE("/artists/:id", handlers.DeleteArtist)

	// Роуты для жанров
	r.GET("/genres", handlers.GetGenres)
	r.POST("/genres", handlers.CreateGenre)
	r.DELETE("/genres/:id", handlers.DeleteGenre)

	// Роуты для песен с пагинацией и фильтрами
	r.GET("/songs", handlers.GetSongs)
	r.POST("/songs", handlers.CreateSong)
	r.PUT("/songs/:id", handlers.UpdateSong)
	r.DELETE("/songs/:id", handlers.DeleteSong)

	// Запускаем сервер
	r.Run(":8080")
}
