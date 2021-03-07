package domain

import userEntity "server/domain/user/entity"
import departmentEntity "server/domain/department/entity"

type User interface {
	QueryById(id int) (*userEntity.User, error)
	Create(user *userEntity.User) (int, error)
}

type Department interface {
	QueryById(id int) (*departmentEntity.Department, error)
}
