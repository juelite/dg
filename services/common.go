package services

import (
	"time"
	"math/rand"
	"crypto/md5"
	"encoding/hex"
)

type CommonSrv struct{}

/**
 * 生成随机字符串
 * @param length int 生成随机字符串的长度
 */
func (c *CommonSrv) RandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

/**
 * 创建密码
 * @param pwd string 明文密码
 * @param salt string 盐
 */
func (c *CommonSrv) GenPwd(pwd string, salt string) string {
	ctx := md5.New()
	ctx.Write([]byte(pwd+salt))
	return hex.EncodeToString(ctx.Sum(nil))
}

/**
 * 校验密码
 * @param pwd string 明文密码
 * @param salt string 盐
 * @param pwdEc string 正确的密码
 */
func (c *CommonSrv) CheckPwd(pwd string, salt string, pwdEc string) bool {
	pwdNew := c.GenPwd(pwd, salt)
	if pwdNew == pwdEc {
		return true
	}
	return false
}
