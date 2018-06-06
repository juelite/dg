package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"dg/services"
)

//用户模型
type Users struct {
	Id 					int
	Phone				string			//手机号码
	Password 			string			//密码
	Salt 			    string			//盐
	CreateTime			time.Time
	UpdateTime			time.Time
}

func init() {
	orm.RegisterModel(new(Users))
}

//用户注册
func (u *Users) Add(o orm.Ormer, user Users) (int64, error) {
	user.UpdateTime = time.Now()
	user.CreateTime = time.Now()
	commonSrv := &services.CommonSrv{}
	user.Salt = commonSrv.RandomString(8)
	user.Password = commonSrv.GenPwd(user.Password, user.Salt)
	id, err := o.Insert(&user)
	return id, err
}

//修改密码
func (u *Users) ChangePwd(o orm.Ormer, user Users) (int64, error) {
	user.UpdateTime = time.Now()
	commonSrv := &services.CommonSrv{}
	user.Salt = commonSrv.RandomString(8)
	user.Password = commonSrv.GenPwd(user.Password, user.Salt)
	num, err := o.Update(&user)
	return num, err
}