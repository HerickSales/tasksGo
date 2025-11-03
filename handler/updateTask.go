package handler

import (
	"github.com/elipe1/tasksGo-tarefa2/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateTaskHandler(ctx *gin.Context) {
	request := UpdateTaskRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

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
			"error": "task not found",
		})
		return
	}

	if request.Nome != "" {
		task.Nome = request.Nome
	}

	if request.Concluida != nil {
		task.Concluida = *request.Concluida
	}

	if request.ConcluinteID != nil {
		task.ConcluinteID = request.ConcluinteID
	}

	if err := db.Save(&task).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed updating task",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "task updated successfully",
		"data":    task,
	})
}
