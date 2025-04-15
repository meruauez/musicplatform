package routes

import (
	"musicplatform/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Маршруты для песен
	songGroup := r.Group("/songs")
	{
		songGroup.GET("/", handlers.GetSongs)
		songGroup.GET("/:id", handlers.GetSongByID)
		songGroup.POST("/", handlers.CreateSong)
		songGroup.PUT("/:id", handlers.UpdateSong)
		songGroup.DELETE("/:id", handlers.DeleteSong)
	}

	// Маршруты для артистов
	artistGroup := r.Group("/artists")
	{
		artistGroup.GET("/", handlers.GetArtists)
		artistGroup.GET("/:id", handlers.GetArtistByID)
		artistGroup.POST("/", handlers.CreateArtist)
		artistGroup.PUT("/:id", handlers.UpdateArtist)
		artistGroup.DELETE("/:id", handlers.DeleteArtist)
	}

	// Маршруты для жанров
	genreGroup := r.Group("/genres")
	{
		genreGroup.GET("/", handlers.GetGenres)
		genreGroup.GET("/:id", handlers.GetGenreByID)
		genreGroup.POST("/", handlers.CreateGenre)
		genreGroup.PUT("/:id", handlers.UpdateGenre)
		genreGroup.DELETE("/:id", handlers.DeleteGenre)
	}

	return r
}
