package controllers

import (
	"net/http"
	"strconv"
	"time"
	"todo-app/data"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

func CreateTodoItem(c *gin.Context) {
	username := c.MustGet("username").(string)

	listIDParam := c.Param("id")
	listID, err := strconv.Atoi(listIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz liste ID"})
		return
	}

	var input struct {
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz giriş verisi"})
		return
	}

	item := models.TodoItem{
		ListID:    listID,
		Content:   input.Content,
		IsDone:    false,
		CreatedAt: time.Now(),
		Username:  username,
	}

	created := data.AddTodoItem(item)
	c.JSON(http.StatusCreated, gin.H{"item": created})
}

func GetTodoItems(c *gin.Context) {
	listIDParam := c.Param("id")
	listID, err := strconv.Atoi(listIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz liste ID"})
		return
	}

	items := data.GetTodoItemsByListID(listID)
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func UpdateTodoItem(c *gin.Context) {
	username := c.MustGet("username").(string)
	role := c.MustGet("role").(string)

	itemIDParam := c.Param("id")
	itemID, err := strconv.Atoi(itemIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz madde ID"})
		return
	}

	var input struct {
		Content string `json:"content" binding:"required"`
		IsDone  bool   `json:"isDone"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz giriş verisi"})
		return
	}

	updated, err := data.UpdateTodoItem(itemID, input.Content, input.IsDone, username, role)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"item": updated})
}

func DeleteTodoItem(c *gin.Context) {
	username := c.MustGet("username").(string)
	role := c.MustGet("role").(string)

	itemIDParam := c.Param("id")
	itemID, err := strconv.Atoi(itemIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz madde ID"})
		return
	}

	err = data.DeleteTodoItem(itemID, username, role)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "TO-DO maddesi silindi"})
}
