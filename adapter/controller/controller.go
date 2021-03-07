package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	InitRoute(engine *gin.Engine)
	QueryUserById(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
}

func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
