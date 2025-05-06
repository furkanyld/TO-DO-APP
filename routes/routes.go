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
	protected.PUT("/updateTodo/:id", controllers.UpdateTodo)
	protected.DELETE("/deleteTodo/:id", controllers.DeleteTodo)
	protected.POST("/lists/:id/addItem", controllers.CreateTodoItem)
	protected.GET("/lists/:id/getItems", controllers.GetTodoItems)
	protected.PUT("/updateItem/:id", controllers.UpdateTodoItem)
	protected.DELETE("/deleteItem/:id", controllers.DeleteTodoItem)

}
