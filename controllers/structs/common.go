package structs

const (
	SUCCESS_CODE = 0
	ERROR_CODE = 1

)

//token 解析模板
type TokenTmp struct {
	Uid 		int
	UnixTime 	int64
	Phone 		string
}