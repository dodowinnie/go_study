package main

import (
	_ "goweb/myapp/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

