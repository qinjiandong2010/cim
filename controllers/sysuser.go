package controllers

import (
	"cim/models"
	"github.com/astaxie/beego"
)

var userDB = models.UserDB{}

type UserContrller struct {
	beego.Controller
}

func init() {
	log.SetLogger("file", `{"filename":"cim.log"}`)
}

func (this *UserContrller) RegisterPage() {
	this.TplNames = "register.tpl"
}

//用户注册
func (this *UserContrller) Register() {
	username := this.GetString("username")
	password := this.GetString("password")
	email := this.GetString("email")
	sex, _ := this.GetInt("sex")

	user := userDB.GetUser(username)
	if user.Id != 0 {
		this.Data["Error"] = "User name already exists."
	} else {
		user.Username = username
		user.Password = password
		user.Email = email
		user.Sex = sex
		if userDB.AddUser(user) > 0 {
			this.Data["Success"] = "User registration successful."
			log.Debug("user[%s] add successful.", user.Username)
		} else {
			this.Data["Error"] = "User registration fail."
		}
	}
	this.TplNames = "register.tpl"
}
