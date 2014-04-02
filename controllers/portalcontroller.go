package controllers

import (
	"github.com/astaxie/beego"
)

type PortalContrller struct {
	beego.Controller
}

func init() {
	log.SetLogger("file", `{"filename":"cim.log"}`)
}
func (this *PortalContrller) Get() {
	this.TplNames = "login.tpl"
}

func (this *PortalContrller) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	user := userDB.GetUser(username)

	if user.Id != 0 && user.Password == password {
		log.Debug("user[%s]login successful.", user.Username)
		user.Password = ""
		//add user to session
		this.SetSession("user", user)
		this.Redirect("/", 302)
	} else {
		log.Debug("user[%s]login fail. ", username)
		this.Data["Error"] = "The user name or password is not correct."
		this.TplNames = "login.tpl"
	}
}
