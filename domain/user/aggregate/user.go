package aggregate

import (
	departmentEntity "server/domain/department/entity"
	userEntity "server/domain/user/entity"
)

type UserAggregate struct {
	Id           int    `json:"id"`
	Name         string `json:"name" binding:"required"`
	Phone        string `json:"phone"`
	DepartmentId int    `json:"department_id"`
	Department   string `json:"department"`
	LeaderId     int    `json:"leader_id"`
	Leader       string `json:"leader"`
}

func NewUserAggregate(user, leader *userEntity.User, department *departmentEntity.Department) *UserAggregate {
	return &UserAggregate{
		Id:           user.Id,
		Name:         user.Name,
		Phone:        user.Phone,
		DepartmentId: user.Id,
		Department:   department.Name,
		LeaderId:     user.LeaderId,
		Leader:       leader.Name,
	}
}
