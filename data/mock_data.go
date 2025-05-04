package data

import "todo-app/models"

var Users = []models.User{
	{Username: "admin", Password: "admin123", Role: "admin"},
	{Username: "user", Password: "user123", Role: "user"},
}
