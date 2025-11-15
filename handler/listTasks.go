package handler

import (
	"github.com/elipe1/tasksGo-tarefa2/schemas"
	"github.com/gin-gonic/gin"
)

func ListTasksHandler(ctx *gin.Context) {
	tasks := []schemas.Tarefa{}

	if err := db.Find(&tasks).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": "error listing tasks",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "success showing tasks",
		"data":    tasks,
	})
}
