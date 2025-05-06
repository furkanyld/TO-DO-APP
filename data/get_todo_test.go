
package data

import (
	"testing"
	"time"
	"todo-app/models"
)

func TestGetTodoItemsByListID(t *testing.T) {
	list := AddTodoList(models.TodoList{
		Name:      "GetItemList",
		Username:  "testuser",
		CreatedAt: time.Now(),
	})

	AddTodoItem(models.TodoItem{
		ListID:    list.ID,
		Content:   "Item 1",
		IsDone:    false,
		CreatedAt: time.Now(),
		Username:  "testuser",
	})

	AddTodoItem(models.TodoItem{
		ListID:    list.ID,
		Content:   "Item 2",
		IsDone:    false,
		CreatedAt: time.Now(),
		Username:  "testuser",
	})

	items := GetTodoItemsByListID(list.ID)
	if len(items) != 2 {
		t.Errorf("Beklenen madde sayısı: 2, gelen: %d", len(items))
	}
}

func TestGetUserTodoLists(t *testing.T) {
	AddTodoList(models.TodoList{
		Name:      "User1 Listesi",
		Username:  "user1",
		CreatedAt: time.Now(),
	})

	AddTodoList(models.TodoList{
		Name:      "User2 Listesi",
		Username:  "user2",
		CreatedAt: time.Now(),
	})

	user1Lists := GetUserTodoLists("user1")
	if len(user1Lists) != 1 {
		t.Errorf("user1 için beklenen liste sayısı: 1, gelen: %d", len(user1Lists))
	}

	if user1Lists[0].Username != "user1" {
		t.Errorf("Liste sahibi user1 olmalıydı, gelen: %s", user1Lists[0].Username)
	}
}
