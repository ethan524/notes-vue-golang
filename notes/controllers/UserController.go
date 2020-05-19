package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"notes/common"
	"notes/dao"
	"notes/model"
	"notes/response"
	"notes/responsitory"
	"time"
)

type User struct{
	DB *gorm.DB
}

func NewUser() *User {
	return &User{
		DB : dao.GetDB(),
	}
}

// 用户注册
func (u *User) Register(c *gin.Context){
	var user = model.User{}
	c.Bind(&user)
	email := user.Email
	password := user.Password

	if !common.ValidateEmail(email){
		response.Faild(c,nil,"邮箱格式错误")
		return
	}

	if len(password) < 6 || len(password) > 20{
		response.Faild(c,nil,"用户密码不能为空，长度在6~20个字符之间")
		return
	}

	_, ok := responsitory.FindUserByEmail(email)
	if ok {
		response.Faild(c,nil,"该邮箱已存在")
		return
	}

	pwd, _ := common.GeneratePassword(password)
	user = model.User{
		Email: email,
		Password: string(pwd),
		Status: "1",
		Create_at:time.Now().Format("2006-01-02 15:04:05"),
		Update_at: "0000-01-01 00:00:00",
	}
	ok = responsitory.InsertUser(&user)
	if !ok {
		response.Faild(c,nil,"注册失败")
		return
	}

	response.Success(c,gin.H{"data":user},"注册成功")
	return
}

// 用户登录
func (u *User) Login(c *gin.Context){
	var user = model.User{}
	c.Bind(&user)

	if len(user.Username) < 2 || len(user.Username) > 20{
		response.Faild(c,gin.H{},"请输入用户名，长度在2~20个字符之间")
		return
	}

	if len(user.Password) < 6 || len(user.Password) > 20{
		response.Faild(c,gin.H{},"请输入用户名密码，长度在6~20个字符之间")
		return
	}

	res, ok := responsitory.FindUserByName(user.Username)
	if !ok {
		response.Faild(c,gin.H{},"用户名或密码错误")
		return
	}

	isOk, _ :=common.ValidatePassword(user.Password, res.Password)
	if !isOk{
		response.Faild(c,gin.H{},"用户名或密码错误")
		return
	}

	c.Set("user", &user)

	response.Success(c,gin.H{"data":user},"登录成功")
	return
}





















