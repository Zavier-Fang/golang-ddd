package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"server/adapter/controller"
	"server/application"
	"server/application/user"
	"server/domain/user/entity"
	"strconv"
)

type Controller struct {
	App application.UserApplication
}

func NewUserController(engine *gin.Engine) controller.UserController {
	ctrl := Controller{
		App: user.NewUserApplication(),
	}
	engine.GET("ping", controller.Ping)
	engine.GET("user/:id", ctrl.QueryUserById)
	engine.POST("user", ctrl.CreateUser)

	return &ctrl
}

func (ctrl Controller) QueryUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(400, gin.H{
			"message": "bad request",
		})
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("user id not number, error: %s, id : %v", err.Error(), id)
		ctx.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	users, err := ctrl.App.QueryUserById(userId)
	if err != nil {
		log.Printf("user id not number, error: %s, id : %v", err.Error(), id)
		ctx.JSON(500, gin.H{
			"message": "server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "success",
		"data":    users,
	})
}

func (ctrl Controller) CreateUser(ctx *gin.Context) {
	var user entity.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}
	id, err := ctrl.App.CreateUser(&user)
	if err != nil {
		log.Printf("create user failed, user: %v, err: %s", user, err.Error())
		ctx.JSON(500, gin.H{
			"message": "server error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "success",
		"data":    id,
	})
}