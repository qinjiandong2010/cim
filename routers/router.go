package routers

import (
	"cim/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/api/login", &controllers.PortalContrller{})
	beego.Router("/api/user", &controllers.UserContrller{}, "post:Register")

	beego.SetStaticPath("/", "views/html/")
}
