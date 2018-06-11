package structs

type RegisterIn struct {
	Email 				string	`valid:"Email" form:"email"`
	Password 			string 	`valid:"Required;MinSize(6);MaxSize(15)" form:"password"`
	PasswordConfirm 	string 	`valid:"Required;MinSize(6);MaxSize(15)" form:"password_confirm"`
}

type LoginIn struct {
	Email 				string	`valid:"Email" form:"email"`
	Password 			string 	`valid:"Required;MinSize(6);MaxSize(15)" form:"password"`
}

type ChangePwd struct {
	Token 				string 	`valid:"Required" form:"token"`
	Password 			string 	`valid:"Required;MinSize(6);MaxSize(15)" form:"password"`
}
