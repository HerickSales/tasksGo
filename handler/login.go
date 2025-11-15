package handler

import (
	"time"

	"github.com/elipe1/tasksGo-tarefa2/schemas"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("secret-key")

func LoginHandler(ctx *gin.Context) {
	var request LoginRequest

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{
			"error": "invalid request",
		})
		return
	}

	var user schemas.Usuario

	if err := db.Where("nome = ?", request.Nome).First(&user).Error; err != nil {
		ctx.JSON(401, gin.H{
			"error": "invalid credentials - username",
		})
		return
	}

	// Verifica se a senha fornecida corresponde ao hash armazenado no banco de dados
	if err := bcrypt.CompareHashAndPassword([]byte(user.Senha), []byte(request.Senha)); err != nil {
		ctx.JSON(401, gin.H{
			"error": "invalid credentials - password",
		})
		return
	}

	// Gera o token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":  user.ID,
		"nome":    user.Nome,
		"isAdmin": user.IsAdmin,
		"exp":     time.Now().Add(48 * time.Hour).Unix(), // Token expira em 24h - após isso, usuário deve fazer login novamente
	})

	// Assina o token com a chave secreta
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"token": tokenString,
		"user": gin.H{
			"id":      user.ID,
			"nome":    user.Nome,
			"isAdmin": user.IsAdmin,
		},
	})
}
