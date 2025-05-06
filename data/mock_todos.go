package data

import (
	"errors"
	"time"
	"todo-app/models"
)

var TodoLists []models.TodoList
var listIDCounter = 1

func AddTodoList(list models.TodoList) models.TodoList {
	list.ID = listIDCounter
	listIDCounter++
	TodoLists = append(TodoLists, list)
	return list
}

func GetAllTodoLists() []models.TodoList {
	var result []models.TodoList
	for _, list := range TodoLists {
		if list.DeletedAt == nil {
			result = append(result, list)
		}
	}
	return result
}

func GetUserTodoLists(username string) []models.TodoList {
	var result []models.TodoList
	for _, list := range TodoLists {
		if list.Username == username && list.DeletedAt == nil {
			result = append(result, list)
		}
	}
	return result
}

func UpdateTodoList(id int, name, username, role string) (models.TodoList, error) {
	for i, list := range TodoLists {
		if list.ID == id && list.DeletedAt == nil {
			if role != "admin" && list.Username != username {
				return models.TodoList{}, errors.New("bu listeyi güncelleme yetkiniz yok")
			}
			TodoLists[i].Name = name
			now := time.Now()
			TodoLists[i].UpdatedAt = now
			return TodoLists[i], nil
		}
	}
	return models.TodoList{}, errors.New("liste bulunamadı")
}

func DeleteTodoList(id int, username, role string) error {
	for i, list := range TodoLists {
		if list.ID == id && list.DeletedAt == nil {
			if role != "admin" && list.Username != username {
				return errors.New("bu listeyi silme yetkiniz yok")
			}
			now := time.Now()
			TodoLists[i].DeletedAt = &now
			return nil
		}
	}
	return errors.New("liste bulunamadı")
}
