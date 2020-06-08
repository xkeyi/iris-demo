package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"login/model"
	"login/util"
)

type AuthController struct {
	Ctx iris.Context
}

type UserRegister struct {
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
}

func (ac *AuthController) Get() {
	iris.New().Logger().Info(" Get Test ")
}

func (ac *AuthController) Login() {
	iris.New().Logger().Info(" Post Login ")
}

func (ac *AuthController) Register() mvc.Result {
	iris.New().Logger().Info(" user Register ")

	var registerData UserRegister
	ac.Ctx.ReadJSON(&registerData)

	// 验证
	// 查询用户名是否存在
	_, isExist := userService.GetByUsername(registerData.Username)
	if isExist {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  "1",
				"msg": "用户名已被占用",
			},
		}
	}
	// 密码加密
	hashed, _ := util.GeneratePassword(registerData.Password)

	user := model.User{
		Username: registerData.Username,
		Name: registerData.Name,
		Password: hashed,
	}

	userId, err := userService.Insert(user)
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  "1",
				"msg": err,
			},
		}
	}

	// 生成 token
	// 返回 token

	return mvc.Response{
		Object: map[string]interface{}{
			"status":  "0",
			"user_id": userId,
		},
	}
}