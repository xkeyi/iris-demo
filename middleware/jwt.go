package middleware

import (
	"github.com/iris-contrib/middleware/jwt"
)

var JWT *jwt.Middleware

func initJWT() {
	JWT = jwt.New(jwt.Config{
		// 设置一个函数返回秘钥，关键在于return []byte("这里设置秘钥")
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("My Secret"), nil
		},

		// 设置一个加密方法
		SigningMethod: jwt.SigningMethodHS256,
	})
}
