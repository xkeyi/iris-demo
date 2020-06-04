package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"login/controller"
)

func Route(app *iris.Application) {
	// 首页  -- 只做 api 的话可以不用
	app.Get("/", func(ctx iris.Context) {
		ctx.View("views/index.html")
	})

	// 用户认证相关
	mvc.New(app.Party("/auth")).Handle(new(controller.AuthController))
}