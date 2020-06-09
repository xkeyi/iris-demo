package controller

import (
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"login/model"
	"login/util"
	"login/middleware"
	"login/http"
)

type AuthController struct {
	Ctx iris.Context
}

type RegisterUser struct {
	Name      string  `json:"name"`
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
	b.Handle("POST", "/me", "Me", middleware.JWT.Serve)
}

func (ac *AuthController) Get() {
	iris.New().Logger().Info(" Get Test ")

	peter := RegisterUser{
		Name: "John",
		Username:  "Doe",
		Password:      "Neither FBI knows!!!",
	}
	//手动设置内容类型: ctx.ContentType("application/javascript")
	ac.Ctx.JSON(peter)
}

func (ac *AuthController) Login() {
	iris.New().Logger().Info(" Post Login ")

	//ac.Ctx.StatusCode(401)
	//ac.Ctx.JSON(iris.Map{
	//	"code": 401,
	//	"msg": "没有权限，请先认真",
	//})
	http.Error401(ac.Ctx, "没有权限，请先认真")
	return

	// 用户ID：1
	token, err := util.GetJWTString(1)
	if err != nil {
		ac.Ctx.StatusCode(500)
		ac.Ctx.JSON(iris.Map{
			"status": 500,
			"msg": err,
		})
	}

	ac.Ctx.JSON(iris.Map{
		"token": token,
		"exp": time.Now().Add(120 * time.Minute * time.Duration(1)).Unix(),
	})
}

func (ac *AuthController) Me() {
	peter := RegisterUser{
		Name: "llz",
		Username:  "llz",
	}
	//手动设置内容类型: ctx.ContentType("application/javascript")
	ac.Ctx.JSON(peter)
}

func (ac *AuthController) Register() {
	iris.New().Logger().Info(" user Register ")

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
	token, _ := util.GetJWTString(userId)

	http.Success201(ac.Ctx, iris.Map{
		"token": token,
		"exp":   time.Now().Add(120 * time.Minute * time.Duration(1)).Unix(),
	})
}