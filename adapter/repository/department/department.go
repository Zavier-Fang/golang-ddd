package department

import (
	"gorm.io/gorm"
	"log"
	"server/adapter/repository"
	"server/domain/department/entity"
)

type DepartmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) repository.DepartmentRepository {
	return &DepartmentRepository{db: db}
}

func (repo *DepartmentRepository) QueryById(id int) (*entity.Department, error) {
	var department entity.Department
	if err := repo.db.Where("id = ?", id).First(&department).Error; err != nil {
		log.Printf("failed to query user by id , id: %d, error: %s", id, err.Error())
		return nil, err
	}
	return &department, nil
}
