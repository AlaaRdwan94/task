package user

import (
	"github.com/InnoSoft/task/entity"
	"github.com/InnoSoft/task/model"
)

func CreateAccountTransform(userModel *model.UserData) *entity.User {
	return &entity.User{
		FirstName: userModel.FName,
		LastName:  userModel.LName,
		FullName:  userModel.FName + " " + userModel.LName,
		PhotoUrl:  userModel.ProfilePictureUrl,
		Email:     userModel.Email,
		PassWord:  userModel.Password,
		Phone:     userModel.Phone,
	}
}

func GetAccountDataTransform(user *entity.User) *model.UserData {
	return &model.UserData{
		ID:                user.ID,
		FName:             user.FirstName,
		LName:             user.LastName,
		FullName:          user.FirstName + " " + user.FirstName,
		Email:             user.Email,
		Password:          user.PassWord,
		ProfilePictureUrl: user.PhotoUrl,
		Phone:             user.Phone,
	}
}
