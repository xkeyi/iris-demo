package http

import (
	"github.com/kataras/iris/v12"
)

// ErrorModel 错误返回模型
type ErrorModel struct {
	Code   int64  `json:"code"`
	Msg    string `json:"msg"`
	Detail string `json:"detail"`
}

func buildError(code int64, msg string, detail string) ErrorModel {
	return ErrorModel{
		Code:   code,
		Msg:    msg,
		Detail: detail,
	}
}

func Error400(ctx iris.Context, detail string) {
	ctx.StatusCode(iris.StatusBadRequest)
	ctx.JSON(buildError(iris.StatusBadRequest, "错误请求", detail))
}

func Error401(ctx iris.Context, detail string) {
	ctx.StatusCode(iris.StatusUnauthorized)
	ctx.JSON(buildError(iris.StatusUnauthorized, "未授权", detail))
}

func Error403(ctx iris.Context, detail string) {
	ctx.StatusCode(iris.StatusForbidden)
	ctx.JSON(buildError(iris.StatusForbidden, "拒绝请求", detail))
}

func Error404(ctx iris.Context, detail string) {
	ctx.StatusCode(iris.StatusNotFound)
	ctx.JSON(buildError(iris.StatusNotFound, "未找到", detail))
}

func Error422(ctx iris.Context, detail string) {
	ctx.StatusCode(iris.StatusUnprocessableEntity)
	ctx.JSON(buildError(iris.StatusUnprocessableEntity, "请求格式错误", detail))
}

func Error500(ctx iris.Context, detail string) {
	ctx.StatusCode(iris.StatusInternalServerError)
	ctx.JSON(buildError(iris.StatusInternalServerError, "服务器内部错误", detail))
}

// 查询获取资源成功时返回查询到的数据
func Success200(ctx iris.Context, data iris.Map) {
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(data)
}

// 创建资源，创建成功后返回新创建的数据
func Success201(ctx iris.Context, data iris.Map) {
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(data)
}

// 操作成功后（如删除操作）返回，没有数据
func Success204(ctx iris.Context) {
	ctx.StatusCode(iris.StatusNoContent)
	ctx.JSON(iris.Map{})
}