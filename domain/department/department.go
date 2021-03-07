package department

import (
	"server/adapter/repository"
	"server/domain/department/entity"
)

type Department struct {
	Repo repository.DepartmentRepository
}

func (d Department) QueryById(id int) (*entity.Department, error) {
	return d.Repo.QueryById(id)
}
