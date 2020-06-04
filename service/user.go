package service

import (
	"github.com/go-xorm/xorm"

	"login/model"
)

type UserService struct {
	db *xorm.Engine
}

func (us UserService) Insert(user model.User) (int64, error) {
	if _, err := us.db.Insert(&user); err != nil {
		return 0, err
	}

	return user.Id, nil
}

func (us UserService) GetByUsername(username string) (model.User, bool) {
	var user model.User
	us.db.Where("username = ?", username).Get(&user)

	return user, user.Id > 0
}