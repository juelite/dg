package controllers

import (
	"github.com/astaxie/beego"
	"dg/services"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	commonSrv := &services.CommonSrv{}
	salt := commonSrv.RandomString(8)
	fmt.Println(salt)
	pwd := "wangyu"
	pwdEc := commonSrv.GenPwd(pwd, salt)
	fmt.Println(pwdEc)
}
