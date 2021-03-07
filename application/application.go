package application

import (
	"server/domain/user/aggregate"
	"server/domain/user/entity"
)

type UserApplication interface {
	QueryUserById(id int) (*aggregate.UserAggregate, error)
	CreateUser(user *entity.User) (int, error)
}
