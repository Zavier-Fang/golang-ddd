package adapter

import (
	"log"
	"server/adapter/controller/user"
	"server/infrastructure/config"
	"server/infrastructure/http"
)

func Run() {
	user.NewUserController(http.Engine)

	if err := http.Engine.Run(":" + config.Config.GetString("http.port")); err != nil {
		log.Fatal("run http server failed! err: ", err.Error())
	}
}

