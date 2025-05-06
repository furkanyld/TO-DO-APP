package data

import (
	"testing"
	"time"
	"todo-app/models"
)

func TestAddTodoItem(t *testing.T) {
	list := AddTodoList(models.TodoList{
		Name:      "Item Test List",
		Username:  "testuser",
		CreatedAt: time.Now(),
	})

	item := models.TodoItem{
		ListID:    list.ID,
		Content:   "Test maddesi",
		IsDone:    false,
		CreatedAt: time.Now(),
		Username:  "testuser",
	}

	created := AddTodoItem(item)

	if created.ID == 0 {
		t.Error("Madde eklenirken ID atanmadı")
	}
	if created.Content != "Test maddesi" {
		t.Errorf("İçerik beklenenden farklı: %s", created.Content)
	}
	if created.ListID != list.ID {
		t.Errorf("ListID eşleşmiyor. Beklenen: %d, Gelen: %d", list.ID, created.ListID)
	}
}

func TestDeleteTodoItem_Unauthorized(t *testing.T) {
	list := AddTodoList(models.TodoList{
		Name:      "Silme Testi",
		Username:  "owner",
		CreatedAt: time.Now(),
	})

	item := AddTodoItem(models.TodoItem{
		ListID:    list.ID,
		Content:   "Silinecek Madde",
		IsDone:    false,
		CreatedAt: time.Now(),
		Username:  "owner",
	})

	err := DeleteTodoItem(item.ID, "unauthorized", "user")

	if err == nil {
		t.Error("Yetkisiz kullanıcı maddeyi silebildi, bu hatadır.")
	}
}
