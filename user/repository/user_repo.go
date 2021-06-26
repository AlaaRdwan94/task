package repository

import (
	"encoding/json"
	"errors"
	"github.com/InnoSoft/task/entity"
	"github.com/InnoSoft/task/user"
	"github.com/jinzhu/gorm"
	"gopkg.in/redis.v3"
	"strconv"
	"time"
)

type UserRepo struct {
	db *gorm.DB
	rd *redis.Client
}

func (u *UserRepo) CreateUser(user entity.User) (uint, error) {
	//create account
	if err := u.db.Create(&user).Error; err != nil {
		return 0, err
	}
	j, err := json.Marshal(&user)
	if err != nil {
		return 0, err
	}
	//we create user by id as a key
	if err := u.rd.Set(strconv.Itoa(int(user.ID)), j, time.Hour).Err(); err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (u *UserRepo) UpdateUser(user *entity.User) error {
	//update account
	if err := u.db.Model(&entity.User{}).Where("id = ? ", user.ID).Update(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) GetUser(userid uint) (*entity.User, error) {
	var usr entity.User
	u.db.Where("id = ?", userid).First(&usr)
	return &usr, nil
}
func (u *UserRepo) GetUserByEmail(email string) (*entity.User, error) {
	var usr entity.User
	u.db.Where("email = ?", email).First(&usr)
	return &usr, nil
}

func (u *UserRepo) GetAllUsers() (*[]entity.User, error) {
	var users []entity.User
	u.db.Find(&users)
	if len(users) == 0 {
		return &users, errors.New("record not found")
	}
	return nil, nil
}

func NewUserRepo(db *gorm.DB, rd *redis.Client) user.Repositoy {
	return &UserRepo{
		db: db,
		rd: rd,
	}
}
