package user

import (
	"server/adapter/repository"
	"server/adapter/repository/user"
	"server/domain"
	"server/domain/user/entity"
)

// 领域服务，与entity的user不同
type User struct {
	Repo repository.UserRepository
}

func NewUserDomain() domain.User {
	return &User{Repo: user.NewUserRepository()}
}

func (u User) QueryById(id int) ([]*entity.User, error) {
	return u.Repo.QueryById(id)
}

func (u User) Create(user *entity.User) (int, error) {
	return u.Repo.Create(user)
}
