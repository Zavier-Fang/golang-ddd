package repository

import userEntity "server/domain/user/entity"
import departmentEntity "server/domain/department/entity"

type UserRepository interface {
	QueryById(id int) (*userEntity.User, error)
	Create(user *userEntity.User) (int, error)
}

type DepartmentRepository interface {
	QueryById(id int) (*departmentEntity.Department, error)
}
