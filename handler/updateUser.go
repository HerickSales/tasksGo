package handler

import (
	"github.com/elipe1/tasksGo-tarefa2/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateUserHandler(ctx *gin.Context) {
	request := UpdateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(400, gin.H{
			"error": "id query parameter is required",
		})
		return
	}

	user := schemas.Usuario{}

	if err := db.First(&user, id).Error; err != nil {
		ctx.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	if request.Nome != "" {
		user.Nome = request.Nome
	}
	if request.IsAdmin != nil {
		user.IsAdmin = *request.IsAdmin
	}

	if err := db.Save(&user).Error; err != nil {
		ctx.JSON(400, gin.H{
			"error": "failed updating user",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "user update successfully",
		"data":    user,
	})
}
