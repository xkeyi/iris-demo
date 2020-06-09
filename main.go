package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"

	_ "github.com/go-sql-driver/mysql"

	"login/config"
	"login/database"
	"login/route"
	"login/http"
)
func main() {
	app := iris.New()
	// 设置日志级别
	app.Logger().SetLevel(config.Viper.GetString("server.logger.level"))
	// Add recover to recover from any http-relative panics
	app.Use(recover.New())
	// Add logger to log the requests to the terminal
	app.Use(logger.New())

	//应用App配置
	configation(app)

	db := database.DB
	app.Logger().Info("数据库测试：", db)

	// Globally allow options method to enable CORS
	app.AllowMethods(iris.MethodOptions)

	//注册视图文件 -- 只做 api 的话可以不用
	app.RegisterView(iris.HTML("./resources", ".html"))

	// Router
	route.Route(app)

	app.Run(iris.Addr(config.Viper.GetString("server.addr")), iris.WithoutServerError((iris.ErrServerClosed)))
}

/**
 * 项目设置
 */
func configation(app *iris.Application) {

	//配置 字符编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))

	//错误配置
	//未发现错误
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		http.Error400(ctx, "未找到相关的资源")
	})

	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		http.Error500(ctx, "服务器出错了，快去扣程序员小哥哥工资吧")
	})
}
