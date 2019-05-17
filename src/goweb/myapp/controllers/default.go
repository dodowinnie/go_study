package controllers

import (
	"github.com/astaxie/beego"
	// "strconv"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// c.Data["Website"] = "golang.org"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"
	c.Ctx.WriteString("appname:" + beego.AppConfig.String("appname") + 
	"\nhttpport:" + beego.AppConfig.String("httpport") + "\nrunmode:" + beego.AppConfig.String("runmode"))

	beego.Trace("trace test1")
	beego.Informational("info test1")
	beego.SetLevel(beego.LevelInformational)
	beego.Trace("trace test2")
	beego.Informational("info test2")
}
