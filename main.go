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

	// safe routes
	auth := r.Group("/")
	auth.Use(middlewares.JWTAuthMiddleware())

	//Artists
	auth.GET("/artists", handlers.GetArtists)
	auth.POST("/artists", handlers.CreateArtist)
	auth.DELETE("/artists/:id", handlers.DeleteArtist)
	auth.GET("/artists/:id", handlers.GetArtistByID)

	// Genre
	auth.GET("/genres", handlers.GetGenres)
	auth.POST("/genres", handlers.CreateGenre)
	auth.DELETE("/genres/:id", handlers.DeleteGenre)
	auth.GET("/genres/:id", handlers.GetGenreByID)

	//Songs
	auth.GET("/songs", handlers.GetSongs) // Получить все песни с пагинацией и фильтром
	auth.POST("/songs", handlers.CreateSong)
	auth.PUT("/songs/:id", handlers.UpdateSong)
	auth.DELETE("/songs/:id", handlers.DeleteSong)
	auth.GET("/songs/:id", handlers.GetSongByID)

	r.Run(":8080")
}
