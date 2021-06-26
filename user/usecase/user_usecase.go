package usecase

import (
	"github.com/InnoSoft/task/entity"
	"github.com/InnoSoft/task/model"
	transformer "github.com/InnoSoft/task/transformer/user"
	"github.com/InnoSoft/task/user"
	"github.com/jinzhu/gorm"
)

type UserUsecase struct {
	userRepo user.Repositoy
}

func (u UserUsecase) CreateAccount(model *model.UserData) (uid uint , err error) {
	userEntity := transformer.CreateAccountTransform(model)
	return u.userRepo.CreateUser(*userEntity)
}

func (u UserUsecase) GetUserById(uid uint) (*model.UserData, error) {
	userdata, err := u.userRepo.GetUser(uid)
	if err != nil {
		return nil, err
	}
	return transformer.GetAccountDataTransform(userdata), nil
}

func (u UserUsecase) GetUserByEmail(email string) (*model.UserData, error) {
	userdata, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return transformer.GetAccountDataTransform(userdata), nil
}

func (u UserUsecase) UpdateProfilePic(url string, id uint) (*model.UserData, error) {
	userdata := entity.User{
		Model:     gorm.Model{
			ID:        id,
		},
		PhotoUrl:  url,
	}
	if err := u.userRepo.UpdateUser(&userdata)
	err != nil {
		return nil, err
	}
	data , err := u.userRepo.GetUser(id)
	if err != nil {
		return nil, err
	}
	return transformer.GetAccountDataTransform(data) , nil
}

func NewUserUsecase(userRepo user.Repositoy)  user.Usecase{
   return &UserUsecase{
	   userRepo: userRepo,
   }
}