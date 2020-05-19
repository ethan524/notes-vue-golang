package main

import (
	"github.com/gin-gonic/gin"
	"notes/dao"
)

func main() {

	dao.InitMysql()
	defer dao.Close()

	r := gin.Default()
	r.Static("/upload", "./statics/uploads")
	r = Routers(r)

	r.Run(":8081")

}
