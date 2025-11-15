package handler

import (
	"github.com/elipe1/tasksGo-tarefa2/schemas"
	"github.com/gin-gonic/gin"
)

func ListUsersHandler(ctx *gin.Context) {
	users := []schemas.Usuario{}

	if err := db.Find(&users).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": "error listing users",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "success showing users",
		"data":    users,
	})
}
