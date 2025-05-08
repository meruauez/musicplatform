package middlewares

import (
	"log"
	"net/http"
	"strings"
	"time"

	"musicplatform/music-service/handlers"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtKey = []byte("supersecret")

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &handlers.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("claims", claims)

		c.Next()
	}
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		requestID := uuid.New().String()

		start := time.Now()

		log.Printf("[%s] [RequestID: %s] %s %s - Started", time.Now().UTC().Format(time.RFC3339), requestID, c.Request.Method, c.Request.URL.Path)

		c.Next()

		duration := time.Since(start)
		statusCode := c.Writer.Status()

		log.Printf("[%s] [RequestID: %s] %s %s - %d - Duration: %v", time.Now().UTC().Format(time.RFC3339), requestID, c.Request.Method, c.Request.URL.Path, statusCode, duration)
	}
}
