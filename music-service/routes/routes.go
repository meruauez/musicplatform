package routes

import (
	"musicplatform/music-service/handlers"
	"musicplatform/music-service/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	auth := r.Group("/")
	auth.Use(middlewares.JWTAuthMiddleware())

	songGroup := r.Group("/songs")
	{
		songGroup.GET("/", handlers.GetSongs)
		songGroup.GET("/:id", handlers.GetSongByID)
		songGroup.POST("/", handlers.CreateSong)
		songGroup.PUT("/:id", handlers.UpdateSong)
		songGroup.DELETE("/:id", handlers.DeleteSong)
		songGroup.GET("/artist/:id", handlers.GetSongsByArtist)
		songGroup.GET("/genre/:id", handlers.GetSongsByGenre)
		songGroup.GET("/search", handlers.SearchSongs)
	}

	artistGroup := r.Group("/artists")
	{
		artistGroup.GET("/", handlers.GetArtists)
		artistGroup.GET("/:id", handlers.GetArtistByID)
		artistGroup.POST("/", handlers.CreateArtist)
		artistGroup.PUT("/:id", handlers.UpdateArtist)
		artistGroup.DELETE("/:id", handlers.DeleteArtist)
	}

	genreGroup := r.Group("/genres")
	{
		genreGroup.GET("/", handlers.GetGenres)
		genreGroup.GET("/:id", handlers.GetGenreByID)
		genreGroup.POST("/", handlers.CreateGenre)
		genreGroup.PUT("/:id", handlers.UpdateGenre)
		genreGroup.DELETE("/:id", handlers.DeleteGenre)
	}

	r.GET("/songs/user/:userID", handlers.GetSongWithUser)

	return r
}
