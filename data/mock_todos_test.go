package data

import (
	"testing"
	"time"
	"todo-app/models"
)

func TestAddTodoList(t *testing.T) {
	// Arrange
	list := models.TodoList{
		Name:      "Test List",
		Username:  "user",
		CreatedAt: time.Now(),
	}

	// Act
	result := AddTodoList(list)

	// Assert
	if result.ID == 0 {
		t.Errorf("Beklenen ID > 0, gelen ID: %d", result.ID)
	}
	if result.Name != "Test List" {
		t.Errorf("Liste adı beklenenden farklı: %s", result.Name)
	}
}

func TestUpdateTodoList_WithUnauthorizedUser(t *testing.T) {
	// Arrange: önce bir liste ekle
	list := AddTodoList(models.TodoList{
		Name:      "Yetki Testi",
		Username:  "ownerUser",
		CreatedAt: time.Now(),
	})

	// Act
	_, err := UpdateTodoList(list.ID, "Yeni Ad", "unauthorizedUser", "user")

	// Assert
	if err == nil {
		t.Error("Yetkisiz kullanıcı listeyi güncelleyebildi, bu hatadır.")
	}
}
