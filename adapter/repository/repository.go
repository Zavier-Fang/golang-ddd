package repository

import "server/domain/user/entity"

type UserRepository interface {
	QueryById(id int) ([]*entity.User, error)
	Create(user *entity.User) (int, error)
}
