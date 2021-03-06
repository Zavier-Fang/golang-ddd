package domain

import "server/domain/user/entity"

type User interface {
	QueryById(id int) ([]*entity.User, error)
	Create(user *entity.User) (int, error)
}
