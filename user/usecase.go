package user

import "github.com/InnoSoft/task/model"

type Usecase interface {
	CreateAccount(model *model.UserData) (uid uint , err error)
	GetUserById(uid uint) (*model.UserData, error)
	UpdateProfilePic(url string, id uint) (*model.UserData, error)
	GetUserByEmail(email string) (*model.UserData, error)
}
