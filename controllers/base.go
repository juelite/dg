package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"reflect"
	"strings"
)

type BaseController struct {
	beego.Controller
}

/**
 * 接收参数方法
 * @param param interface{} 对应 struct 地址
 */
func (b *BaseController) RequestData(param interface{}) {
	b.ParseForm(param)
	valida := validation.Validation{}
	valida.Valid(param)
	message := ""
	if valida.HasErrors() {
		for _ , e := range valida.Errors {
			message = e.Key + " " + e.Message
			break
		}
		var data map[string]interface{}
		b.ReturnData( 3999 , message , data)
	}
}


/**
 * 返回数据，输出json中断当前请求
 * @param code int 返回状态码
 * @param message string 返回提示信息
 * @param data map[string]interface{} 返回的数据体
 */
func (b *BaseController) ReturnData(code int, message string, data map[string]interface{}) {
	b.Data["json"] = map[string]interface{}{
		"code":    code,
		"message": message,
		"data":    data,
	}
	b.ServeJSON()
	b.StopRun()
}

/**
 * 将struct转化成map
 * @param : obj struct 要转换的struct
 * @return : data map[string]interface{} 转换后的map
 */
func (b *BaseController) Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[snakeString(t.Field(i).Name)] = v.Field(i).Interface()
	}
	return data
}

func snakeString(s string) string {
	data := make([]byte, 0, len(s) * 2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}
