package middleware

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"

	"login/http"
)

var JWT *jwt.Middleware

func initJWT() {
	JWT = jwt.New(jwt.Config{
		// 错误处理
		ErrorHandler: func(ctx iris.Context, err error) {
			if err == nil {
				return
			}
			ctx.StopExecution()

			http.Error401(ctx, err.Error())
		},

		// 设置一个函数返回秘钥，关键在于return []byte("这里设置秘钥")
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("My Secret"), nil
		},

		// 设置一个加密方法
		SigningMethod: jwt.SigningMethodHS256,
	})
}
