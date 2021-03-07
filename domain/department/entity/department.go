package entity

type Department struct {
	Id   int    `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (Department) TableName() string {
	return "department"
}
