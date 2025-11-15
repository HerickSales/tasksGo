package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.JSON(401, gin.H{
				"error": "authorization header required",
			})
			ctx.Abort()
			return
		}
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(401, gin.H{
				"error": "invalid token",
			})
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx.Set("userID", uint(claims["userID"].(float64)))
			ctx.Set("nome", claims["nome"].(string))
			ctx.Set("isAdmin", claims["isAdmin"].(bool))
		}

		ctx.Next()
	}
}
