package router

import (
	"github.com/elipe1/tasksGo-tarefa2/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	router.POST("/user", handler.CreateUserHandler)
	router.POST("/login", handler.LoginHandler)

	protected := router.Group("/")
	protected.Use(handler.AuthMiddleware())
	{
		protected.POST("/task", handler.CreateTaskHandler)
		// protected.GET("tasks", handler.ListTasksHandler)
		protected.PATCH("task/:id", handler.UpdateTaskHandler)
		protected.DELETE("task/:id", handler.DeleteTaskHandler)

		protected.PATCH("user/:id", handler.UpdateUserHandler)
		protected.GET("users", handler.ListUsersHandler)
		protected.DELETE("user/:id", handler.DeleteUserHandler)
	}
}
