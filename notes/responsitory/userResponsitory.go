package responsitory

import (
	"notes/dao"
	"notes/model"
)

/**
通过用户名查找用户
*/
func FindUserByName(username string) (*model.User, bool ) {
	var user model.User
	if ok := dao.DB.Where("username = ?", username).First(&user).RecordNotFound(); ok {
		return nil, false
	}
	return &user,true
}

/**
通过邮箱查询用户
 */
func FindUserByEmail(email string) (*model.User, bool ) {
	var user model.User
	if ok := dao.DB.Where("email = ?", email).First(&user).RecordNotFound(); ok {
		return nil, false
	}
	return &user,true
}

/**
保存用户信息
*/
func InsertUser(user *model.User) bool {
	if dao.DB.NewRecord(user){
		err := dao.DB.Create(user).Error
		if err != nil{
			return false
		}
		return true
	}
	return false
}
