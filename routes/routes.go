package routes

import (
	"todo-app/controllers"
	"todo-app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.POST("/login", controllers.Login)
	protected := router.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())

	protected.GET("/getTodos", controllers.GetTodos)
	protected.POST("/createTodo", controllers.CreateTodo)
	protected.PUT("/updateTodos/:id", controllers.UpdateTodo)
	protected.DELETE("/deleteTodos/:id", controllers.DeleteTodo)
	protected.POST("/lists/:id/items", controllers.CreateTodoItem)
	protected.GET("/lists/:id/items", controllers.GetTodoItems)
	protected.PUT("/items/:id", controllers.UpdateTodoItem)
	protected.DELETE("/items/:id", controllers.DeleteTodoItem)
}
