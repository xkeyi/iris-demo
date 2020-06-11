package service

import (
	"github.com/go-xorm/xorm"

	"login/model"
	"login/util"
)

type UserService struct {
	db *xorm.Engine
}

func (us *UserService) Insert(user model.User) (int64, error) {
	if _, err := us.db.Insert(&user); err != nil {
		return 0, err
	}

	return user.Id, nil
}

func (us *UserService) GetByUsername(username string) (model.User, bool) {
	var user model.User
	us.db.Where("username = ?", username).Get(&user)

	return user, user.Id > 0
}

func (us *UserService) GetByUserId(userId int64) (model.User) {
	var user model.User
	us.db.Where("id = ?", userId).Get(&user)

	return user
}

func (us *UserService) Login(username string, password string) (int64) {
	var user model.User
	us.db.Where("username = ?", username).Get(&user)

	if user.Id == 0 {
		return 0
	}

	// 验证密码
	if !util.ValidatePassword(password, user.Password) {
		return 0
	}

	return user.Id
}