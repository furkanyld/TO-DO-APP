package controllers

import (
	"net/http"
	"strconv"
	"time"
	"todo-app/data"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	username := c.MustGet("username").(string)

	var input struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz giriş verisi"})
		return
	}

	newList := models.TodoList{
		Name:      input.Name,
		Username:  username,
		CreatedAt: time.Now(),
	}

	createdList := data.AddTodoList(newList)

	c.JSON(http.StatusCreated, gin.H{"todo": createdList})
}

func GetTodos(c *gin.Context) {
	username := c.MustGet("username").(string)
	role := c.MustGet("role").(string)

	var todos []models.TodoList
	if role == "admin" {
		todos = data.GetAllTodoLists()
	} else {
		todos = data.GetUserTodoLists(username)
	}

	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func UpdateTodo(c *gin.Context) {
	username := c.MustGet("username").(string)
	role := c.MustGet("role").(string)

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ID"})
		return
	}

	var input struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz giriş verisi"})
		return
	}

	updatedList, err := data.UpdateTodoList(id, input.Name, username, role)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todo": updatedList})
}

func DeleteTodo(c *gin.Context) {
	username := c.MustGet("username").(string)
	role := c.MustGet("role").(string)

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ID"})
		return
	}

	err = data.DeleteTodoList(id, username, role)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "TO-DO listesi silindi"})
}
