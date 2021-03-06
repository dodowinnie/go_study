package main

import (
	"goweb/beeblog/controllers"
	"goweb/beeblog/models"
	_ "goweb/beeblog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	// 自动路由
	beego.AutoRouter(&controllers.TopicController{})

	beego.Run()
}
