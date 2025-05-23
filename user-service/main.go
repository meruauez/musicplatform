package main

import (
	"musicplatform/user-service/config"
	"musicplatform/user-service/handlers"
	"musicplatform/user-service/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Можно ограничить, например: []string{"http://localhost"}
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.Use(middlewares.LoggingMiddleware())

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	auth := r.Group("/")
	auth.Use(middlewares.JWTAuthMiddleware())

	auth.GET("/profile", handlers.GetCurrentUser)
	auth.PUT("/profile", handlers.UpdateCurrentUser)

	r.GET("/users/:id", handlers.GetUserByID)

	r.Run(":8081")
}
