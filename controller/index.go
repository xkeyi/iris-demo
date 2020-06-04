package controller

import (
	"login/database"
	"login/service"
)

var db = database.DB
var userService = service.NewUserService()
