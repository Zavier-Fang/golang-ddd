package user

import (
	"gorm.io/gorm"
	"log"
	"server/adapter/repository"
	"server/domain/user/entity"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) QueryById(id int) (*entity.User, error) {
	var user entity.User
	if err := repo.db.Where("id = ?", id).First(&user).Error; err != nil {
		log.Printf("failed to query user by id , id: %d, error: %s", id, err.Error())
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) Create(user *entity.User) (int, error) {
	if err := repo.db.Create(user).Error; err != nil {
		log.Printf("failed to insert user, err: %s", err.Error())
		return 0, err
	}

	return user.Id, nil
}
