package service

import (
	"login/database"
)

func NewUserService() UserService {
	return UserService{
		db: database.DB,
	}
}