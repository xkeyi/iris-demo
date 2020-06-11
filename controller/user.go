package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"login/middleware"
	"login/util"
)

type UserController struct {
	Ctx iris.Context
}

func (uc *UserController) BeforeActivation(b mvc.BeforeActivation) {
	// b.Dependencies().Add/Remove
	// b.Router().Use/UseGlobal/Done
	// and any standard Router API call you already know

	// 1-> Method
	// 2-> Path
	// 3-> The controller's function name to be parsed as handler
	// 4-> Any handlers that should run before the MyCustomHandler
	//b.Handle("GET", "/something/{id:long}", "MyCustomHandler", anyMiddleware...)

	b.Handle("POST", "/me", "Me", middleware.JWT.Serve)
}

func (uc *UserController) Me() {
	userId := util.GetTokenUserId(uc.Ctx)

	user := userService.GetByUserId(userId)
	uc.Ctx.JSON(user.ResponseUser())
}