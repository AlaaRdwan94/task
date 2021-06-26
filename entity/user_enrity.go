package entity

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	FullName  string `sql:"-"`
	PhotoUrl  string
	Email     string `gorm:"unique"`
	PassWord  string
	Phone     string
}

func (p *User) TableName() string {
	return "users"
}
