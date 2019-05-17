package controllers

import (
	"goweb/beeblog/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplName = "home.html"
	this.Data["IsHome"] = true

	res := checkAccount(this.Ctx)
	this.Data["isLogin"] = res

	topics, err := models.GetAllTopics(true)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics

}
