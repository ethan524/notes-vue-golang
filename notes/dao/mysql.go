package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/ini.v1"
)

var (
	DB *gorm.DB
)

func InitMysql()(*gorm.DB){

	cfg,err := ini.Load("./config/config.ini")
	if err != nil{
		fmt.Printf("load config.ini failed, err:%v", err)
		return nil
	}

	myIni := map[string]string{
		"username":cfg.Section("mysql").Key("username").String(),
		"password":cfg.Section("mysql").Key("passowrd").String(),
		"address":cfg.Section("mysql").Key("address").String(),
		"port":cfg.Section("mysql").Key("port").String(),
		"database":cfg.Section("mysql").Key("database").String(),
	}
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", myIni["username"], myIni["password"], myIni["address"], myIni["port"], myIni["database"])

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil
	}
	DB = db
	return DB
}

func GetDB() *gorm.DB{
	return DB
}

func Close(){
	DB.Close()
}
