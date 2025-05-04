package controllers

import (
	"net/http"
	"time"
	"todo-app/data"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

// TO-DO Listesi oluştur
func CreateTodo(c *gin.Context) {
	var input struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri"})
		return
	}

	username := c.GetString("username")

	list := models.TodoList{
		Name:           input.Name,
		Username:       username,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		CompletionRate: 0,
	}

	created := data.AddTodoList(list)
	c.JSON(http.StatusCreated, created)
}

// Kullanıcının TO-DO Listelerini getir
func GetTodos(c *gin.Context) {
	username := c.GetString("username")
	role := c.GetString("role")

	var lists []models.TodoList
	if role == "admin" {
		// admin tüm listeleri görür
		lists = data.TodoLists
	} else {
		lists = data.GetUserTodoLists(username)
	}

	c.JSON(http.StatusOK, lists)
}
