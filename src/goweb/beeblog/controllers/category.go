package controllers

import (
	"goweb/beeblog/models"

	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	this.TplName = "category.html"
	this.Data["IsCategory"] = true
	this.Data["isLogin"] = checkAccount(this.Ctx)

	opt := this.GetString("opt")
	switch opt {
	case "add":
		name := this.GetString("name")
		if name == "" {
			return
		}
		err := models.AddCatetory(name)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	case "delete":
		id := this.GetString("id")
		if len(id) == 0 {
			break
		}
		err := models.DeleteCategory(id)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	}
	var err error
	this.Data["Categorys"], err = models.GetAllCategory()
	if err != nil {
		beego.Error(err)
	}
}
