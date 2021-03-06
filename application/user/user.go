package user

import (
	"server/application"
	"server/domain"
	"server/domain/user"
	"server/domain/user/entity"
)

type Application struct {
	User domain.User
}

func NewUserApplication() application.UserApplication {
	return &Application{User: user.NewUserDomain()}
}

func (app Application) QueryUserById(id int) ([]*entity.User, error) {
	return app.User.QueryById(id)
}

func (app Application) CreateUser(user *entity.User) (int, error){
	return app.User.Create(user)
}
