package handler

import (
	"github.com/elipe1/tasksGo-tarefa2/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteTaskHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(400, gin.H{
			"error": "id parameter is required",
		})
		return
	}

	task := schemas.Tarefa{}
	if err := db.First(&task, id).Error; err != nil {
		ctx.JSON(404, gin.H{
			"error": "task not found with this ID",
		})
		return
	}

	if err := db.Delete(&task).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": "error deleting task with this ID",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "success deleting this task",
		"data":    task,
	})
}
