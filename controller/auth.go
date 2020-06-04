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

func (ac *AuthController) PostRegister() mvc.Result {
	iris.New().Logger().Info(" user login ")

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