package middlewares

import (
	"net/http"
	"strings"

	"musicplatform/handlers"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("supersecret")

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Извлекаем токен из заголовка Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			c.Abort()
			return
		}

		// Извлекаем сам токен
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Создаем структуру claims, куда будут записаны данные из токена
		claims := &handlers.Claims{}

		// Разбираем токен и извлекаем claims
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		// Проверяем на ошибки или недействительный токен
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Если все в порядке, сохраняем claims в контексте
		c.Set("claims", claims)

		// Переход к следующему обработчику
		c.Next()
	}
}
