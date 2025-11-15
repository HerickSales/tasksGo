package handler

import (
	"github.com/elipe1/tasksGo-tarefa2/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateTaskHandler(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	isAdmin, _ := ctx.Get("isAdmin")

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

	isTaskOwner := task.ConcluinteID != nil && *task.ConcluinteID == userID.(uint)

	// Se NÃO é admin, NEM responsável pela tarefa
	if !isAdmin.(bool) && !isTaskOwner {
		ctx.JSON(403, gin.H{
			"error": "you don't have permission to update this task",
		})
		return
	}

	// Se É o responsável (mas NÃO é admin), só pode marcar como concluída
	if !isAdmin.(bool) && isTaskOwner {
		if request.Nome != "" || request.ConcluinteID != nil {
			ctx.JSON(403, gin.H{
				"error": "you can only mark this task as completed",
			})
			return
		}
		if request.Concluida != nil {
			task.Concluida = *request.Concluida
		}
	} else {
		if request.Nome != "" {
			task.Nome = request.Nome
		}

		if request.Concluida != nil {
			task.Concluida = *request.Concluida
		}

		if request.ConcluinteID != nil {
			task.ConcluinteID = request.ConcluinteID
		}
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
