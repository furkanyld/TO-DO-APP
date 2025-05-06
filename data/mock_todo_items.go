package data

import (
	"errors"
	"time"
	"todo-app/models"
)

var TodoItems []models.TodoItem
var itemIDCounter = 1

func AddTodoItem(item models.TodoItem) models.TodoItem {
	item.ID = itemIDCounter
	itemIDCounter++
	TodoItems = append(TodoItems, item)
	return item
}

func GetTodoItemsByListID(listID int) []models.TodoItem {
	var result []models.TodoItem
	for _, item := range TodoItems {
		if item.ListID == listID && item.DeletedAt == nil {
			result = append(result, item)
		}
	}
	return result
}

func UpdateTodoItem(id int, content string, isDone bool, username, role string) (models.TodoItem, error) {
	for i, item := range TodoItems {
		if item.ID == id && item.DeletedAt == nil {
			list, err := GetListByID(item.ListID)
			if err != nil {
				return models.TodoItem{}, err
			}
			if role != "admin" && list.Username != username {
				return models.TodoItem{}, errors.New("bu maddeyi güncelleme yetkiniz yok")
			}

			TodoItems[i].Content = content
			TodoItems[i].IsDone = isDone
			TodoItems[i].UpdatedAt = time.Now()
			return TodoItems[i], nil
		}
	}
	return models.TodoItem{}, errors.New("madde bulunamadı")
}

func DeleteTodoItem(id int, username, role string) error {
	for i, item := range TodoItems {
		if item.ID == id && item.DeletedAt == nil {
			list, err := GetListByID(item.ListID)
			if err != nil {
				return err
			}
			if role != "admin" && list.Username != username {
				return errors.New("bu maddeyi silme yetkiniz yok")
			}
			now := time.Now()
			TodoItems[i].DeletedAt = &now
			return nil
		}
	}
	return errors.New("madde bulunamadı")
}

func GetListByID(id int) (models.TodoList, error) {
	for _, list := range TodoLists {
		if list.ID == id && list.DeletedAt == nil {
			return list, nil
		}
	}
	return models.TodoList{}, errors.New("liste bulunamadı")
}
