package container

import (
	"log"
	"server/adapter/controller"
	userController "server/adapter/controller/user"
	"server/adapter/repository"
	departmentRepo "server/adapter/repository/department"
	userRepo "server/adapter/repository/user"
	"server/application"
	userApp "server/application/user"
	"server/domain"
	departmentDomain "server/domain/department"
	userDomain "server/domain/user"
	"server/infrastructure/config"
	"server/infrastructure/http"
	"server/infrastructure/mysql"
)

const (
	SINGLE = "singleton"
)

type Container struct {
	userController controller.UserController
	userRepo       repository.UserRepository
	departmentRepo repository.DepartmentRepository

	app          application.UserApplication
	departDomain domain.Department
	userDomain   domain.User
}

func (c *Container) Register(mode string) *Container {
	if mode == SINGLE {
		c.userRepo = userRepo.NewUserRepository(mysql.DB)
		c.departmentRepo = departmentRepo.NewDepartmentRepository(mysql.DB)
		c.userDomain = userDomain.User{
			Repo: c.userRepo,
		}
		c.departDomain = departmentDomain.Department{
			Repo: c.departmentRepo,
		}
		c.app = userApp.Application{
			User:       c.userDomain,
			Department: c.departDomain,
		}
		c.userController = userController.UserController{
			App: c.app,
		}

		return c
	}
	log.Fatal("unsupported mode: ", mode)
	return nil
}

func (c *Container) Run() {
	c.userController.InitRoute(http.Engine)

	if err := http.Engine.Run(":" + config.Config.GetString("http.port")); err != nil {
		log.Fatal("run http server failed! err: ", err.Error())
	}
}
