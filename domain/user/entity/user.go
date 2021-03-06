package entity

type User struct {
	Id           int    `gorm:"column:id" json:"id"`
	Name         string `gorm:"column:name" json:"name" binding:"required"`
	Phone        string `gorm:"column:phone" json:"phone"`
	DepartmentId int    `gorm:"column:department_id" json:"department_id"`
	LeaderId     int    `gorm:"column:leader_id" json:"leader_id"`
}

func (User) TableName() string {
	return "user"
}
