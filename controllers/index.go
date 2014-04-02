package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var log = logs.NewLogger(1)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	user := this.GetSession("user")
	if user != nil {
		this.Data["Website"] = "beego.me"
		this.Data["Email"] = "astaxie@gmail.com"
		this.TplNames = "index.tpl"
	} else {
		this.Redirect("/login", 302)
	}
}
