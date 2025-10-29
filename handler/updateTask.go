package handler

import (
	"github.com/gin-gonic/gin"
)

func UpdateTaskHandler(ctx *gin.Context) {
	request := UpdateTaskRequest{}
	ctx.BindJSON(&request)
}
