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

	//mvc.New(app.Party("/auth")).
	//	Handle(new(controller.AuthController))

	v1 := app.Party("/v1")

	{
		// 用户认证相关
		mvc.New(v1.Party("/auth")).Handle(new(controller.AuthController))

		// 用户相关
		mvc.New(v1.Party("/users")).Handle(new(controller.UserController))

		//authController := &controller.AuthController{}
		//v1.Get("/test", authController.GetTest)
		//routeAuth(v1)
	}

	// 用户认证相关
	//mvc.New(app.Party("/auth")).Handle(new(controller.AuthController))
}

func routeUser(v1 iris.Party) {
	// 需要中间键的路由组
	//user1 := v1.Party("/user", middleware)
	//{
	//	user1.Post()
	//}
	//
	//// 不需要中间键的路由组
	//user2 := v1.Party("/users")
	//{
	//	user2.Post()
	//}
}