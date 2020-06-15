package controller

import (
  "github.com/kataras/iris/v12"
)

type TestController struct {
}

func (_ TestController) Handler(ctx iris.Context) {
	ctx.Text("/handler")
}