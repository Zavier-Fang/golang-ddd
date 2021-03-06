package application

import (
	"server/domain/user/entity"
)

type UserApplication interface {
	QueryUserById(id int) ([]*entity.User, error)
	CreateUser(user *entity.User) (int, error)
}


