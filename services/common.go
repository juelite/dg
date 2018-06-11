package services

import (
	"time"
	"math/rand"
	"crypto/md5"
	"encoding/hex"
	"dg/lib"
	"strconv"
	"strings"
	"dg/controllers/structs"
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

/**
 * 生成token
 * @param uid int 用户id
 * @param phone string 手机
 */
func (c *CommonSrv) GenToken(uid int, phone string) string {
	tool := &lib.Tool{}
	str := strconv.Itoa(uid) + "," + strconv.FormatInt(time.Now().Unix(), 10) + "," + phone
	mw := tool.Encode(str)
	return mw
}

/**
 * 解析token
 * @param token string
 */
func (c *CommonSrv) ParseToken(token string) (*structs.TokenTmp) {
	tool := &lib.Tool{}
	mw := tool.Decode(token)
	str := strings.Split(mw, ",")
	sess := &structs.TokenTmp{}
	if len(str) == 3 {
		uid, _ := strconv.Atoi(str[0])
		phone := str[2]
		t, _ := strconv.ParseInt(str[1], 10, 64)
		sess.Phone = phone
		sess.Uid = uid
		sess.UnixTime = t
	}
	return sess
}
