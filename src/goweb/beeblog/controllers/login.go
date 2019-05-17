package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	isExist := this.GetString("exist") == "true"
	if isExist {
		// 清除cookie
		this.Ctx.SetCookie("username", "", -1, "/")
		this.Ctx.SetCookie("password", "", -1, "/")
		// 重定向首页
		this.Redirect("/login", 301)
		return
	}

	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	autoLogin := this.GetString("autologin") == "on"

	if username == "" || username != beego.AppConfig.String("username") {
		return
	}

	if password == "" || password != beego.AppConfig.String("password") {
		return
	}

	maxAge := 0
	if autoLogin {
		maxAge = 60
	}

	// 设置cookie
	this.Ctx.SetCookie("username", username, maxAge, "/")
	this.Ctx.SetCookie("password", password, maxAge, "/")

	// 重定向首页
	this.Redirect("/", 301)
	return

}

func checkAccount(ctx *context.Context) bool {
	username := ctx.GetCookie("username")
	password := ctx.GetCookie("password")
	if username == "" || password == "" {
		return false
	}
	res := username == beego.AppConfig.String("username") && password == beego.AppConfig.String("password")

	return res
}
