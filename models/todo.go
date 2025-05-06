package models

import "time"

type TodoList struct {
	ID             int
	Name           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	CompletionRate float32
	Username       string // Creator
}

type TodoItem struct {
	ID        int
	ListID    int
	Content   string
	IsDone    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Username  string // Creator
}
