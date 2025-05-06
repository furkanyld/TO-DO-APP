
package data

import (
	"testing"
	"time"
	"todo-app/models"
)

func TestGetUserTodoLists_ExcludesDeleted(t *testing.T) {
	list := AddTodoList(models.TodoList{
		Name:      "Silinecek Liste",
		Username:  "userA",
		CreatedAt: time.Now(),
	})

	// Soft delete
	now := time.Now()
	for i := range TodoLists {
		if TodoLists[i].ID == list.ID {
			TodoLists[i].DeletedAt = &now
		}
	}

	userLists := GetUserTodoLists("userA")
	if len(userLists) != 0 {
		t.Errorf("Soft delete edilmiş liste gösterilmemeliydi. Gelen: %d", len(userLists))
	}
}

func TestGetAllTodoLists_ExcludesDeleted(t *testing.T) {
	list := AddTodoList(models.TodoList{
		Name:      "Silinmiş Admin Listesi",
		Username:  "admin",
		CreatedAt: time.Now(),
	})

	// Soft delete
	now := time.Now()
	for i := range TodoLists {
		if TodoLists[i].ID == list.ID {
			TodoLists[i].DeletedAt = &now
		}
	}

	allLists := GetAllTodoLists()
	for _, l := range allLists {
		if l.ID == list.ID {
			t.Error("Soft delete edilmiş liste tüm listelerden çıkarılmamış.")
		}
	}
}
