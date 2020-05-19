package common

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

/**
	密码加密
 */
func GeneratePassword(pwd string) ([]byte, error){
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}

/**
	密码对比 明文对比加密
 */
func ValidatePassword(pwd string, hashed string) (isOk bool, err error){
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pwd)); err != nil{
		return false, errors.New("密码错误")
	}
	return true, nil
}

/**
	验证邮箱
 */
func ValidateEmail(email string) bool {
	regx := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	r := regexp.MustCompile(regx)
	return r.MatchString(email)
}

/**
	验证手机号码
 */
func ValidatePhone(telphone string) bool{
	regx := `^1([358][0-9]|4[579]|66|7[0135678]|9[89])[0-9]{8}$`
	r := regexp.MustCompile(regx)
	return r.MatchString(telphone)
}