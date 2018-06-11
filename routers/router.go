package routers

import (
	"dg/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/register", &controllers.UsersController{}, "POST:Register")
    beego.Router("/login", &controllers.UsersController{}, "POST:Login")
    beego.Router("/changePwd", &controllers.UsersController{}, "POST:ChangePwd")
}
