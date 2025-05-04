package main

import (
	"todo-app/controllers"
	"todo-app/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/login", controllers.Login)

	// Auth gerektiren grup
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/todos", controllers.GetTodos)
		auth.POST("/todos", controllers.CreateTodo)
	}

	r.Run(":8080")
}
