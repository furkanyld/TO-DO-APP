package middleware

import (
	"net/http"
	"strings"
	"todo-app/utils"

	"github.com/gin-gonic/gin"
)

// Bu middleware token'ı çözüp doğrular, hata varsa 401 döner
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token gerekli"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Hatalı token formatı"})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(tokenParts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token geçersiz"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}
