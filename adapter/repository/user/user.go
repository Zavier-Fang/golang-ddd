package user

import (
	"gorm.io/gorm"
	"log"
	"server/adapter/repository"
	"server/domain/user/entity"
	"server/infrastructure/mysql"
)

type Repository struct {
	db *gorm.DB
}

func NewUserRepository() repository.UserRepository {
	return &Repository{db: mysql.DB}
}

func (repo *Repository) QueryById(id int) ([]*entity.User, error) {
	var users []*entity.User
	if err := repo.db.Where("id = ?", id).Find(&users).Error; err != nil {
		log.Printf("failed to query user by id , id: %d, error: %s", id, err.Error())
		return nil, err
	}
	return users, nil
}

func (repo *Repository) Create(user *entity.User) (int, error) {
	if err := repo.db.Create(user).Error; err != nil {
		log.Printf("failed to insert user, err: %s", err.Error())
		return 0, err
	}

	return user.Id, nil
}
