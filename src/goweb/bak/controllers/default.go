package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["TrueCond"] = true
	c.Data["FalseCond"] = false
	c.TplName = "index.tpl"

	type u struct{
		Name string
		Age int
		Sex string
	}

	user := &u{
		Name : "brandon",
		Age : 20,
		Sex : "male",
	}
	c.Data["User"] = user

	nums := []int{1,2,3,4,5,6,7,8,9}
	c.Data["Nums"] = nums
	c.Data["TmpVal"] = "hey guys"

	c.Data["html"] = "<div>hello beego</div>"

	c.Data["pipe"] = "<div>hello beego</div>"




}
