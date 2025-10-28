package router

import (
	"github.com/elipe1/tasksGo-tarefa2/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	router.POST("/user", handler.CreateUserHandler)
	router.POST("/task", handler.CreateTaskHandler)
}
