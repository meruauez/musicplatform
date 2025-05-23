package main

import (
	"musicplatform/music-service/config"
	"musicplatform/music-service/handlers"
	"musicplatform/music-service/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost", "http://localhost:80"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.Use(middlewares.LoggingMiddleware())

	auth := r.Group("/")
	auth.Use(middlewares.JWTAuthMiddleware())

	auth.GET("/artists", handlers.GetArtists)
	auth.POST("/artists", handlers.CreateArtist)
	auth.DELETE("/artists/:id", handlers.DeleteArtist)
	auth.GET("/artists/:id", handlers.GetArtistByID)
	auth.GET("/artists/:id/songs", handlers.GetSongsByArtist)

	auth.GET("/genres", handlers.GetGenres)
	auth.POST("/genres", handlers.CreateGenre)
	auth.DELETE("/genres/:id", handlers.DeleteGenre)
	auth.GET("/genres/:id", handlers.GetGenreByID)
	auth.GET("/genres/:id/songs", handlers.GetSongsByGenre)

	auth.GET("/songs", handlers.GetSongs)
	auth.POST("/songs", handlers.CreateSong)
	auth.PUT("/songs/:id", handlers.UpdateSong)
	auth.DELETE("/songs/:id", handlers.DeleteSong)
	auth.GET("/songs/:id", handlers.GetSongByID)
	auth.GET("/songs/search", handlers.SearchSongs)

	r.GET("/songs/user/:userID", handlers.GetSongWithUser)

	r.Run(":8082")
}
