package handler

import (
	"github.com/elipe1/tasksGo-tarefa2/schemas"
	"github.com/gin-gonic/gin"
)

func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		ctx.JSON(500, gin.H{
			"error": "invalid request body",
		})
		return
	}

	user := schemas.Usuario{
		Nome:    request.Nome,
		IsAdmin: request.IsAdmin,
	}

	if err := db.Create(&user).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to create user",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "user created successfully",
		"data":    user,
	})
}
