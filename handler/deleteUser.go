package handler

import (
	"github.com/elipe1/tasksGo-tarefa2/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(ctx *gin.Context) {
	isAdmin, exists := ctx.Get("isAdmin")
	if !exists {
		ctx.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}

	if !isAdmin.(bool) {
		ctx.JSON(401, gin.H{
			"error": "only admins can delete users",
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

	user := schemas.Usuario{}

	if err := db.Preload("TarefasDesignadas").First(&user, id).Error; err != nil {
		ctx.JSON(404, gin.H{
			"error": "user not found with this ID",
		})
		return
	}

	pendingTasks := 0
	for _, task := range user.TarefasDesignadas {
		if !task.Concluida {
			pendingTasks++
		}
	}

	if pendingTasks > 0 {
		ctx.JSON(400, gin.H{
			"error": "user has pendig tasks and cannot be deleted",
		})
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": "error deleting user with this ID",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "success deleting this user",
		"data":    user,
	})
}

// Ver ID do criador da tarefa (0)
