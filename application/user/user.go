package user

import (
	"log"
	"server/domain"
	"server/domain/user/aggregate"
	"server/domain/user/entity"
)

type Application struct {
	User       domain.User
	Department domain.Department
}

func (app Application) QueryUserById(id int) (*aggregate.UserAggregate, error) {
	user, err := app.User.QueryById(id)
	if err != nil {
		log.Printf("failed to query user by id, id: %d, err: %s", id, err.Error())
		return nil, err
	}
	leader, err := app.User.QueryById(user.LeaderId)
	if err != nil {
		log.Printf("failed to query leader by id, id: %d, err: %s", user.LeaderId, err.Error())
		return nil, err
	}
	department, err := app.Department.QueryById(user.DepartmentId)
	if err != nil {
		log.Printf("failed to query department by id, id: %d, err: %s", user.DepartmentId, err.Error())
		return nil, err
	}
	return aggregate.NewUserAggregate(user, leader, department), nil
}

func (app Application) CreateUser(user *entity.User) (int, error) {
	return app.User.Create(user)
}
