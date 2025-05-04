package data

import (
	"todo-app/models"
)

var TodoLists []models.TodoList
var TodoItems []models.TodoItem

var listIDCounter = 1
var itemIDCounter = 1

// Yeni bir TO-DO listesi ekler
func AddTodoList(list models.TodoList) models.TodoList {
	list.ID = listIDCounter
	listIDCounter++
	TodoLists = append(TodoLists, list)
	return list
}

// Kullanıcıya ait TO-DO listelerini döner
func GetUserTodoLists(username string) []models.TodoList {
	var result []models.TodoList
	for _, list := range TodoLists {
		if list.Username == username && list.DeletedAt == nil {
			result = append(result, list)
		}
	}
	return result
}
