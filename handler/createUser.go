package handler

import (
	"github.com/elipe1/tasksGo-tarefa2/schemas"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		ctx.JSON(400, gin.H{ // Bad Request
			"error": err.Error(),
		})
		return
	}

	if request.Nome == "" {
		ctx.JSON(400, gin.H{
			"error": "name parameter is required",
		})
		return
	}

	if request.Senha == "" {
		ctx.JSON(400, gin.H{
			"error": "senha parameter is required",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Senha), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	user := schemas.Usuario{
		Nome:    request.Nome,
		IsAdmin: request.IsAdmin,
		Senha:   string(hashedPassword),
	}

	if request.Nome == "" {
		ctx.JSON(400, gin.H{
			"error": "name parameter is required",
		})
		return
	}

	if request.Senha == "" {
		ctx.JSON(400, gin.H{
			"error": "senha parameter is required",
		})
		return
	}

	if err := db.Create(&user).Error; err != nil {
		ctx.JSON(500, gin.H{ // Internal Server Error
			"error": "failed to create user",
		})
		return
	}

	ctx.JSON(201, gin.H{ // Created
		"message": "user created successfully",
		"data":    user,
	})
}
