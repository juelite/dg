package controllers

import (
	"dg/controllers/structs"
	"dg/models"
	"github.com/astaxie/beego/orm"
	"dg/services"
	"fmt"
)

type UsersController struct {
	BaseController
}

var data map[string]interface{}

//注册
func (u *UsersController) Register() {
	param := &structs.RegisterIn{}
	u.RequestData(param)
	if param.Password != param.PasswordConfirm {
		u.ReturnData(structs.ERROR_CODE, "两次密码输入不一致", data)
	}
	users := &models.Users{
		Email:param.Email,
		Password:param.Password,
	}
	o := orm.NewOrm()
	id, err := users.Add(o, *users)
	if err != nil || id < 1 {
		u.ReturnData(structs.ERROR_CODE, "注册失败或该手机号码已存在", data)
	}
	u.ReturnData(structs.SUCCESS_CODE, "success", data)
}

//登录
func (u *UsersController) Login() {
	param := &structs.LoginIn{}
	u.RequestData(param)
	users_model := &models.Users{}
	user, err := users_model.GetUserByPhone(param.Email)
	if err != nil {
		u.ReturnData(structs.ERROR_CODE, "手机号不存在或密码错误", data)
	}
	commonSrv := &services.CommonSrv{}
	if commonSrv.CheckPwd(param.Password, user.Salt, user.Password) == false {
		u.ReturnData(structs.ERROR_CODE, "手机号不存在或密码错误", data)
	}
	ret := make(map[string]interface{}, 1)
	ret["token"] = commonSrv.GenToken(user.Id, user.Email)
	data = ret
	u.ReturnData(structs.SUCCESS_CODE, "success", data)
}

//修改密码
func (u *UsersController) ChangePwd() {
	param := &structs.ChangePwd{}
	u.RequestData(param)
	commonSrv := &services.CommonSrv{}
	fmt.Println(commonSrv.ParseToken(param.Token))
	u.ReturnData(structs.SUCCESS_CODE, "success", data)
}
