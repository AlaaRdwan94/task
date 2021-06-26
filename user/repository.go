package user

import "github.com/InnoSoft/task/entity"

type Repositoy interface {
	CreateUser(user entity.User) (id uint ,err error)
	UpdateUser(user *entity.User) error
	GetUser(userid uint) (*entity.User, error)
	GetAllUsers() (*[]entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}
