package structs

import "github.com/astaxie/beego/validation"

var MessageTmpls = map[string]string{
	"Required":     "不能为空或者值不正确",
	"Min":          "不能小于%d",
	"Max":          "不能大于%d",
	"Range":        "范围在 %d 和 %d 之间",
	"MinSize":      "长度必须大于 %d",
	"MaxSize":      "长度必须小于 %d",
	"Length":       "格式不正确",
	"Alpha":        "必须是一个alpha字符",
	"Numeric":      "格式不正确",
	"AlphaNumeric": "必须是一个alpha字符或者数字",
	"Match":        "格式不正确",
	"NoMatch":      "格式不正确",
	"AlphaDash":    "必须为alpha 字符或数字或横杠",
	"Email":        "格式不正确",
	"IP":           "Ip格式不正确",
	"Base64":       "不是一个有效的base64值",
	"Mobile":       "手机号码格式不正确",
	"Tel":          "电话号码格式不正确",
	"Phone":        "手机或固定号码格式不正确",
	"ZipCode":      "邮编格式不正确",
}

func init()  {
	validation.SetDefaultMessage(MessageTmpls)
}

//字段名映射到中文
func GetCnField(field string) string {
	//关系表
	fieldMapping := map[string]string{
		"CompanyAddress":"单位地址",
		"CompanyName":"单位名称",
		"Email":"邮箱",
		"IdCard":"身份证",
		"IssuingOrgan":"发证机关",
		"Name":"姓名",
		"Phone":"手机号",
		"UserId":"用户id",
	}
	if val, ok := fieldMapping[field]; ok {
		return val
	}
	return ""
}