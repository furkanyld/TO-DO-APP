package models

import "time"

// TO-DO Listesi
type TodoList struct {
	ID             int
	Name           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	CompletionRate float32
	Username       string // Listeyi oluşturan kullanıcı
}

// TO-DO Listesine ait adımlar
type TodoItem struct {
	ID        int
	ListID    int
	Content   string
	IsDone    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
