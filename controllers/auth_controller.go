package controllers

import (
	"net/http"
	"todo-app/data"
	"todo-app/utils"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Kullanıcı adı/parola kontrolü ve JWT üretimi
func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz giriş verisi"})
		return
	}

	// Kullanıcı kontrolü
	for _, user := range data.Users {
		if input.Username == user.Username && input.Password == user.Password {
			// JWT oluştur
			token, err := utils.GenerateJWT(user.Username, user.Role)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Token üretilemedi"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"token": token})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Geçersiz kullanıcı adı veya parola"})
}
