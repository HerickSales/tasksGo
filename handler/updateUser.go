package handler

import (
	"github.com/elipe1/tasksGo-tarefa2/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateUserHandler(ctx *gin.Context) {
	isAdmin, exists := ctx.Get("isAdmin")
	if !exists {
		ctx.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}

	if !isAdmin.(bool) {
		ctx.JSON(401, gin.H{
			"error": "only admins can update users",
		})
		return
	}

	request := UpdateUserRequest{}
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

	user := schemas.Usuario{}

	if err := db.First(&user, id).Error; err != nil {
		ctx.JSON(404, gin.H{
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
		ctx.JSON(500, gin.H{
			"error": "failed updating user",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "user updated successfully",
		"data":    user,
	})
}
