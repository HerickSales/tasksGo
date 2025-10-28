package handler

import (
	"github.com/elipe1/tasksGo-tarefa2/schemas"
	"github.com/gin-gonic/gin"
)

func CreateTaskHandler(ctx *gin.Context) {
	request := CreateTaskRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	task := schemas.Tarefa{
		Nome:         request.Nome,
		CriadorID:    request.CriadorID,
		ConcluinteID: request.ConcluinteID,
	}

	if err := db.Create(&task).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "task created successfully",
		"data":    task,
	})

	// Apenas usuários responsaveis por tarefa x podem conclui-la (Verificação de id na hora de concluir tarefa)
}
