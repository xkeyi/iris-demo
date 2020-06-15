package controller

import (
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"login/http"
	"login/model"
	"login/util"
	"login/middleware"
)

type AuthController struct {
	Ctx iris.Context
}

type RegisterUser struct {
	Name      string  `json:"name"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
}

type LoginUser struct {
	Username  string  `json:"username"`
	Password  string  `json:"password"`
}

func (ac *AuthController) BeforeActivation(b mvc.BeforeActivation) {
	// b.Dependencies().Add/Remove
	// b.Router().Use/UseGlobal/Done
	// and any standard Router API call you already know

	// 1-> Method
	// 2-> Path
	// 3-> The controller's function name to be parsed as handler
	// 4-> Any handlers that should run before the MyCustomHandler
	//b.Handle("GET", "/something/{id:long}", "MyCustomHandler", anyMiddleware...)

	b.Handle("POST", "/register", "Register")
	b.Handle("POST", "/login", "Login")
	b.Handle("POST", "/refresh", "Refresh", middleware.JWT.Serve)
}

func (ac *AuthController) Login() {
	var login_user LoginUser
	ac.Ctx.ReadJSON(&login_user)

	// 验证
	if login_user.Username == "" || login_user.Password == "" {
		http.Error422(ac.Ctx, "用户名和密码均不能为空")
		return
	}

	userId := userService.Login(login_user.Username, login_user.Password)
	if userId == 0 {
		http.Error401(ac.Ctx, "用户名或密码错误")
		return
	}

	http.Success201(ac.Ctx, getToken(userId))
}

func (ac *AuthController) Register() {
	var register_user RegisterUser
	ac.Ctx.ReadJSON(&register_user)

	// 验证
	if register_user.Username == "" || register_user.Name == "" || register_user.Password == "" {
		http.Error422(ac.Ctx, "用户名、姓名、密码均不能为空")
		return
	}
	// 查询用户名是否存在
	_, isExist := userService.GetByUsername(register_user.Username)
	if isExist {
		http.Error422(ac.Ctx, "用户名已被占用")
		return
	}
	// 密码加密
	hashed, _ := util.GeneratePassword(register_user.Password)

	user := model.User{
		Username: register_user.Username,
		Name: register_user.Name,
		Password: hashed,
	}

	userId, _ := userService.Insert(user)

	http.Success201(ac.Ctx, getToken(userId))
}

func getToken(userId int64) iris.Map{
	token, _ := util.GetJWTString(userId)

	return iris.Map{
		"token": token,
		"exp":   time.Now().Add(120 * time.Minute * time.Duration(1)).Unix(),
	}
}

func (ac *AuthController) Refresh() {
	userId := util.GetTokenUserId(ac.Ctx)

	ac.Ctx.JSON(userId)
}